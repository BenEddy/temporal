// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package list

import (
	"testing"
	"unsafe"
)

func checkListLen(t *testing.T, l *WeakList[entry], len int) bool {
	if n := l.Len(); n != len {
		t.Errorf("l.Len() = %d, want %d", n, len)
		return false
	}
	return true
}

func toPointer(p *Element[entry]) uintptr {
	return uintptr(unsafe.Pointer(p))
}

func checkListPointers(t *testing.T, l *WeakList[entry], es []*Element[entry]) {
	root := &l.root

	if !checkListLen(t, l, len(es)) {
		return
	}

	// zero length lists must be the zero value or properly initialized (sentinel circle)
	if len(es) == 0 {
		if l.root._next != 0 && l.root._next != toPointer(root) || l.root._prev != 0 && l.root._prev != toPointer(root) {
			t.Errorf("l.root.next = %d, l.root.prev = %d; both should both be nil or %p", l.root._next, l.root._prev, root)
		}
		return
	}
	// len(es) > 0

	// check internal and external prev/next connections
	for i, e := range es {
		prev := root
		Prev := (*Element[entry])(nil)
		if i > 0 {
			prev = es[i-1]
			Prev = prev
		}
		if p := e._prev; p != toPointer(prev) {
			t.Errorf("elt[%d](%p).prev = %d, want %p", i, e, p, prev)
		}
		if p := e.Prev(); p != Prev {
			t.Errorf("elt[%d](%p).Prev() = %p, want %p", i, e, p, Prev)
		}

		next := root
		Next := (*Element[entry])(nil)
		if i < len(es)-1 {
			next = es[i+1]
			Next = next
		}
		if n := e._next; n != toPointer(next) {
			t.Errorf("elt[%d](%p).next = %d, want %p", i, e, n, next)
		}
		if n := e.Next(); n != Next {
			t.Errorf("elt[%d](%p).Next() = %p, want %p", i, e, n, Next)
		}
	}
}

func TestList(t *testing.T) {
	l := New[entry]()
	checkListPointers(t, l, []*Element[entry]{})

	// Single element list
	e := l.PushFront(&entry{value: 0})
	checkListPointers(t, l, []*Element[entry]{e})
	l.MoveToFront(e)
	checkListPointers(t, l, []*Element[entry]{e})
	l.MoveToBack(e)
	checkListPointers(t, l, []*Element[entry]{e})
	l.Remove(e)
	checkListPointers(t, l, []*Element[entry]{})

	// Bigger list
	e2 := l.PushFront(&entry{value: 2})
	e1 := l.PushFront(&entry{value: 1})
	e3 := l.PushBack(&entry{value: 3})
	e4 := l.PushBack(&entry{value: 0})
	checkListPointers(t, l, []*Element[entry]{e1, e2, e3, e4})

	l.Remove(e2)
	checkListPointers(t, l, []*Element[entry]{e1, e3, e4})

	l.MoveToFront(e3) // move from middle
	checkListPointers(t, l, []*Element[entry]{e3, e1, e4})

	l.MoveToFront(e1)
	l.MoveToBack(e3) // move from middle
	checkListPointers(t, l, []*Element[entry]{e1, e4, e3})

	l.MoveToFront(e3) // move from back
	checkListPointers(t, l, []*Element[entry]{e3, e1, e4})
	l.MoveToFront(e3) // should be no-op
	checkListPointers(t, l, []*Element[entry]{e3, e1, e4})

	l.MoveToBack(e3) // move from front
	checkListPointers(t, l, []*Element[entry]{e1, e4, e3})
	l.MoveToBack(e3) // should be no-op
	checkListPointers(t, l, []*Element[entry]{e1, e4, e3})

	e2 = l.InsertBefore(&entry{value: 2}, e1) // insert before front
	checkListPointers(t, l, []*Element[entry]{e2, e1, e4, e3})
	l.Remove(e2)
	e2 = l.InsertBefore(&entry{value: 2}, e4) // insert before middle
	checkListPointers(t, l, []*Element[entry]{e1, e2, e4, e3})
	l.Remove(e2)
	e2 = l.InsertBefore(&entry{value: 2}, e3) // insert before back
	checkListPointers(t, l, []*Element[entry]{e1, e4, e2, e3})
	l.Remove(e2)

	e2 = l.InsertAfter(&entry{value: 2}, e1) // insert after front
	checkListPointers(t, l, []*Element[entry]{e1, e2, e4, e3})
	l.Remove(e2)
	e2 = l.InsertAfter(&entry{value: 2}, e4) // insert after middle
	checkListPointers(t, l, []*Element[entry]{e1, e4, e2, e3})
	l.Remove(e2)
	e2 = l.InsertAfter(&entry{value: 2}, e3) // insert after back
	checkListPointers(t, l, []*Element[entry]{e1, e4, e3, e2})
	l.Remove(e2)

	// Check standard iteration.
	sum := 0
	for e := l.Front(); e != nil; e = e.Next() {
		sum += e.Value().value
	}
	if sum != 4 {
		t.Errorf("sum over l = %d, want 4", sum)
	}

	// Clear all elements by iterating
	var next *Element[entry]
	for e := l.Front(); e != nil; e = next {
		next = e.Next()
		l.Remove(e)
	}
	checkListPointers(t, l, []*Element[entry]{})
}

