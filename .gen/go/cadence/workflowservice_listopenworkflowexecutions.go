// Copyright (c) 2017 Uber Technologies, Inc.
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

// Code generated by thriftrw v1.8.0. DO NOT EDIT.
// @generated

package cadence

import (
	"errors"
	"fmt"
	"github.com/uber/cadence/.gen/go/shared"
	"go.uber.org/thriftrw/wire"
	"strings"
)

// WorkflowService_ListOpenWorkflowExecutions_Args represents the arguments for the WorkflowService.ListOpenWorkflowExecutions function.
//
// The arguments for ListOpenWorkflowExecutions are sent and received over the wire as this struct.
type WorkflowService_ListOpenWorkflowExecutions_Args struct {
	ListRequest *shared.ListOpenWorkflowExecutionsRequest `json:"listRequest,omitempty"`
}

// ToWire translates a WorkflowService_ListOpenWorkflowExecutions_Args struct into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
//
// An error is returned if the struct or any of its fields failed to
// validate.
//
//   x, err := v.ToWire()
//   if err != nil {
//     return err
//   }
//
//   if err := binaryProtocol.Encode(x, writer); err != nil {
//     return err
//   }
func (v *WorkflowService_ListOpenWorkflowExecutions_Args) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	if v.ListRequest != nil {
		w, err = v.ListRequest.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 1, Value: w}
		i++
	}

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _ListOpenWorkflowExecutionsRequest_Read(w wire.Value) (*shared.ListOpenWorkflowExecutionsRequest, error) {
	var v shared.ListOpenWorkflowExecutionsRequest
	err := v.FromWire(w)
	return &v, err
}

// FromWire deserializes a WorkflowService_ListOpenWorkflowExecutions_Args struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a WorkflowService_ListOpenWorkflowExecutions_Args struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v WorkflowService_ListOpenWorkflowExecutions_Args
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *WorkflowService_ListOpenWorkflowExecutions_Args) FromWire(w wire.Value) error {
	var err error

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TStruct {
				v.ListRequest, err = _ListOpenWorkflowExecutionsRequest_Read(field.Value)
				if err != nil {
					return err
				}

			}
		}
	}

	return nil
}

