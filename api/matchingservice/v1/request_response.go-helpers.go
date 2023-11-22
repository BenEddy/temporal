// The MIT License
//
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.
//
// Copyright (c) 2020 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package matchingservice

import (
	"google.golang.org/protobuf/proto"
)

// Marshal an object of type PollWorkflowTaskQueueRequest to the protobuf v3 wire format
func (val *PollWorkflowTaskQueueRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type PollWorkflowTaskQueueRequest from the protobuf v3 wire format
func (val *PollWorkflowTaskQueueRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *PollWorkflowTaskQueueRequest) Size() int {
	return proto.Size(val)
}

// Equal returns whether two PollWorkflowTaskQueueRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *PollWorkflowTaskQueueRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *PollWorkflowTaskQueueRequest
	switch t := that.(type) {
	case *PollWorkflowTaskQueueRequest:
		that1 = t
	case PollWorkflowTaskQueueRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type PollWorkflowTaskQueueResponse to the protobuf v3 wire format
func (val *PollWorkflowTaskQueueResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type PollWorkflowTaskQueueResponse from the protobuf v3 wire format
func (val *PollWorkflowTaskQueueResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *PollWorkflowTaskQueueResponse) Size() int {
	return proto.Size(val)
}

// Equal returns whether two PollWorkflowTaskQueueResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *PollWorkflowTaskQueueResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *PollWorkflowTaskQueueResponse
	switch t := that.(type) {
	case *PollWorkflowTaskQueueResponse:
		that1 = t
	case PollWorkflowTaskQueueResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type PollActivityTaskQueueRequest to the protobuf v3 wire format
func (val *PollActivityTaskQueueRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type PollActivityTaskQueueRequest from the protobuf v3 wire format
func (val *PollActivityTaskQueueRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *PollActivityTaskQueueRequest) Size() int {
	return proto.Size(val)
}

// Equal returns whether two PollActivityTaskQueueRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *PollActivityTaskQueueRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *PollActivityTaskQueueRequest
	switch t := that.(type) {
	case *PollActivityTaskQueueRequest:
		that1 = t
	case PollActivityTaskQueueRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type PollActivityTaskQueueResponse to the protobuf v3 wire format
func (val *PollActivityTaskQueueResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type PollActivityTaskQueueResponse from the protobuf v3 wire format
func (val *PollActivityTaskQueueResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *PollActivityTaskQueueResponse) Size() int {
	return proto.Size(val)
}

// Equal returns whether two PollActivityTaskQueueResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *PollActivityTaskQueueResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *PollActivityTaskQueueResponse
	switch t := that.(type) {
	case *PollActivityTaskQueueResponse:
		that1 = t
	case PollActivityTaskQueueResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type AddWorkflowTaskRequest to the protobuf v3 wire format
func (val *AddWorkflowTaskRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type AddWorkflowTaskRequest from the protobuf v3 wire format
func (val *AddWorkflowTaskRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *AddWorkflowTaskRequest) Size() int {
	return proto.Size(val)
}

// Equal returns whether two AddWorkflowTaskRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *AddWorkflowTaskRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *AddWorkflowTaskRequest
	switch t := that.(type) {
	case *AddWorkflowTaskRequest:
		that1 = t
	case AddWorkflowTaskRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type AddWorkflowTaskResponse to the protobuf v3 wire format
func (val *AddWorkflowTaskResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type AddWorkflowTaskResponse from the protobuf v3 wire format
func (val *AddWorkflowTaskResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *AddWorkflowTaskResponse) Size() int {
	return proto.Size(val)
}

// Equal returns whether two AddWorkflowTaskResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *AddWorkflowTaskResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *AddWorkflowTaskResponse
	switch t := that.(type) {
	case *AddWorkflowTaskResponse:
		that1 = t
	case AddWorkflowTaskResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type AddActivityTaskRequest to the protobuf v3 wire format
func (val *AddActivityTaskRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type AddActivityTaskRequest from the protobuf v3 wire format
func (val *AddActivityTaskRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *AddActivityTaskRequest) Size() int {
	return proto.Size(val)
}

// Equal returns whether two AddActivityTaskRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *AddActivityTaskRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *AddActivityTaskRequest
	switch t := that.(type) {
	case *AddActivityTaskRequest:
		that1 = t
	case AddActivityTaskRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type AddActivityTaskResponse to the protobuf v3 wire format
func (val *AddActivityTaskResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type AddActivityTaskResponse from the protobuf v3 wire format
func (val *AddActivityTaskResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *AddActivityTaskResponse) Size() int {
	return proto.Size(val)
}

// Equal returns whether two AddActivityTaskResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *AddActivityTaskResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *AddActivityTaskResponse
	switch t := that.(type) {
	case *AddActivityTaskResponse:
		that1 = t
	case AddActivityTaskResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type QueryWorkflowRequest to the protobuf v3 wire format
func (val *QueryWorkflowRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type QueryWorkflowRequest from the protobuf v3 wire format
func (val *QueryWorkflowRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *QueryWorkflowRequest) Size() int {
	return proto.Size(val)
}

// Equal returns whether two QueryWorkflowRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *QueryWorkflowRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *QueryWorkflowRequest
	switch t := that.(type) {
	case *QueryWorkflowRequest:
		that1 = t
	case QueryWorkflowRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type QueryWorkflowResponse to the protobuf v3 wire format
func (val *QueryWorkflowResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type QueryWorkflowResponse from the protobuf v3 wire format
func (val *QueryWorkflowResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *QueryWorkflowResponse) Size() int {
	return proto.Size(val)
}

// Equal returns whether two QueryWorkflowResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *QueryWorkflowResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *QueryWorkflowResponse
	switch t := that.(type) {
	case *QueryWorkflowResponse:
		that1 = t
	case QueryWorkflowResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type RespondQueryTaskCompletedRequest to the protobuf v3 wire format
func (val *RespondQueryTaskCompletedRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type RespondQueryTaskCompletedRequest from the protobuf v3 wire format
func (val *RespondQueryTaskCompletedRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *RespondQueryTaskCompletedRequest) Size() int {
	return proto.Size(val)
}

// Equal returns whether two RespondQueryTaskCompletedRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *RespondQueryTaskCompletedRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *RespondQueryTaskCompletedRequest
	switch t := that.(type) {
	case *RespondQueryTaskCompletedRequest:
		that1 = t
	case RespondQueryTaskCompletedRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type RespondQueryTaskCompletedResponse to the protobuf v3 wire format
func (val *RespondQueryTaskCompletedResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type RespondQueryTaskCompletedResponse from the protobuf v3 wire format
func (val *RespondQueryTaskCompletedResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *RespondQueryTaskCompletedResponse) Size() int {
	return proto.Size(val)
}

// Equal returns whether two RespondQueryTaskCompletedResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *RespondQueryTaskCompletedResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *RespondQueryTaskCompletedResponse
	switch t := that.(type) {
	case *RespondQueryTaskCompletedResponse:
		that1 = t
	case RespondQueryTaskCompletedResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type CancelOutstandingPollRequest to the protobuf v3 wire format
func (val *CancelOutstandingPollRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type CancelOutstandingPollRequest from the protobuf v3 wire format
func (val *CancelOutstandingPollRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *CancelOutstandingPollRequest) Size() int {
	return proto.Size(val)
}

// Equal returns whether two CancelOutstandingPollRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *CancelOutstandingPollRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *CancelOutstandingPollRequest
	switch t := that.(type) {
	case *CancelOutstandingPollRequest:
		that1 = t
	case CancelOutstandingPollRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type CancelOutstandingPollResponse to the protobuf v3 wire format
func (val *CancelOutstandingPollResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type CancelOutstandingPollResponse from the protobuf v3 wire format
func (val *CancelOutstandingPollResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *CancelOutstandingPollResponse) Size() int {
	return proto.Size(val)
}

// Equal returns whether two CancelOutstandingPollResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *CancelOutstandingPollResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *CancelOutstandingPollResponse
	switch t := that.(type) {
	case *CancelOutstandingPollResponse:
		that1 = t
	case CancelOutstandingPollResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type DescribeTaskQueueRequest to the protobuf v3 wire format
func (val *DescribeTaskQueueRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type DescribeTaskQueueRequest from the protobuf v3 wire format
func (val *DescribeTaskQueueRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *DescribeTaskQueueRequest) Size() int {
	return proto.Size(val)
}

// Equal returns whether two DescribeTaskQueueRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *DescribeTaskQueueRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *DescribeTaskQueueRequest
	switch t := that.(type) {
	case *DescribeTaskQueueRequest:
		that1 = t
	case DescribeTaskQueueRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type DescribeTaskQueueResponse to the protobuf v3 wire format
func (val *DescribeTaskQueueResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type DescribeTaskQueueResponse from the protobuf v3 wire format
func (val *DescribeTaskQueueResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *DescribeTaskQueueResponse) Size() int {
	return proto.Size(val)
}

// Equal returns whether two DescribeTaskQueueResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *DescribeTaskQueueResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *DescribeTaskQueueResponse
	switch t := that.(type) {
	case *DescribeTaskQueueResponse:
		that1 = t
	case DescribeTaskQueueResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type ListTaskQueuePartitionsRequest to the protobuf v3 wire format
func (val *ListTaskQueuePartitionsRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type ListTaskQueuePartitionsRequest from the protobuf v3 wire format
func (val *ListTaskQueuePartitionsRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *ListTaskQueuePartitionsRequest) Size() int {
	return proto.Size(val)
}

// Equal returns whether two ListTaskQueuePartitionsRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *ListTaskQueuePartitionsRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *ListTaskQueuePartitionsRequest
	switch t := that.(type) {
	case *ListTaskQueuePartitionsRequest:
		that1 = t
	case ListTaskQueuePartitionsRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type ListTaskQueuePartitionsResponse to the protobuf v3 wire format
func (val *ListTaskQueuePartitionsResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type ListTaskQueuePartitionsResponse from the protobuf v3 wire format
func (val *ListTaskQueuePartitionsResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *ListTaskQueuePartitionsResponse) Size() int {
	return proto.Size(val)
}

// Equal returns whether two ListTaskQueuePartitionsResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *ListTaskQueuePartitionsResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *ListTaskQueuePartitionsResponse
	switch t := that.(type) {
	case *ListTaskQueuePartitionsResponse:
		that1 = t
	case ListTaskQueuePartitionsResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type UpdateWorkerBuildIdCompatibilityRequest to the protobuf v3 wire format
func (val *UpdateWorkerBuildIdCompatibilityRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type UpdateWorkerBuildIdCompatibilityRequest from the protobuf v3 wire format
func (val *UpdateWorkerBuildIdCompatibilityRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *UpdateWorkerBuildIdCompatibilityRequest) Size() int {
	return proto.Size(val)
}

// Equal returns whether two UpdateWorkerBuildIdCompatibilityRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *UpdateWorkerBuildIdCompatibilityRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *UpdateWorkerBuildIdCompatibilityRequest
	switch t := that.(type) {
	case *UpdateWorkerBuildIdCompatibilityRequest:
		that1 = t
	case UpdateWorkerBuildIdCompatibilityRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type UpdateWorkerBuildIdCompatibilityResponse to the protobuf v3 wire format
func (val *UpdateWorkerBuildIdCompatibilityResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type UpdateWorkerBuildIdCompatibilityResponse from the protobuf v3 wire format
func (val *UpdateWorkerBuildIdCompatibilityResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *UpdateWorkerBuildIdCompatibilityResponse) Size() int {
	return proto.Size(val)
}

// Equal returns whether two UpdateWorkerBuildIdCompatibilityResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *UpdateWorkerBuildIdCompatibilityResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *UpdateWorkerBuildIdCompatibilityResponse
	switch t := that.(type) {
	case *UpdateWorkerBuildIdCompatibilityResponse:
		that1 = t
	case UpdateWorkerBuildIdCompatibilityResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type GetWorkerBuildIdCompatibilityRequest to the protobuf v3 wire format
func (val *GetWorkerBuildIdCompatibilityRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type GetWorkerBuildIdCompatibilityRequest from the protobuf v3 wire format
func (val *GetWorkerBuildIdCompatibilityRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *GetWorkerBuildIdCompatibilityRequest) Size() int {
	return proto.Size(val)
}

// Equal returns whether two GetWorkerBuildIdCompatibilityRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *GetWorkerBuildIdCompatibilityRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *GetWorkerBuildIdCompatibilityRequest
	switch t := that.(type) {
	case *GetWorkerBuildIdCompatibilityRequest:
		that1 = t
	case GetWorkerBuildIdCompatibilityRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type GetWorkerBuildIdCompatibilityResponse to the protobuf v3 wire format
func (val *GetWorkerBuildIdCompatibilityResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type GetWorkerBuildIdCompatibilityResponse from the protobuf v3 wire format
func (val *GetWorkerBuildIdCompatibilityResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *GetWorkerBuildIdCompatibilityResponse) Size() int {
	return proto.Size(val)
}

// Equal returns whether two GetWorkerBuildIdCompatibilityResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *GetWorkerBuildIdCompatibilityResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *GetWorkerBuildIdCompatibilityResponse
	switch t := that.(type) {
	case *GetWorkerBuildIdCompatibilityResponse:
		that1 = t
	case GetWorkerBuildIdCompatibilityResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type GetTaskQueueUserDataRequest to the protobuf v3 wire format
func (val *GetTaskQueueUserDataRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type GetTaskQueueUserDataRequest from the protobuf v3 wire format
func (val *GetTaskQueueUserDataRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *GetTaskQueueUserDataRequest) Size() int {
	return proto.Size(val)
}

// Equal returns whether two GetTaskQueueUserDataRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *GetTaskQueueUserDataRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *GetTaskQueueUserDataRequest
	switch t := that.(type) {
	case *GetTaskQueueUserDataRequest:
		that1 = t
	case GetTaskQueueUserDataRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type GetTaskQueueUserDataResponse to the protobuf v3 wire format
func (val *GetTaskQueueUserDataResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type GetTaskQueueUserDataResponse from the protobuf v3 wire format
func (val *GetTaskQueueUserDataResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *GetTaskQueueUserDataResponse) Size() int {
	return proto.Size(val)
}

// Equal returns whether two GetTaskQueueUserDataResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *GetTaskQueueUserDataResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *GetTaskQueueUserDataResponse
	switch t := that.(type) {
	case *GetTaskQueueUserDataResponse:
		that1 = t
	case GetTaskQueueUserDataResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type ApplyTaskQueueUserDataReplicationEventRequest to the protobuf v3 wire format
func (val *ApplyTaskQueueUserDataReplicationEventRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type ApplyTaskQueueUserDataReplicationEventRequest from the protobuf v3 wire format
func (val *ApplyTaskQueueUserDataReplicationEventRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *ApplyTaskQueueUserDataReplicationEventRequest) Size() int {
	return proto.Size(val)
}

// Equal returns whether two ApplyTaskQueueUserDataReplicationEventRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *ApplyTaskQueueUserDataReplicationEventRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *ApplyTaskQueueUserDataReplicationEventRequest
	switch t := that.(type) {
	case *ApplyTaskQueueUserDataReplicationEventRequest:
		that1 = t
	case ApplyTaskQueueUserDataReplicationEventRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type ApplyTaskQueueUserDataReplicationEventResponse to the protobuf v3 wire format
func (val *ApplyTaskQueueUserDataReplicationEventResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type ApplyTaskQueueUserDataReplicationEventResponse from the protobuf v3 wire format
func (val *ApplyTaskQueueUserDataReplicationEventResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *ApplyTaskQueueUserDataReplicationEventResponse) Size() int {
	return proto.Size(val)
}

// Equal returns whether two ApplyTaskQueueUserDataReplicationEventResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *ApplyTaskQueueUserDataReplicationEventResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *ApplyTaskQueueUserDataReplicationEventResponse
	switch t := that.(type) {
	case *ApplyTaskQueueUserDataReplicationEventResponse:
		that1 = t
	case ApplyTaskQueueUserDataReplicationEventResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type GetBuildIdTaskQueueMappingRequest to the protobuf v3 wire format
func (val *GetBuildIdTaskQueueMappingRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type GetBuildIdTaskQueueMappingRequest from the protobuf v3 wire format
func (val *GetBuildIdTaskQueueMappingRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *GetBuildIdTaskQueueMappingRequest) Size() int {
	return proto.Size(val)
}

// Equal returns whether two GetBuildIdTaskQueueMappingRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *GetBuildIdTaskQueueMappingRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *GetBuildIdTaskQueueMappingRequest
	switch t := that.(type) {
	case *GetBuildIdTaskQueueMappingRequest:
		that1 = t
	case GetBuildIdTaskQueueMappingRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type GetBuildIdTaskQueueMappingResponse to the protobuf v3 wire format
func (val *GetBuildIdTaskQueueMappingResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type GetBuildIdTaskQueueMappingResponse from the protobuf v3 wire format
func (val *GetBuildIdTaskQueueMappingResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *GetBuildIdTaskQueueMappingResponse) Size() int {
	return proto.Size(val)
}

// Equal returns whether two GetBuildIdTaskQueueMappingResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *GetBuildIdTaskQueueMappingResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *GetBuildIdTaskQueueMappingResponse
	switch t := that.(type) {
	case *GetBuildIdTaskQueueMappingResponse:
		that1 = t
	case GetBuildIdTaskQueueMappingResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type ForceUnloadTaskQueueRequest to the protobuf v3 wire format
func (val *ForceUnloadTaskQueueRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type ForceUnloadTaskQueueRequest from the protobuf v3 wire format
func (val *ForceUnloadTaskQueueRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *ForceUnloadTaskQueueRequest) Size() int {
	return proto.Size(val)
}

// Equal returns whether two ForceUnloadTaskQueueRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *ForceUnloadTaskQueueRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *ForceUnloadTaskQueueRequest
	switch t := that.(type) {
	case *ForceUnloadTaskQueueRequest:
		that1 = t
	case ForceUnloadTaskQueueRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type ForceUnloadTaskQueueResponse to the protobuf v3 wire format
func (val *ForceUnloadTaskQueueResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type ForceUnloadTaskQueueResponse from the protobuf v3 wire format
func (val *ForceUnloadTaskQueueResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *ForceUnloadTaskQueueResponse) Size() int {
	return proto.Size(val)
}

// Equal returns whether two ForceUnloadTaskQueueResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *ForceUnloadTaskQueueResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *ForceUnloadTaskQueueResponse
	switch t := that.(type) {
	case *ForceUnloadTaskQueueResponse:
		that1 = t
	case ForceUnloadTaskQueueResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type UpdateTaskQueueUserDataRequest to the protobuf v3 wire format
func (val *UpdateTaskQueueUserDataRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type UpdateTaskQueueUserDataRequest from the protobuf v3 wire format
func (val *UpdateTaskQueueUserDataRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *UpdateTaskQueueUserDataRequest) Size() int {
	return proto.Size(val)
}

// Equal returns whether two UpdateTaskQueueUserDataRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *UpdateTaskQueueUserDataRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *UpdateTaskQueueUserDataRequest
	switch t := that.(type) {
	case *UpdateTaskQueueUserDataRequest:
		that1 = t
	case UpdateTaskQueueUserDataRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type UpdateTaskQueueUserDataResponse to the protobuf v3 wire format
func (val *UpdateTaskQueueUserDataResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type UpdateTaskQueueUserDataResponse from the protobuf v3 wire format
func (val *UpdateTaskQueueUserDataResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *UpdateTaskQueueUserDataResponse) Size() int {
	return proto.Size(val)
}

// Equal returns whether two UpdateTaskQueueUserDataResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *UpdateTaskQueueUserDataResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *UpdateTaskQueueUserDataResponse
	switch t := that.(type) {
	case *UpdateTaskQueueUserDataResponse:
		that1 = t
	case UpdateTaskQueueUserDataResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type ReplicateTaskQueueUserDataRequest to the protobuf v3 wire format
func (val *ReplicateTaskQueueUserDataRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type ReplicateTaskQueueUserDataRequest from the protobuf v3 wire format
func (val *ReplicateTaskQueueUserDataRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *ReplicateTaskQueueUserDataRequest) Size() int {
	return proto.Size(val)
}

// Equal returns whether two ReplicateTaskQueueUserDataRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *ReplicateTaskQueueUserDataRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *ReplicateTaskQueueUserDataRequest
	switch t := that.(type) {
	case *ReplicateTaskQueueUserDataRequest:
		that1 = t
	case ReplicateTaskQueueUserDataRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type ReplicateTaskQueueUserDataResponse to the protobuf v3 wire format
func (val *ReplicateTaskQueueUserDataResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type ReplicateTaskQueueUserDataResponse from the protobuf v3 wire format
func (val *ReplicateTaskQueueUserDataResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *ReplicateTaskQueueUserDataResponse) Size() int {
	return proto.Size(val)
}

// Equal returns whether two ReplicateTaskQueueUserDataResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *ReplicateTaskQueueUserDataResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *ReplicateTaskQueueUserDataResponse
	switch t := that.(type) {
	case *ReplicateTaskQueueUserDataResponse:
		that1 = t
	case ReplicateTaskQueueUserDataResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}