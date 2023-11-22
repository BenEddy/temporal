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

package persistence

import (
	"google.golang.org/protobuf/proto"
)

// Marshal an object of type Predicate to the protobuf v3 wire format
func (val *Predicate) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type Predicate from the protobuf v3 wire format
func (val *Predicate) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *Predicate) Size() int {
	return proto.Size(val)
}

// Equal returns whether two Predicate values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *Predicate) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *Predicate
	switch t := that.(type) {
	case *Predicate:
		that1 = t
	case Predicate:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type UniversalPredicateAttributes to the protobuf v3 wire format
func (val *UniversalPredicateAttributes) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type UniversalPredicateAttributes from the protobuf v3 wire format
func (val *UniversalPredicateAttributes) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *UniversalPredicateAttributes) Size() int {
	return proto.Size(val)
}

// Equal returns whether two UniversalPredicateAttributes values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *UniversalPredicateAttributes) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *UniversalPredicateAttributes
	switch t := that.(type) {
	case *UniversalPredicateAttributes:
		that1 = t
	case UniversalPredicateAttributes:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type EmptyPredicateAttributes to the protobuf v3 wire format
func (val *EmptyPredicateAttributes) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type EmptyPredicateAttributes from the protobuf v3 wire format
func (val *EmptyPredicateAttributes) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *EmptyPredicateAttributes) Size() int {
	return proto.Size(val)
}

// Equal returns whether two EmptyPredicateAttributes values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *EmptyPredicateAttributes) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *EmptyPredicateAttributes
	switch t := that.(type) {
	case *EmptyPredicateAttributes:
		that1 = t
	case EmptyPredicateAttributes:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type AndPredicateAttributes to the protobuf v3 wire format
func (val *AndPredicateAttributes) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type AndPredicateAttributes from the protobuf v3 wire format
func (val *AndPredicateAttributes) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *AndPredicateAttributes) Size() int {
	return proto.Size(val)
}

// Equal returns whether two AndPredicateAttributes values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *AndPredicateAttributes) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *AndPredicateAttributes
	switch t := that.(type) {
	case *AndPredicateAttributes:
		that1 = t
	case AndPredicateAttributes:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type OrPredicateAttributes to the protobuf v3 wire format
func (val *OrPredicateAttributes) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type OrPredicateAttributes from the protobuf v3 wire format
func (val *OrPredicateAttributes) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *OrPredicateAttributes) Size() int {
	return proto.Size(val)
}

// Equal returns whether two OrPredicateAttributes values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *OrPredicateAttributes) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *OrPredicateAttributes
	switch t := that.(type) {
	case *OrPredicateAttributes:
		that1 = t
	case OrPredicateAttributes:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type NotPredicateAttributes to the protobuf v3 wire format
func (val *NotPredicateAttributes) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type NotPredicateAttributes from the protobuf v3 wire format
func (val *NotPredicateAttributes) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *NotPredicateAttributes) Size() int {
	return proto.Size(val)
}

// Equal returns whether two NotPredicateAttributes values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *NotPredicateAttributes) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *NotPredicateAttributes
	switch t := that.(type) {
	case *NotPredicateAttributes:
		that1 = t
	case NotPredicateAttributes:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type NamespaceIdPredicateAttributes to the protobuf v3 wire format
func (val *NamespaceIdPredicateAttributes) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type NamespaceIdPredicateAttributes from the protobuf v3 wire format
func (val *NamespaceIdPredicateAttributes) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *NamespaceIdPredicateAttributes) Size() int {
	return proto.Size(val)
}

// Equal returns whether two NamespaceIdPredicateAttributes values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *NamespaceIdPredicateAttributes) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *NamespaceIdPredicateAttributes
	switch t := that.(type) {
	case *NamespaceIdPredicateAttributes:
		that1 = t
	case NamespaceIdPredicateAttributes:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type TaskTypePredicateAttributes to the protobuf v3 wire format
func (val *TaskTypePredicateAttributes) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type TaskTypePredicateAttributes from the protobuf v3 wire format
func (val *TaskTypePredicateAttributes) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *TaskTypePredicateAttributes) Size() int {
	return proto.Size(val)
}

// Equal returns whether two TaskTypePredicateAttributes values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *TaskTypePredicateAttributes) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *TaskTypePredicateAttributes
	switch t := that.(type) {
	case *TaskTypePredicateAttributes:
		that1 = t
	case TaskTypePredicateAttributes:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}