func checkList(t *testing.T, l *WeakList[entry], es []any) {
	if !checkListLen(t, l, len(es)) {
		return
	}

	i := 0
	for e := l.Front(); e != nil; e = e.Next() {
		le := e.Value()
		if le.value != es[i] {
			t.Errorf("elt[%d].Value = %v, want %v", i, le, es[i])
		}
		i++
	}
}

func PushBackList(source, dest *WeakList[entry]) {
	for i, e := source.Len(), source.Front(); i > 0; i, e = i-1, e.Next() {
		dest.PushBack(e.Value())
	}
}

func PushFrontList(source, dest *WeakList[entry]) {
	for i, e := source.Len(), source.Back(); i > 0; i, e = i-1, e.Prev() {
		dest.PushFront(e.Value())
	}
}

type (
	entry struct {
		value int
	}

	cache struct {
		entries map[string]*entry
	}
)

func (c *cache) put(key string, value int) *entry {
	e := &entry{value: value}
	c.entries[key] = e
	return e
}

func TestExtending(t *testing.T) {
	l1 := New[entry]()
	l2 := New[entry]()

	l1.PushBack(&entry{value: 1})
	l1.PushBack(&entry{value: 2})
	l1.PushBack(&entry{value: 3})

	l2.PushBack(&entry{value: 4})
	l2.PushBack(&entry{value: 5})

	l3 := New[entry]()

	PushBackList(l1, l3)
	checkList(t, l3, []any{1, 2, 3})

	PushBackList(l2, l3)
	checkList(t, l3, []any{1, 2, 3, 4, 5})

	l3 = New[entry]()
	PushFrontList(l2, l3)
	checkList(t, l3, []any{4, 5})
	PushFrontList(l1, l3)
	checkList(t, l3, []any{1, 2, 3, 4, 5})

	checkList(t, l1, []any{1, 2, 3})
	checkList(t, l2, []any{4, 5})

	l3 = New[entry]()
	PushBackList(l1, l3)
	checkList(t, l3, []any{1, 2, 3})
	PushBackList(l3, l3)
	checkList(t, l3, []any{1, 2, 3, 1, 2, 3})

	l3 = New[entry]()
	PushFrontList(l1, l3)
	checkList(t, l3, []any{1, 2, 3})
	PushFrontList(l3, l3)
	checkList(t, l3, []any{1, 2, 3, 1, 2, 3})

	l3 = New[entry]()
	PushBackList(l3, l1)
	checkList(t, l1, []any{1, 2, 3})
	PushFrontList(l3, l1)
	checkList(t, l1, []any{1, 2, 3})
}

func TestRemove(t *testing.T) {
	l := New[entry]()
	e1 := l.PushBack(&entry{value: 1})
	e2 := l.PushBack(&entry{value: 2})
	checkListPointers(t, l, []*Element[entry]{e1, e2})
	e := l.Front()
	l.Remove(e)
	checkListPointers(t, l, []*Element[entry]{e2})
	l.Remove(e)
	checkListPointers(t, l, []*Element[entry]{e2})
}

func TestIssue4103(t *testing.T) {
	l1 := New[entry]()
	l1.PushBack(&entry{value: 1})
	l1.PushBack(&entry{value: 2})

	l2 := New[entry]()
	l2.PushBack(&entry{value: 3})
	l2.PushBack(&entry{value: 4})

	e := l1.Front()
	l2.Remove(e) // l2 should not change because e is not an element of l2
	if n := l2.Len(); n != 2 {
		t.Errorf("l2.Len() = %d, want 2", n)
	}

	l1.InsertBefore(&entry{value: 8}, e)
	if n := l1.Len(); n != 3 {
		t.Errorf("l1.Len() = %d, want 3", n)
	}
}

