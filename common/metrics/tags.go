package metrics

import (
	"fmt"
	"strconv"
	"strings"

	enumspb "go.temporal.io/api/enums/v1"
	historypb "go.temporal.io/api/history/v1"
	enumsspb "go.temporal.io/server/api/enums/v1"
	"go.temporal.io/server/common/locks"
	"go.temporal.io/server/common/primitives"
	"go.temporal.io/server/common/util"
)

const (
	gitRevisionTag   = "git_revision"
	buildDateTag     = "build_date"
	buildVersionTag  = "build_version"
	buildPlatformTag = "build_platform"
	goVersionTag     = "go_version"

	instance       = "instance"
	namespace      = "namespace"
	namespaceID    = "namespace_id"
	namespaceState = "namespace_state"
	sourceCluster  = "source_cluster"
	targetCluster  = "target_cluster"
	fromCluster    = "from_cluster"
	toCluster      = "to_cluster"
	taskQueue      = "taskqueue"
	workflowType   = "workflowType"
	activityType   = "activityType"
	commandType    = "commandType"
	serviceName    = "service_name"
	actionType     = "action_type"
	workerBuildId  = "worker-build-id"
	destination    = "destination"
	// Generic reason tag can be used anywhere a reason is needed.
	reason = "reason"
	// See server.api.enums.v1.ReplicationTaskType
	replicationTaskType     = "replicationTaskType"
	replicationTaskPriority = "replicationTaskPriority"
	taskExpireStage         = "task_expire_stage"
	versioningBehavior      = "versioning_behavior"
	isFirstAttempt          = "first-attempt"
	workflowStatus          = "workflow_status"
	behaviorBefore          = "behavior_before"
	behaviorAfter           = "behavior_after"
	runInitiator            = "run_initiator"
	fromUnversioned         = "from_unversioned"
	toUnversioned           = "to_unversioned"
	queryTypeTag            = "query_type"
	namespaceAllValue       = "all"
	unknownValue            = "_unknown_"
	totalMetricSuffix       = "_total"
	tagExcludedValue        = "_tag_excluded_"
	falseValue              = "false"
	trueValue               = "true"
	errorPrefix             = "*"

	queryTypeStackTrace       = "__stack_trace"
	queryTypeOpenSessions     = "__open_sessions"
	queryTypeWorkflowMetadata = "__temporal_workflow_metadata"
	queryTypeUserDefined      = "__user_defined"

	newRun      = "new"
	existingRun = "existing"
	childRun    = "child"
	canRun      = "can"
	retryRun    = "retry"
	cronRun     = "cron"
	unknownRun  = "unknown"
)

// Tag is an interface to define metrics tags
type Tag struct {
	key   string
	value string
}

func newTagImpl(key, value string) Tag {
	return Tag{
		key:   key,
		value: value,
	}
}

func (v Tag) Key() string {
	return v.key
}

func (v Tag) Value() string {
	return v.value
}

func (v Tag) String() string {
	return fmt.Sprintf("tag{key: %q, value: %q}", v.Key(), v.Value())
}

// NamespaceTag returns a new namespace tag. For timers, this also ensures that we
// dual emit the metric with the all tag. If a blank namespace is provided then
// this converts that to an unknown namespace.
func NamespaceTag(value string) Tag {
	if len(value) == 0 {
		value = unknownValue
	}
	return newTagImpl(namespace, value)
}

// NamespaceIDTag returns a new namespace ID tag.
func NamespaceIDTag(value string) Tag {
	if len(value) == 0 {
		value = unknownValue
	}
	return newTagImpl(namespaceID, value)
}

var namespaceUnknownTag = newTagImpl(namespace, unknownValue)

// NamespaceUnknownTag returns a new namespace:unknown tag-value
func NamespaceUnknownTag() Tag {
	return namespaceUnknownTag
}

// NamespaceStateTag returns a new namespace state tag.
func NamespaceStateTag(value string) Tag {
	if len(value) == 0 {
		value = unknownValue
	}
	return newTagImpl(namespaceState, value)
}

var taskQueueUnknownTag = newTagImpl(taskQueue, unknownValue)

// TaskQueueUnknownTag returns a new taskqueue:unknown tag-value
func TaskQueueUnknownTag() Tag {
	return taskQueueUnknownTag
}

// InstanceTag returns a new instance tag
func InstanceTag(value string) Tag {
	return newTagImpl(instance, value)
}

// SourceClusterTag returns a new source cluster tag.
func SourceClusterTag(value string) Tag {
	if len(value) == 0 {
		value = unknownValue
	}
	return newTagImpl(sourceCluster, value)
}

// TargetClusterTag returns a new target cluster tag.
func TargetClusterTag(value string) Tag {
	if len(value) == 0 {
		value = unknownValue
	}
	return newTagImpl(targetCluster, value)
}

// FromClusterIDTag returns a new from cluster tag.
func FromClusterIDTag(value int32) Tag {
	return newTagImpl(fromCluster, strconv.FormatInt(int64(value), 10))
}

// ToClusterIDTag returns a new to cluster tag.
func ToClusterIDTag(value int32) Tag {
	return newTagImpl(toCluster, strconv.FormatInt(int64(value), 10))
}