// String returns a readable string representation of a WorkflowService_ListOpenWorkflowExecutions_Args
// struct.
func (v *WorkflowService_ListOpenWorkflowExecutions_Args) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [1]string
	i := 0
	if v.ListRequest != nil {
		fields[i] = fmt.Sprintf("ListRequest: %v", v.ListRequest)
		i++
	}

	return fmt.Sprintf("WorkflowService_ListOpenWorkflowExecutions_Args{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this WorkflowService_ListOpenWorkflowExecutions_Args match the
// provided WorkflowService_ListOpenWorkflowExecutions_Args.
//
// This function performs a deep comparison.
func (v *WorkflowService_ListOpenWorkflowExecutions_Args) Equals(rhs *WorkflowService_ListOpenWorkflowExecutions_Args) bool {
	if !((v.ListRequest == nil && rhs.ListRequest == nil) || (v.ListRequest != nil && rhs.ListRequest != nil && v.ListRequest.Equals(rhs.ListRequest))) {
		return false
	}

	return true
}

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the arguments.
//
// This will always be "ListOpenWorkflowExecutions" for this struct.
func (v *WorkflowService_ListOpenWorkflowExecutions_Args) MethodName() string {
	return "ListOpenWorkflowExecutions"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be Call for this struct.
func (v *WorkflowService_ListOpenWorkflowExecutions_Args) EnvelopeType() wire.EnvelopeType {
	return wire.Call
}

// WorkflowService_ListOpenWorkflowExecutions_Helper provides functions that aid in handling the
// parameters and return values of the WorkflowService.ListOpenWorkflowExecutions
// function.
var WorkflowService_ListOpenWorkflowExecutions_Helper = struct {
	// Args accepts the parameters of ListOpenWorkflowExecutions in-order and returns
	// the arguments struct for the function.
	Args func(
		listRequest *shared.ListOpenWorkflowExecutionsRequest,
	) *WorkflowService_ListOpenWorkflowExecutions_Args

	// IsException returns true if the given error can be thrown
	// by ListOpenWorkflowExecutions.
	//
	// An error can be thrown by ListOpenWorkflowExecutions only if the
	// corresponding exception type was mentioned in the 'throws'
	// section for it in the Thrift file.
	IsException func(error) bool

	// WrapResponse returns the result struct for ListOpenWorkflowExecutions
	// given its return value and error.
	//
	// This allows mapping values and errors returned by
	// ListOpenWorkflowExecutions into a serializable result struct.
	// WrapResponse returns a non-nil error if the provided
	// error cannot be thrown by ListOpenWorkflowExecutions
	//
	//   value, err := ListOpenWorkflowExecutions(args)
	//   result, err := WorkflowService_ListOpenWorkflowExecutions_Helper.WrapResponse(value, err)
	//   if err != nil {
	//     return fmt.Errorf("unexpected error from ListOpenWorkflowExecutions: %v", err)
	//   }
	//   serialize(result)
	WrapResponse func(*shared.ListOpenWorkflowExecutionsResponse, error) (*WorkflowService_ListOpenWorkflowExecutions_Result, error)

	// UnwrapResponse takes the result struct for ListOpenWorkflowExecutions
	// and returns the value or error returned by it.
	//
	// The error is non-nil only if ListOpenWorkflowExecutions threw an
	// exception.
	//
	//   result := deserialize(bytes)
	//   value, err := WorkflowService_ListOpenWorkflowExecutions_Helper.UnwrapResponse(result)
	UnwrapResponse func(*WorkflowService_ListOpenWorkflowExecutions_Result) (*shared.ListOpenWorkflowExecutionsResponse, error)
}{}

func init() {
	WorkflowService_ListOpenWorkflowExecutions_Helper.Args = func(
		listRequest *shared.ListOpenWorkflowExecutionsRequest,
	) *WorkflowService_ListOpenWorkflowExecutions_Args {
		return &WorkflowService_ListOpenWorkflowExecutions_Args{
			ListRequest: listRequest,
		}
	}

	WorkflowService_ListOpenWorkflowExecutions_Helper.IsException = func(err error) bool {
		switch err.(type) {
		case *shared.BadRequestError:
			return true
		case *shared.InternalServiceError:
			return true
		case *shared.EntityNotExistsError:
			return true
		case *shared.ServiceBusyError:
			return true
		default:
			return false
		}
	}

	WorkflowService_ListOpenWorkflowExecutions_Helper.WrapResponse = func(success *shared.ListOpenWorkflowExecutionsResponse, err error) (*WorkflowService_ListOpenWorkflowExecutions_Result, error) {
		if err == nil {
			return &WorkflowService_ListOpenWorkflowExecutions_Result{Success: success}, nil
		}

		switch e := err.(type) {
		case *shared.BadRequestError:
			if e == nil {
				return nil, errors.New("WrapResponse received non-nil error type with nil value for WorkflowService_ListOpenWorkflowExecutions_Result.BadRequestError")
			}
			return &WorkflowService_ListOpenWorkflowExecutions_Result{BadRequestError: e}, nil
		case *shared.InternalServiceError:
			if e == nil {
				return nil, errors.New("WrapResponse received non-nil error type with nil value for WorkflowService_ListOpenWorkflowExecutions_Result.InternalServiceError")
			}
			return &WorkflowService_ListOpenWorkflowExecutions_Result{InternalServiceError: e}, nil
		case *shared.EntityNotExistsError:
			if e == nil {
				return nil, errors.New("WrapResponse received non-nil error type with nil value for WorkflowService_ListOpenWorkflowExecutions_Result.EntityNotExistError")
			}
			return &WorkflowService_ListOpenWorkflowExecutions_Result{EntityNotExistError: e}, nil
		case *shared.ServiceBusyError:
			if e == nil {
				return nil, errors.New("WrapResponse received non-nil error type with nil value for WorkflowService_ListOpenWorkflowExecutions_Result.ServiceBusyError")
			}
			return &WorkflowService_ListOpenWorkflowExecutions_Result{ServiceBusyError: e}, nil
		}

		return nil, err
	}
	WorkflowService_ListOpenWorkflowExecutions_Helper.UnwrapResponse = func(result *WorkflowService_ListOpenWorkflowExecutions_Result) (success *shared.ListOpenWorkflowExecutionsResponse, err error) {
		if result.BadRequestError != nil {
			err = result.BadRequestError
			return
		}
		if result.InternalServiceError != nil {
			err = result.InternalServiceError
			return
		}
		if result.EntityNotExistError != nil {
			err = result.EntityNotExistError
			return
		}
		if result.ServiceBusyError != nil {
			err = result.ServiceBusyError
			return
		}

		if result.Success != nil {
			success = result.Success
			return
		}

		err = errors.New("expected a non-void result")
		return
	}

}

// WorkflowService_ListOpenWorkflowExecutions_Result represents the result of a WorkflowService.ListOpenWorkflowExecutions function call.
//
// The result of a ListOpenWorkflowExecutions execution is sent and received over the wire as this struct.
//
// Success is set only if the function did not throw an exception.
type WorkflowService_ListOpenWorkflowExecutions_Result struct {
	// Value returned by ListOpenWorkflowExecutions after a successful execution.
	Success              *shared.ListOpenWorkflowExecutionsResponse `json:"success,omitempty"`
	BadRequestError      *shared.BadRequestError                    `json:"badRequestError,omitempty"`
	InternalServiceError *shared.InternalServiceError               `json:"internalServiceError,omitempty"`
	EntityNotExistError  *shared.EntityNotExistsError               `json:"entityNotExistError,omitempty"`
	ServiceBusyError     *shared.ServiceBusyError                   `json:"serviceBusyError,omitempty"`
}

// ToWire translates a WorkflowService_ListOpenWorkflowExecutions_Result struct into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
//
// An error is returned if the struct or any of its fields failed to
// validate.
//
//   x, err := v.ToWire()
//   if err != nil {
//     return err
//   }
//
//   if err := binaryProtocol.Encode(x, writer); err != nil {
//     return err
//   }
func (v *WorkflowService_ListOpenWorkflowExecutions_Result) ToWire() (wire.Value, error) {
	var (
		fields [5]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	if v.Success != nil {
		w, err = v.Success.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 0, Value: w}
		i++
	}
	if v.BadRequestError != nil {
		w, err = v.BadRequestError.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 1, Value: w}
		i++
	}
	if v.InternalServiceError != nil {
		w, err = v.InternalServiceError.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 2, Value: w}
		i++
	}
	if v.EntityNotExistError != nil {
		w, err = v.EntityNotExistError.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 3, Value: w}
		i++
	}
	if v.ServiceBusyError != nil {
		w, err = v.ServiceBusyError.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 4, Value: w}
		i++
	}

	if i != 1 {
		return wire.Value{}, fmt.Errorf("WorkflowService_ListOpenWorkflowExecutions_Result should have exactly one field: got %v fields", i)
	}

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _ListOpenWorkflowExecutionsResponse_Read(w wire.Value) (*shared.ListOpenWorkflowExecutionsResponse, error) {
	var v shared.ListOpenWorkflowExecutionsResponse
	err := v.FromWire(w)
	return &v, err
}

