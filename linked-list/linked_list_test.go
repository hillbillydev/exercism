// API:
//
// type Node struct
// type List struct
// var ErrEmptyList
//
// func (e *Node) Next() *Node
// func (e *Node) Prev() *Node
// func NewList(args ...interface{}) *List
// func (l *List) PushFront(v interface{})
// func (l *List) PushBack(v interface{})
// func (l *List) PopFront() (interface{}, error)
// func (l *List) PopBack() (interface{}, error)
// func (l *List) Reverse() *List
// func (l *List) First() *Node
// func (l *List) Last() *Node
package linkedlist

import (
	"bytes"
	"fmt"
	"testing"
)

var newListTestCases = []struct {
	name     string
	in       []interface{}
	expected []interface{}
}{
	{
		name:     "from 5 elements",
		in:       []interface{}{1, 2, 3, 4, 5},
		expected: []interface{}{1, 2, 3, 4, 5},
	},
	{
		name:     "from 2 elements",
		in:       []interface{}{1, 2},
		expected: []interface{}{1, 2},
	},
	{
		name:     "from no element",
		in:       []interface{}{},
		expected: []interface{}{},
	},
	{
		name:     "from 1 element",
		in:       []interface{}{999},
		expected: []interface{}{999},
	},
}

var reverseTestCases = []struct {
	name     string
	in       []interface{}
	expected []interface{}
}{
	{
		name:     "from 5 elements",
		in:       []interface{}{1, 2, 3, 4, 5},
		expected: []interface{}{5, 4, 3, 2, 1},
	},
	{
		name:     "from 2 elements",
		in:       []interface{}{1, 2},
		expected: []interface{}{2, 1},
	},
	{
		name:     "from no element",
		in:       []interface{}{},
		expected: []interface{}{},
	},
	{
		name:     "from 1 element",
		in:       []interface{}{999},
		expected: []interface{}{999},
	},
}

var pushPopTestCases = []struct {
	name     string
	in       []interface{}
	actions  []checkedAction
	expected []interface{}
}{
	{
		name: "PushFront only",
		in:   []interface{}{},
		actions: []checkedAction{
			pushFront(4),
			pushFront(3),
			pushFront(2),
			pushFront(1),
		},
		expected: []interface{}{1, 2, 3, 4},
	},
	{
		name: "PushBack only",
		in:   []interface{}{},
		actions: []checkedAction{
			pushBack(1),
			pushBack(2),
			pushBack(3),
			pushBack(4),
		},
		expected: []interface{}{1, 2, 3, 4},
	},
	{
		name: "PopFront only, pop some elements",
		in:   []interface{}{1, 2, 3, 4},
		actions: []checkedAction{
			popFront(1, nil),
			popFront(2, nil),
		},
		expected: []interface{}{3, 4},
	},
	{
		name: "PopFront only, pop till empty",
		in:   []interface{}{1, 2, 3, 4},
		actions: []checkedAction{
			popFront(1, nil),
			popFront(2, nil),
			popFront(3, nil),
			popFront(4, nil),
			popFront(nil, ErrEmptyList),
		},
		expected: []interface{}{},
	},
	{
		name: "PopBack only, pop some elements",
		in:   []interface{}{1, 2, 3, 4},
		actions: []checkedAction{
			popBack(4, nil),
			popBack(3, nil),
		},
		expected: []interface{}{1, 2},
	},
	{
		name: "PopBack only, pop till empty",
		in:   []interface{}{1, 2, 3, 4},
		actions: []checkedAction{
			popBack(4, nil),
			popBack(3, nil),
			popBack(2, nil),
			popBack(1, nil),
			popBack(nil, ErrEmptyList),
		},
		expected: []interface{}{},
	},
	{
		name: "mixed actions",
		in:   []interface{}{2, 3},
		actions: []checkedAction{
			pushFront(1),
			pushBack(4),
			popFront(1, nil),
			popFront(2, nil),
			popBack(4, nil),
			popBack(3, nil),
			popBack(nil, ErrEmptyList),
			popFront(nil, ErrEmptyList),
			pushFront(8),
			pushBack(7),
			pushFront(9),
			pushBack(6),
		},
		expected: []interface{}{9, 8, 7, 6},
	},
}

// checkedAction calls a function of the linked list and (possibly) checks the result
type checkedAction func(*testing.T, *List)

