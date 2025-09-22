package list

import (
	"unsafe"
)

// WeakList is a doubly linked list based on the standard library container/list. The list uses
// unsafe, weak references so that internal references are not visible to the garbage collector.
// The caller is responsible for holding references to list elements to prevent them from being
// garbage collected.
type (
	WeakList[T any] struct {
		root Element[T]
		len  int
	}

	Element[T any] struct {
		_next, _prev uintptr // weak references to next and previous elements
		_list        uintptr // weak reference to the parent list
		value        *T
	}
)

func (e *Element[T]) next() *Element[T] {
	return (*Element[T])(unsafe.Pointer(e._next))
}

func (e *Element[T]) prev() *Element[T] {
	return (*Element[T])(unsafe.Pointer(e._prev))
}

func (e *Element[T]) list() *WeakList[T] {
	return (*WeakList[T])(unsafe.Pointer(e._list))
}

func (e *Element[T]) Value() *T {
	return e.value
}

// Next returns the next list element or nil.
func (e *Element[T]) Next() *Element[T] {
	if p := e.next(); e.list() != nil && p != &e.list().root {
		return p
	}
	return nil
}

// Prev returns the previous list element or nil.
func (e *Element[T]) Prev() *Element[T] {
	if p := e.prev(); e.list() != nil && p != &e.list().root {
		return p
	}
	return nil
}

// Init initializes or clears list l.
func (l *WeakList[T]) Init() *WeakList[T] {
	l.root._next = uintptr(unsafe.Pointer(&l.root))
	l.root._prev = uintptr(unsafe.Pointer(&l.root))
	l.len = 0
	return l
}

// New returns an initialized list.
func New[T any]() *WeakList[T] {
	return new(WeakList[T]).Init()
}

// Len returns the number of elements of list l.
// The complexity is O(1).
func (l *WeakList[T]) Len() int { return l.len }

// Front returns the first element of list l or nil if the list is empty.
func (l *WeakList[T]) Front() *Element[T] {
	if l.len == 0 {
		return nil
	}

	if l.root._next == 0 {
		return nil
	}

	return l.root.next()
}

// Back returns the last element of list l or nil if the list is empty.
func (l *WeakList[T]) Back() *Element[T] {
	if l.len == 0 {
		return nil
	}

	if l.root._prev == 0 {
		return nil
	}

	return l.root.prev()
}

// lazyInit lazily initializes a zero List value.
func (l *WeakList[T]) lazyInit() {
	if l.root._next == 0 {
		l.Init()
	}
}

// insert inserts e after at, increments l.len, and returns e.
func (l *WeakList[T]) insert(e, at *Element[T]) *Element[T] {
	// fmt.Printf("insert %+v at=%+v\n", e.Value, at)
	e._prev = uintptr(unsafe.Pointer(at))
	e._next = at._next
	e.prev()._next = uintptr(unsafe.Pointer(e))
	e.next()._prev = uintptr(unsafe.Pointer(e))
	e._list = uintptr(unsafe.Pointer(l))
	l.len++
	// fmt.Printf("inserted %+v at=%+v e=%+v\n", e.Value, nil, e)
	return e
}

// insertValue is a convenience wrapper for insert(&Element{Value: v}, at).
func (l *WeakList[T]) insertValue(v *T, at *Element[T]) *Element[T] {
	return l.insert(&Element[T]{value: v}, at)
}

// remove removes e from its list, decrements l.len
func (l *WeakList[T]) remove(e *Element[T]) {
	e.prev()._next = e._next
	e.next()._prev = e._prev
	e._next = 0 // avoid memory leaks
	e._prev = 0 // avoid memory leaks
	e._list = 0
	l.len--
}

// move moves e to next to at.
func (l *WeakList[T]) move(e, at *Element[T]) {
	if e == at {
		return
	}
	e.prev()._next = e._next
	e.next()._prev = e._prev

	e._prev = uintptr(unsafe.Pointer(at))
	e._next = at._next
	e.prev()._next = uintptr(unsafe.Pointer(e))
	e.next()._prev = uintptr(unsafe.Pointer(e))
}

// Remove removes e from l if e is an element of list l.
// It returns the element value e.Value.
// The element must not be nil.
func (l *WeakList[T]) Remove(e *Element[T]) any {
	if e.list() == l {
		// if e.list == l, l must have been initialized when e was inserted
		// in l or l == nil (e is a zero Element) and l.remove will crash
		l.remove(e)
	}
	return e.value
}

// PushFront inserts a new element e with value v at the front of list l and returns e.
func (l *WeakList[T]) PushFront(v *T) *Element[T] {
	l.lazyInit()
	return l.insertValue(v, &l.root)
}

// PushBack inserts a new element e with value v at the back of list l and returns e.
func (l *WeakList[T]) PushBack(v *T) *Element[T] {
	l.lazyInit()
	return l.insertValue(v, l.root.prev())
}

// InsertBefore inserts a new element e with value v immediately before mark and returns e.
// If mark is not an element of l, the list is not modified.
// The mark must not be nil.
func (l *WeakList[T]) InsertBefore(v *T, mark *Element[T]) *Element[T] {
	if mark.list() != l {
		return nil
	}
	// see comment in List.Remove about initialization of l
	return l.insertValue(v, mark.prev())
}

// InsertAfter inserts a new element e with value v immediately after mark and returns e.
// If mark is not an element of l, the list is not modified.
// The mark must not be nil.
func (l *WeakList[T]) InsertAfter(v *T, mark *Element[T]) *Element[T] {
	if mark.list() != l {
		return nil
	}
	// see comment in List.Remove about initialization of l
	return l.insertValue(v, mark)
}

// MoveToFront moves element e to the front of list l.
// If e is not an element of l, the list is not modified.
// The element must not be nil.
func (l *WeakList[T]) MoveToFront(e *Element[T]) {
	if e.list() != l || l.root.next() == e {
		return
	}
	// see comment in List.Remove about initialization of l
	l.move(e, &l.root)
}

// MoveToBack moves element e to the back of list l.
// If e is not an element of l, the list is not modified.
// The element must not be nil.
func (l *WeakList[T]) MoveToBack(e *Element[T]) {
	if e.list() != l || l.root.prev() == e {
		return
	}
	// see comment in List.Remove about initialization of l
	l.move(e, l.root.prev())
}

// MoveBefore moves element e to its new position before mark.
// If e or mark is not an element of l, or e == mark, the list is not modified.
// The element and mark must not be nil.
func (l *WeakList[T]) MoveBefore(e, mark *Element[T]) {
	if e.list() != l || e == mark || mark.list() != l {
		return
	}
	l.move(e, mark.prev())
}

// MoveAfter moves element e to its new position after mark.
// If e or mark is not an element of l, or e == mark, the list is not modified.
// The element and mark must not be nil.
func (l *WeakList[T]) MoveAfter(e, mark *Element[T]) {
	if e.list() != l || e == mark || mark.list() != l {
		return
	}
	l.move(e, mark)
}
