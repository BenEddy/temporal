package persistence_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	commonpb "go.temporal.io/api/common/v1"
	enumspb "go.temporal.io/api/enums/v1"
	historypb "go.temporal.io/api/history/v1"
	"go.temporal.io/api/serviceerror"
	enumsspb "go.temporal.io/server/api/enums/v1"
	historyspb "go.temporal.io/server/api/history/v1"
	persistencespb "go.temporal.io/server/api/persistence/v1"
	"go.temporal.io/server/chasm"
	"go.temporal.io/server/common"
	"go.temporal.io/server/common/dynamicconfig"
	"go.temporal.io/server/common/log"
	p "go.temporal.io/server/common/persistence"
	mockp "go.temporal.io/server/common/persistence/mock"
	"go.temporal.io/server/common/persistence/serialization"
	"go.temporal.io/server/common/persistence/versionhistory"
	"go.temporal.io/server/common/testing/testlogger"
	"go.temporal.io/server/service/history/tasks"
	"go.uber.org/mock/gomock"
)

func newTestUpdateRequest(keys []tasks.Key) *p.UpdateWorkflowExecutionRequest {
	toDelete := map[tasks.Category][]tasks.Key{
		tasks.CategoryTimer: keys,
	}
	return &p.UpdateWorkflowExecutionRequest{
		ShardID:     1,
		RangeID:     1,
		Mode:        p.UpdateWorkflowModeUpdateCurrent,
		ArchetypeID: chasm.WorkflowArchetypeID,
		UpdateWorkflowMutation: p.WorkflowMutation{
			ExecutionInfo:         &persistencespb.WorkflowExecutionInfo{NamespaceId: "ns", WorkflowId: "wid", ExecutionStats: &persistencespb.ExecutionStats{}},
			ExecutionState:        &persistencespb.WorkflowExecutionState{RunId: "rid", State: enumsspb.WORKFLOW_EXECUTION_STATE_RUNNING, Status: enumspb.WORKFLOW_EXECUTION_STATUS_RUNNING},
			NextEventID:           2,
			BestEffortDeleteTasks: toDelete,
		},
	}
}

func TestExecutionManager_DeletesTasksWhenEnabled(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store := mockp.NewMockExecutionStore(ctrl)
	store.EXPECT().GetName().AnyTimes().Return("mock-store")
	// Expect the main update call
	store.EXPECT().UpdateWorkflowExecution(gomock.Any(), gomock.Any()).Return(nil)

	// Expect CompleteHistoryTask for each key
	// Use gomock.Any for ctx and request pointer type
	expectedKeys := []tasks.Key{
		tasks.NewKey(time.Now().UTC(), 123),
		tasks.NewKey(time.Now().UTC().Add(time.Second), 456),
	}
	// Allow any order
	store.EXPECT().CompleteHistoryTask(gomock.Any(), gomock.Any()).Times(len(expectedKeys)).DoAndReturn(
		func(_ context.Context, req *p.CompleteHistoryTaskRequest) error {
			if req.ShardID != 1 || req.TaskCategory != tasks.CategoryTimer {
				t.Fatalf("unexpected CompleteHistoryTask request: %#v", req)
			}
			return nil
		},
	)

	serializer := serialization.NewSerializer()
	em := p.NewExecutionManager(
		store,
		serializer,
		nil,
		log.NewNoopLogger(),
		dynamicconfig.GetIntPropertyFn(1024*1024),
		dynamicconfig.GetBoolPropertyFn(true),
	)

	_, err := em.UpdateWorkflowExecution(context.Background(), newTestUpdateRequest(expectedKeys))
	if err != nil {
		t.Fatalf("UpdateWorkflowExecution returned error: %v", err)
	}
}

func TestExecutionManager_DoesNotDeleteTasksWhenDisabled(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store := mockp.NewMockExecutionStore(ctrl)
	store.EXPECT().GetName().AnyTimes().Return("mock-store")
	store.EXPECT().UpdateWorkflowExecution(gomock.Any(), gomock.Any()).Return(nil)
	// No CompleteHistoryTask expected

	serializer := serialization.NewSerializer()
	em := p.NewExecutionManager(
		store,
		serializer,
		nil,
		log.NewNoopLogger(),
		dynamicconfig.GetIntPropertyFn(1024*1024),
		dynamicconfig.GetBoolPropertyFn(false),
	)

	keys := []tasks.Key{tasks.NewKey(time.Now().UTC(), 789)}
	_, err := em.UpdateWorkflowExecution(context.Background(), newTestUpdateRequest(keys))
	if err != nil {
		t.Fatalf("UpdateWorkflowExecution returned error: %v", err)
	}
}