// UnsafeTaskQueueTag returns a new task queue tag.
// WARNING: Do not use this function directly in production code as it may create high number of unique task queue tag
// values that can trouble the observability stack. Instead, use one of the following helper functions and pass a proper
// breakdown boolean (typically based on the task queue dynamic configs):
// - `workflow.PerTaskQueueFamilyScope`
// - `tqid.PerTaskQueueFamilyScope`
// - `tqid.PerTaskQueueScope`
// - `tqid.PerTaskQueuePartitionScope`
func UnsafeTaskQueueTag(value string) Tag {
	if len(value) == 0 {
		value = unknownValue
	}
	return newTagImpl(taskQueue, value)
}

func TaskQueueTypeTag(tqType enumspb.TaskQueueType) Tag {
	return newTagImpl(TaskTypeTagName, tqType.String())
}

// Consider passing the value of "metrics.breakdownByBuildID" dynamic config to this function.
func WorkerBuildIdTag(buildId string, buildIdBreakdown bool) Tag {
	if buildId == "" {
		buildId = "__unversioned__"
	} else if !buildIdBreakdown {
		buildId = "__versioned__"
	}
	return newTagImpl(workerBuildId, buildId)
}

// WorkflowTypeTag returns a new workflow type tag.
func WorkflowTypeTag(value string) Tag {
	if len(value) == 0 {
		value = unknownValue
	}
	return newTagImpl(workflowType, value)
}

// ActivityTypeTag returns a new activity type tag.
func ActivityTypeTag(value string) Tag {
	if len(value) == 0 {
		value = unknownValue
	}
	return newTagImpl(activityType, value)
}

// CommandTypeTag returns a new command type tag.
func CommandTypeTag(value string) Tag {
	if len(value) == 0 {
		value = unknownValue
	}
	return newTagImpl(commandType, value)
}

// Returns a new service role tag.
func ServiceRoleTag(value string) Tag {
	if len(value) == 0 {
		value = unknownValue
	}
	return newTagImpl(ServiceRoleTagName, value)
}

// Returns a new failure type tag
func FailureTag(value string) Tag {
	if len(value) == 0 {
		value = unknownValue
	}
	return newTagImpl(FailureTagName, value)
}

func FirstAttemptTag(attempt int32) Tag {
	value := falseValue
	if attempt == 1 {
		value = trueValue
	}
	return newTagImpl(isFirstAttempt, value)
}

func FailureSourceTag(value string) Tag {
	if len(value) == 0 {
		value = unknownValue
	}
	return newTagImpl(FailureSourceTagName, value)
}

func TaskCategoryTag(value string) Tag {
	if len(value) == 0 {
		value = unknownValue
	}
	return newTagImpl(TaskCategoryTagName, value)
}

func TaskTypeTag(value string) Tag {
	if len(value) == 0 {
		value = unknownValue
	}
	return newTagImpl(TaskTypeTagName, value)
}

func PartitionTag(partition string) Tag {
	return newTagImpl(PartitionTagName, partition)
}

func TaskPriorityTag(value string) Tag {
	if len(value) == 0 {
		value = unknownValue
	}
	return newTagImpl(TaskPriorityTagName, value)
}

func QueueReaderIDTag(readerID int64) Tag {
	return newTagImpl(QueueReaderIDTagName, strconv.Itoa(int(readerID)))
}

func QueueActionTag(value string) Tag {
	if len(value) == 0 {
		value = unknownValue
	}
	return newTagImpl(QueueActionTagName, value)
}

func QueueTypeTag(value string) Tag {
	if len(value) == 0 {
		value = unknownValue
	}
	return newTagImpl(QueueTypeTagName, value)
}

func VisibilityPluginNameTag(value string) Tag {
	if value == "" {
		value = unknownValue
	}
	return newTagImpl(visibilityPluginNameTagName, value)
}

func VisibilityIndexNameTag(value string) Tag {
	if value == "" {
		value = unknownValue
	}
	return newTagImpl(visibilityIndexNameTagName, value)
}

// VersionedTag represents whether a loaded task queue manager represents a specific version set or build ID or not.
func VersionedTag(versioned string) Tag {
	return newTagImpl(versionedTagName, versioned)
}

func ServiceErrorTypeTag(err error) Tag {
	return newTagImpl(ErrorTypeTagName, strings.TrimPrefix(util.ErrorType(err), errorPrefix))
}

func OutcomeTag(outcome string) Tag {
	return newTagImpl(outcomeTagName, outcome)
}

func NexusMethodTag(value string) Tag {
	return newTagImpl(nexusMethodTagName, value)
}

func NexusEndpointTag(value string) Tag {
	if len(value) == 0 {
		value = unknownValue
	}
	return newTagImpl(nexusEndpointTagName, value)
}

func NexusServiceTag(value string) Tag {
	return newTagImpl(nexusServiceTagName, value)
}

func NexusOperationTag(value string) Tag {
	return newTagImpl(nexusOperationTagName, value)
}