func pushFront(arg interface{}) checkedAction {
	return func(t *testing.T, ll *List) {
		ll.PushFront(arg)
	}
}

func pushBack(arg interface{}) checkedAction {
	return func(t *testing.T, ll *List) {
		ll.PushBack(arg)
	}
}

func popFront(expected interface{}, expectedErr error) checkedAction {
	return func(t *testing.T, ll *List) {
		v, err := ll.PopFront()
		if err != expectedErr {
			t.Errorf("PopFront() returned wrong, expected no error, got= %v", err)
		}

		if expectedErr == nil && v != expected {
			t.Errorf("PopFront() returned wrong, expected= %v, got= %v", expected, v)
		}
	}
}

func popBack(expected interface{}, expectedErr error) checkedAction {
	return func(t *testing.T, ll *List) {
		v, err := ll.PopBack()
		if err != expectedErr {
			t.Errorf("PopBack() returned wrong, expected no error, got= %v", err)
		}

		if expectedErr == nil && v != expected {
			t.Errorf("PopBack() returned wrong, expected= %v, got= %v", expected, v)
		}
	}
}

func TestNew(t *testing.T) {
	for _, tc := range newListTestCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := NewList(tc.in...)
			fmt.Println(actual)

			checkDoublyLinkedList(t, actual, tc.expected)
		})
	}
}

func TestReverse(t *testing.T) {
	for _, tc := range reverseTestCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := NewList(tc.in...)
			actual.Reverse()

			checkDoublyLinkedList(t, actual, tc.expected)
		})
	}
}

func TestPushPop(t *testing.T) {
	for _, tc := range pushPopTestCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := NewList(tc.in...)
			for _, ac := range tc.actions {
				ac(t, actual)
			}

			checkDoublyLinkedList(t, actual, tc.expected)
		})
	}
}

// checkDoublyLinkedList checks that the linked list is constructed correctly.
func checkDoublyLinkedList(t *testing.T, ll *List, expected []interface{}) {
	// check that length and elements are correct (scan once from begin -> end)
	elem, count, idx := ll.First(), 0, 0
	for ; elem != nil && idx < len(expected); elem, count, idx = elem.Next(), count+1, idx+1 {
		if elem.Val != expected[idx] {
			t.Errorf("wrong value from %d-th element, expected= %v, got= %v", idx, expected[idx], elem.Val)
		}
	}
	if !(elem == nil && idx == len(expected)) {
		t.Errorf("expected %d elements, got= %d", len(expected), count)
	}

	// if elements are the same, we also need to examine the links (next & prev)
	switch {
	case ll.First() == nil && ll.Last() == nil: // empty list
		return
	case ll.First() != nil && ll.Last() != nil && ll.First().Next() == nil: // 1 element
		valid := ll.First() == ll.Last() &&
			ll.First().Next() == nil &&
			ll.First().Prev() == nil &&
			ll.Last().Next() == nil &&
			ll.Last().Prev() == nil

		if !valid {
			t.Errorf("expected to only have 1 element and no links, got= %v", ll.debugString())
		}
	}

	// >1 element
	if ll.First().Prev() != nil {
		t.Errorf("expected Head.Prev() == nil, got= %v", ll.First().Prev())
	}

	prev := ll.First()
	cur := ll.First().Next()
	for idx := 0; cur != nil; idx++ {
		if !(prev.Next() == cur && cur.Prev() == prev) {
			t.Errorf("%d-th element's links is wrong", idx)
		}

		prev = cur
		cur = cur.Next()
	}

	if ll.Last().Next() != nil {
		t.Errorf("expected Last().Next() == nil, got= %v", ll.Last().Next())
	}
}

// debugString prints the linked list with both node's Val, next & prev pointers.
func (ll *List) debugString() string {
	buf := bytes.NewBuffer([]byte{'{'})
	buf.WriteString(fmt.Sprintf("First()= %p; ", ll.First()))

	for cur := ll.First(); cur != nil; cur = cur.Next() {
		buf.WriteString(fmt.Sprintf("[Prev()= %p, Val= %p (%v), Next()= %p] <-> ", cur.Prev(), cur, cur.Val, cur.Next()))
	}

	buf.WriteString(fmt.Sprintf("; Last()= %p; ", ll.Last()))
	buf.WriteByte('}')

	return buf.String()
}