func TestExecutionManager_DeleteTasksError_DoesNotFailUpdate(t *testing.T) {
	// Test that errors in CompleteHistoryTask don't cause the UpdateWorkflowExecution to fail.
	// Task deletion is best-effort and should only be logged.
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store := mockp.NewMockExecutionStore(ctrl)
	store.EXPECT().GetName().AnyTimes().Return("mock-store")
	// Expect the main update call to succeed
	store.EXPECT().UpdateWorkflowExecution(gomock.Any(), gomock.Any()).Return(nil)

	// Expect CompleteHistoryTask to fail
	expectedKeys := []tasks.Key{
		tasks.NewKey(time.Now().UTC(), 123),
	}
	store.EXPECT().CompleteHistoryTask(gomock.Any(), gomock.Any()).Times(len(expectedKeys)).Return(
		serviceerror.NewInternal("simulated deletion error"),
	)

	serializer := serialization.NewSerializer()
	em := p.NewExecutionManager(
		store,
		serializer,
		nil,
		log.NewNoopLogger(),
		dynamicconfig.GetIntPropertyFn(1024*1024),
		dynamicconfig.GetBoolPropertyFn(true),
	)

	// UpdateWorkflowExecution should succeed even though CompleteHistoryTask failed
	_, err := em.UpdateWorkflowExecution(context.Background(), newTestUpdateRequest(expectedKeys))
	if err != nil {
		t.Fatalf("UpdateWorkflowExecution should succeed even if task deletion fails, got error: %v", err)
	}
}

func TestExecutionManager_TrimHistoryBranchSkipped_NonWorkflow(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store := mockp.NewMockExecutionStore(ctrl)
	store.EXPECT().GetName().AnyTimes().Return("mock-store")
	store.EXPECT().UpdateWorkflowExecution(gomock.Any(), gomock.Any()).Return(&p.ConditionFailedError{Msg: "conflict"})
	// trimHistoryNode logic skipped for non-workflow archetype and should be no call to GetWorkflowExecution.
	store.EXPECT().GetWorkflowExecution(gomock.Any(), gomock.Any()).Times(0)

	serializer := serialization.NewSerializer()
	em := p.NewExecutionManager(
		store,
		serializer,
		nil,
		// TrimHistoryBranch logic is best effort and only emits error log if failed.
		// Capture error logs to verify no unexpected errors.
		testlogger.NewTestLogger(t, testlogger.FailOnAnyUnexpectedError),
		dynamicconfig.GetIntPropertyFn(1024*1024),
		dynamicconfig.GetBoolPropertyFn(false),
	)

	req := newTestUpdateRequest(nil)
	req.ArchetypeID = chasm.WorkflowArchetypeID + 1 // Non-workflow archetype
	_, err := em.UpdateWorkflowExecution(context.Background(), req)
	if err == nil {
		t.Fatal("expected UpdateWorkflowExecution to return error")
	}
	if _, ok := err.(*p.ConditionFailedError); !ok {
		t.Fatalf("expected ConditionFailedError, got %T", err)
	}
}

func TestExecutionManager_TrimHistoryBranchSkipped_EmptyBranchToken(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store := mockp.NewMockExecutionStore(ctrl)
	store.EXPECT().GetName().AnyTimes().Return("mock-store")
	store.EXPECT().UpdateWorkflowExecution(gomock.Any(), gomock.Any()).Return(&p.ConditionFailedError{Msg: "conflict"})
	store.EXPECT().GetHistoryBranchUtil().Times(0)

	serializer := serialization.NewSerializer()
	versionHistories := versionhistory.NewVersionHistories(&historyspb.VersionHistory{}) // Empty branch token
	executionInfo := &persistencespb.WorkflowExecutionInfo{
		NamespaceId:      "ns",
		WorkflowId:       "wid",
		ExecutionStats:   &persistencespb.ExecutionStats{},
		VersionHistories: versionHistories,
	}
	executionState := &persistencespb.WorkflowExecutionState{
		RunId:  "rid",
		State:  enumsspb.WORKFLOW_EXECUTION_STATE_RUNNING,
		Status: enumspb.WORKFLOW_EXECUTION_STATUS_RUNNING,
	}
	executionInfoBlob, err := serializer.WorkflowExecutionInfoToBlob(executionInfo)
	if err != nil {
		t.Fatalf("WorkflowExecutionInfoToBlob error: %v", err)
	}
	executionStateBlob, err := serializer.WorkflowExecutionStateToBlob(executionState)
	if err != nil {
		t.Fatalf("WorkflowExecutionStateToBlob error: %v", err)
	}

	store.EXPECT().GetWorkflowExecution(gomock.Any(), gomock.Any()).Return(
		&p.InternalGetWorkflowExecutionResponse{
			State: &p.InternalWorkflowMutableState{
				ExecutionInfo:  executionInfoBlob,
				ExecutionState: executionStateBlob,
			},
		},
		nil,
	)

	em := p.NewExecutionManager(
		store,
		serializer,
		nil,
		// TrimHistoryBranch logic is best effort and only emits error log if failed.
		// Capture error logs to verify no unexpected errors.
		testlogger.NewTestLogger(t, testlogger.FailOnAnyUnexpectedError),
		dynamicconfig.GetIntPropertyFn(1024*1024),
		dynamicconfig.GetBoolPropertyFn(false),
	)

	_, err = em.UpdateWorkflowExecution(context.Background(), newTestUpdateRequest(nil))
	if err == nil {
		t.Fatal("expected UpdateWorkflowExecution to return error")
	}
	if _, ok := err.(*p.ConditionFailedError); !ok {
		t.Fatalf("expected ConditionFailedError, got %T", err)
	}
}