func TestIssue6349(t *testing.T) {
	l := New[entry]()
	l.PushBack(&entry{value: 1})
	l.PushBack(&entry{value: 2})

	e := l.Front()
	l.Remove(e)
	if e.Value().value != 1 {
		t.Errorf("e.value = %d, want 1", e.value)
	}
	if e.Next() != nil {
		t.Errorf("e.Next() != nil")
	}
	if e.Prev() != nil {
		t.Errorf("e.Prev() != nil")
	}
}

func TestMove(t *testing.T) {
	l := New[entry]()
	e1 := l.PushBack(&entry{value: 1})
	e2 := l.PushBack(&entry{value: 2})
	e3 := l.PushBack(&entry{value: 3})
	e4 := l.PushBack(&entry{value: 4})

	l.MoveAfter(e3, e3)
	checkListPointers(t, l, []*Element[entry]{e1, e2, e3, e4})
	l.MoveBefore(e2, e2)
	checkListPointers(t, l, []*Element[entry]{e1, e2, e3, e4})

	l.MoveAfter(e3, e2)
	checkListPointers(t, l, []*Element[entry]{e1, e2, e3, e4})
	l.MoveBefore(e2, e3)
	checkListPointers(t, l, []*Element[entry]{e1, e2, e3, e4})

	l.MoveBefore(e2, e4)
	checkListPointers(t, l, []*Element[entry]{e1, e3, e2, e4})
	e2, e3 = e3, e2

	l.MoveBefore(e4, e1)
	checkListPointers(t, l, []*Element[entry]{e4, e1, e2, e3})
	e1, e2, e3, e4 = e4, e1, e2, e3

	l.MoveAfter(e4, e1)
	checkListPointers(t, l, []*Element[entry]{e1, e4, e2, e3})
	e2, e3, e4 = e4, e2, e3

	l.MoveAfter(e2, e3)
	checkListPointers(t, l, []*Element[entry]{e1, e3, e2, e4})
}

// Test PushFront, PushBack, PushFrontList, PushBackList with uninitialized List
func TestZeroList(t *testing.T) {
	var l1 = new(WeakList[entry])
	l1.PushFront(&entry{value: 1})
	checkList(t, l1, []any{1})

	var l2 = new(WeakList[entry])
	l2.PushBack(&entry{value: 1})
	checkList(t, l2, []any{1})

	var l3 = new(WeakList[entry])
	PushFrontList(l1, l3)
	checkList(t, l3, []any{1})

	var l4 = new(WeakList[entry])
	PushBackList(l2, l4)
	checkList(t, l4, []any{1})
}

// Test that a list l is not modified when calling InsertBefore with a mark that is not an element of l.
func TestInsertBeforeUnknownMark(t *testing.T) {
	var l WeakList[entry]
	l.PushBack(&entry{value: 1})
	l.PushBack(&entry{value: 2})
	l.PushBack(&entry{value: 3})
	l.InsertBefore(&entry{value: 1}, new(Element[entry]))
	checkList(t, &l, []any{1, 2, 3})
}

// Test that a list l is not modified when calling InsertAfter with a mark that is not an element of l.
func TestInsertAfterUnknownMark(t *testing.T) {
	var l WeakList[entry]
	l.PushBack(&entry{value: 1})
	l.PushBack(&entry{value: 2})
	l.PushBack(&entry{value: 3})
	l.InsertAfter(&entry{value: 1}, new(Element[entry]))
	checkList(t, &l, []any{1, 2, 3})
}

// Test that a list l is not modified when calling MoveAfter or MoveBefore with a mark that is not an element of l.
func TestMoveUnknownMark(t *testing.T) {
	var l1 WeakList[entry]
	e1 := l1.PushBack(&entry{value: 1})

	var l2 WeakList[entry]
	e2 := l2.PushBack(&entry{value: 2})

	l1.MoveAfter(e1, e2)
	checkList(t, &l1, []any{1})
	checkList(t, &l2, []any{2})

	l1.MoveBefore(e1, e2)
	checkList(t, &l1, []any{1})
	checkList(t, &l2, []any{2})
}