// HttpStatusTag returns a new httpStatusTag.
func HttpStatusTag(value int) Tag {
	return newTagImpl(httpStatusTagName, strconv.Itoa(value))
}

func ResourceExhaustedCauseTag(cause enumspb.ResourceExhaustedCause) Tag {
	return newTagImpl(resourceExhaustedTag, cause.String())
}

func ResourceExhaustedScopeTag(scope enumspb.ResourceExhaustedScope) Tag {
	return newTagImpl(resourceExhaustedScopeTag, scope.String())
}

func ServiceNameTag(value primitives.ServiceName) Tag {
	return newTagImpl(serviceName, string(value))
}

func ActionType(value string) Tag {
	return newTagImpl(actionType, value)
}

func OperationTag(value string) Tag {
	return newTagImpl(OperationTagName, value)
}

func StringTag(key string, value string) Tag {
	return newTagImpl(key, value)
}

func CacheTypeTag(value string) Tag {
	return newTagImpl(CacheTypeTagName, value)
}

func PriorityTag(value locks.Priority) Tag {
	return newTagImpl(PriorityTagName, strconv.Itoa(int(value)))
}

// ReasonString is just a string but the special type is defined here to remind callers of ReasonTag to limit the
// cardinality of possible reasons.
type ReasonString string

// ReasonTag is a generic tag can be used anywhere a reason is needed.
// Make sure that the value is of limited cardinality.
func ReasonTag(value ReasonString) Tag {
	return newTagImpl(reason, string(value))
}

// ReplicationTaskTypeTag returns a new replication task type tag.
func ReplicationTaskTypeTag(value enumsspb.ReplicationTaskType) Tag {
	return newTagImpl(replicationTaskType, value.String())
}

// ReplicationTaskPriorityTag returns a replication task priority tag.
func ReplicationTaskPriorityTag(value enumsspb.TaskPriority) Tag {
	return newTagImpl(replicationTaskPriority, value.String())
}

// DestinationTag is a tag for metrics emitted by outbound task executors for the task's destination.
func DestinationTag(value string) Tag {
	return newTagImpl(destination, value)
}

func VersioningBehaviorTag(behavior enumspb.VersioningBehavior) Tag {
	return newTagImpl(versioningBehavior, behavior.String())
}

func WorkflowStatusTag(status string) Tag {
	return newTagImpl(workflowStatus, status)
}

func QueryTypeTag(queryType string) Tag {
	if queryType == queryTypeStackTrace || queryType == queryTypeOpenSessions || queryType == queryTypeWorkflowMetadata {
		return newTagImpl(queryTypeTag, queryType)
	}
	// group all user defined queries into a single tag value
	return newTagImpl(queryTypeTag, queryTypeUserDefined)
}

func VersioningBehaviorBeforeOverrideTag(behavior enumspb.VersioningBehavior) Tag {
	return newTagImpl(behaviorBefore, behavior.String())
}

func VersioningBehaviorAfterOverrideTag(behavior enumspb.VersioningBehavior) Tag {
	return newTagImpl(behaviorAfter, behavior.String())
}

// RunInitiatorTag creates a tag indicating how a workflow run was initiated.
// It handles both new workflow runs and continuations from previous runs.
// When attributes is nil (e.g. during AddWorkflowExecutionOptionsUpdatedEvent),
// it returns a tag indicating an existing run.
func RunInitiatorTag(prevRunID string, attributes *historypb.WorkflowExecutionStartedEventAttributes) Tag {
	if attributes == nil {
		return newTagImpl(runInitiator, existingRun)
	} else if attributes.GetParentWorkflowExecution() != nil {
		return newTagImpl(runInitiator, childRun)
	}

	switch attributes.GetInitiator() {
	case enumspb.CONTINUE_AS_NEW_INITIATOR_UNSPECIFIED:
		return newTagImpl(runInitiator, newRun)
	case enumspb.CONTINUE_AS_NEW_INITIATOR_WORKFLOW:
		return newTagImpl(runInitiator, canRun)
	case enumspb.CONTINUE_AS_NEW_INITIATOR_RETRY:
		return newTagImpl(runInitiator, retryRun)
	case enumspb.CONTINUE_AS_NEW_INITIATOR_CRON_SCHEDULE:
		return newTagImpl(runInitiator, cronRun)
	default:
		return newTagImpl(runInitiator, unknownRun)
	}
}

func FromUnversionedTag(version string) Tag {
	if version == "_unversioned_" {
		return newTagImpl(fromUnversioned, trueValue)
	}
	return newTagImpl(fromUnversioned, falseValue)
}

func ToUnversionedTag(version string) Tag {
	if version == "_unversioned_" {
		return newTagImpl(toUnversioned, trueValue)
	}
	return newTagImpl(toUnversioned, falseValue)
}

var TaskExpireStageReadTag Tag = newTagImpl(taskExpireStage, "read")
var TaskExpireStageMemoryTag Tag = newTagImpl(taskExpireStage, "memory")
var TaskInvalidTag Tag = newTagImpl(taskExpireStage, "invalid")

func PersistenceDBKindTag(kind string) Tag {
	return newTagImpl(PersistenceDBKindTagName, kind)
}