func TestExecutionManager_ReadHistoryBranchReverse_WrongPrevTxnID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	serializer := serialization.NewSerializer()
	branchUtil := p.NewHistoryBranchUtil(serializer)

	// Helper to serialize a batch of events into a blob (the persistence layer returns batch blobs).
	events := func(ids ...int64) *commonpb.DataBlob {
		events := make([]*historypb.HistoryEvent, 0, len(ids))
		for _, id := range ids {
			events = append(events, &historypb.HistoryEvent{EventId: id, Version: 1})
		}
		blob, err := serializer.SerializeEvents(events)
		require.NoError(t, err)
		return blob
	}

	// Create some history events. The original WF history is on root-branch and there are two resets.
	nodes := map[string][]p.InternalHistoryNode{
		"root-branch": {
			{NodeID: 4, TransactionID: 2000, PrevTransactionID: 1000, Events: events(4, 5)},
			{NodeID: 1, TransactionID: 1000, PrevTransactionID: 0, Events: events(1, 2, 3)},
		},
		"reset-1": {
			{NodeID: 9, TransactionID: 4000, PrevTransactionID: 3000, Events: events(9, 10)},
			{NodeID: 6, TransactionID: 3000, PrevTransactionID: 2000, Events: events(6, 7, 8)},
		},
		"reset-2": {
			{NodeID: 14, TransactionID: 6000, PrevTransactionID: 5000, Events: events(14, 15)},
			// Mis-wire the prev txn ID to simulate a bug
			{NodeID: 11, TransactionID: 5000, PrevTransactionID: 4, Events: events(11, 12, 13)},
		},
	}

	store := mockp.NewMockExecutionStore(ctrl)
	store.EXPECT().GetHistoryBranchUtil().AnyTimes().Return(branchUtil)
	store.EXPECT().ReadHistoryBranch(gomock.Any(), gomock.Any()).AnyTimes().DoAndReturn(
		func(_ context.Context, req *p.InternalReadHistoryBranchRequest) (*p.InternalReadHistoryBranchResponse, error) {
			return &p.InternalReadHistoryBranchResponse{Nodes: nodes[req.BranchID]}, nil
		},
	)

	em := p.NewExecutionManager(
		store,
		serializer,
		nil,
		log.NewNoopLogger(),
		dynamicconfig.GetIntPropertyFn(1024*1024),
		dynamicconfig.GetBoolPropertyFn(false),
	)

	// Generate a branch token for the reset-2 branch. It has two ancestors: root-branch and reset-1.
	branchToken, err := branchUtil.NewHistoryBranch(
		"ns", "wf-id", "wf-run-id", "tree", new("reset-2"),
		[]*persistencespb.HistoryBranchRange{
			{BranchId: "root-branch", BeginNodeId: 1, EndNodeId: 6},
			{BranchId: "reset-1", BeginNodeId: 6, EndNodeId: 11},
		},
		0, 0, 0,
	)
	require.NoError(t, err)

	_, err = getWorkflowExecutionHistoryReverse(t, em, branchToken, 6000)
	require.NoError(t, err)
}

// Simulate a client reading history via the GetWorkflowExecutionHistoryReverse API.
func getWorkflowExecutionHistoryReverse(
	t *testing.T,
	em p.ExecutionManager,
	branchToken []byte,
	lastFirstTxnID int64,
) ([]*historypb.HistoryEvent, error) {
	t.Helper()

	var allEvents []*historypb.HistoryEvent
	var pageToken []byte
	var nextEventID int64 = common.EmptyEventID

	// Continue paginating until we're done.
	for {
		events, _, tok, err := p.ReadFullPageEventsReverse(context.Background(), em, &p.ReadHistoryBranchReverseRequest{
			ShardID:                1,
			BranchToken:            branchToken,
			MaxEventID:             nextEventID,
			LastFirstTransactionID: lastFirstTxnID,
			PageSize:               100,
			NextPageToken:          pageToken,
		})
		if err != nil {
			return allEvents, err
		}

		allEvents = append(allEvents, events...)
		pageToken = tok

		// GetHistoryReverse in get_history_util.go computes the continuation token NextEventId.
		if len(events) > 0 {
			nextEventID = events[len(events)-1].EventId - 1
		}

		// Stop pagination when NextEventId drops below FirstEventId.
		if nextEventID < common.FirstEventID {
			break
		}

		// GetWorkflowExecutionHistoryReverse zeros out lastFirstTxnID on paginated calls. It's
		// only populated on the initial call.
		lastFirstTxnID = 0
	}
	return allEvents, nil
}