// FromWire deserializes a WorkflowService_ListOpenWorkflowExecutions_Result struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a WorkflowService_ListOpenWorkflowExecutions_Result struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v WorkflowService_ListOpenWorkflowExecutions_Result
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *WorkflowService_ListOpenWorkflowExecutions_Result) FromWire(w wire.Value) error {
	var err error

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 0:
			if field.Value.Type() == wire.TStruct {
				v.Success, err = _ListOpenWorkflowExecutionsResponse_Read(field.Value)
				if err != nil {
					return err
				}

			}
		case 1:
			if field.Value.Type() == wire.TStruct {
				v.BadRequestError, err = _BadRequestError_Read(field.Value)
				if err != nil {
					return err
				}

			}
		case 2:
			if field.Value.Type() == wire.TStruct {
				v.InternalServiceError, err = _InternalServiceError_Read(field.Value)
				if err != nil {
					return err
				}

			}
		case 3:
			if field.Value.Type() == wire.TStruct {
				v.EntityNotExistError, err = _EntityNotExistsError_Read(field.Value)
				if err != nil {
					return err
				}

			}
		case 4:
			if field.Value.Type() == wire.TStruct {
				v.ServiceBusyError, err = _ServiceBusyError_Read(field.Value)
				if err != nil {
					return err
				}

			}
		}
	}

	count := 0
	if v.Success != nil {
		count++
	}
	if v.BadRequestError != nil {
		count++
	}
	if v.InternalServiceError != nil {
		count++
	}
	if v.EntityNotExistError != nil {
		count++
	}
	if v.ServiceBusyError != nil {
		count++
	}
	if count != 1 {
		return fmt.Errorf("WorkflowService_ListOpenWorkflowExecutions_Result should have exactly one field: got %v fields", count)
	}

	return nil
}

// String returns a readable string representation of a WorkflowService_ListOpenWorkflowExecutions_Result
// struct.
func (v *WorkflowService_ListOpenWorkflowExecutions_Result) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [5]string
	i := 0
	if v.Success != nil {
		fields[i] = fmt.Sprintf("Success: %v", v.Success)
		i++
	}
	if v.BadRequestError != nil {
		fields[i] = fmt.Sprintf("BadRequestError: %v", v.BadRequestError)
		i++
	}
	if v.InternalServiceError != nil {
		fields[i] = fmt.Sprintf("InternalServiceError: %v", v.InternalServiceError)
		i++
	}
	if v.EntityNotExistError != nil {
		fields[i] = fmt.Sprintf("EntityNotExistError: %v", v.EntityNotExistError)
		i++
	}
	if v.ServiceBusyError != nil {
		fields[i] = fmt.Sprintf("ServiceBusyError: %v", v.ServiceBusyError)
		i++
	}

	return fmt.Sprintf("WorkflowService_ListOpenWorkflowExecutions_Result{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this WorkflowService_ListOpenWorkflowExecutions_Result match the
// provided WorkflowService_ListOpenWorkflowExecutions_Result.
//
// This function performs a deep comparison.
func (v *WorkflowService_ListOpenWorkflowExecutions_Result) Equals(rhs *WorkflowService_ListOpenWorkflowExecutions_Result) bool {
	if !((v.Success == nil && rhs.Success == nil) || (v.Success != nil && rhs.Success != nil && v.Success.Equals(rhs.Success))) {
		return false
	}
	if !((v.BadRequestError == nil && rhs.BadRequestError == nil) || (v.BadRequestError != nil && rhs.BadRequestError != nil && v.BadRequestError.Equals(rhs.BadRequestError))) {
		return false
	}
	if !((v.InternalServiceError == nil && rhs.InternalServiceError == nil) || (v.InternalServiceError != nil && rhs.InternalServiceError != nil && v.InternalServiceError.Equals(rhs.InternalServiceError))) {
		return false
	}
	if !((v.EntityNotExistError == nil && rhs.EntityNotExistError == nil) || (v.EntityNotExistError != nil && rhs.EntityNotExistError != nil && v.EntityNotExistError.Equals(rhs.EntityNotExistError))) {
		return false
	}
	if !((v.ServiceBusyError == nil && rhs.ServiceBusyError == nil) || (v.ServiceBusyError != nil && rhs.ServiceBusyError != nil && v.ServiceBusyError.Equals(rhs.ServiceBusyError))) {
		return false
	}

	return true
}

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the result.
//
// This will always be "ListOpenWorkflowExecutions" for this struct.
func (v *WorkflowService_ListOpenWorkflowExecutions_Result) MethodName() string {
	return "ListOpenWorkflowExecutions"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be Reply for this struct.
func (v *WorkflowService_ListOpenWorkflowExecutions_Result) EnvelopeType() wire.EnvelopeType {
	return wire.Reply
}
