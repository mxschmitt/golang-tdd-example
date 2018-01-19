package stack

import (
	"reflect"
	"testing"
)

// TestIsEmpty: Test 1
// Create a Stack and verify that IsEmpty is true.
func TestIsEmpty(t *testing.T) {
	stack := NewStack()
	isEmpty := stack.IsEmpty()
	if !isEmpty {
		t.Fatalf("expected: %t; got: %t", true, isEmpty)
	}
}

// TestPushOne: Test 2
// Push a single object on the Stack and verify that IsEmpty is false.
func TestPushOne(t *testing.T) {
	stack := NewStack()
	stack.Push("foobar")
	isEmpty := stack.IsEmpty()
	if isEmpty {
		t.Fatalf("expected: %t; got: %t", false, isEmpty)
	}
}

// TestPush: Test 3
// Push a single object, Pop the object, and verify that IsEmpty is true.
func TestPush(t *testing.T) {
	stack := NewStack()
	stack.Push("foobar")
	stack.Pop()
	isEmpty := stack.IsEmpty()
	if !isEmpty {
		t.Fatalf("expected: %t; got: %t", true, isEmpty)
	}
}

// TestPushPopContentCheck: Test 4
// Push a single object, remembering what it is; Pop the object
// and verify that the two objects are equal.
func TestPushPopContentCheck(t *testing.T) {
	stack := NewStack()
	ourItem := "foobar"
	stack.Push(ourItem)
	item := stack.Pop()
	if ourItem != item {
		t.Fatalf("expected: %s; got: %s", ourItem, item)
	}
}

// TestPushPopMultipleElements: Test 5
// Push three objects, remembering what they are; Pop each
// one, and verify that they are removed in the correct order
func TestPushPopMultipleElements(t *testing.T) {
	tc := []string{"foo", "bar", "foobar"}
	stack := NewStack()
	for _, item := range tc {
		stack.Push(item)
	}
	i := len(tc) - 1
	for {
		item := stack.Pop()
		if item == nil {
			break
		}
		if item != tc[i] {
			t.Fatalf("expected: %s; got: %s", tc[i], item)
		}
		i--
	}
}

// TestPopEmptyStack: Test 6
// Pop a Stack that has no elements.
func TestPopEmptyStack(t *testing.T) {
	stack := NewStack()
	item := stack.Pop()
	if item != nil {
		t.Fatalf("expected: %v; got: %v", nil, item)
	}
}

// TestPushTop: Test 7
// Push a single object and then call Top. Verify that IsEmpty
// returns false.
func TestPushTop(t *testing.T) {
	stack := NewStack()
	stack.Push("foobar")
	isEmpty := stack.IsEmpty()
	if isEmpty {
		t.Fatalf("expected: %t; got: %t", false, isEmpty)
	}
}

// TestPushTopContentCheckOneElement: Test 8
// Push a single object, remembering what it is; and then call
// Top. Verify that the object that is returned is equal to the
// one that was pushed.
func TestPushTopContentCheckOneElement(t *testing.T) {
	ourItem := "foobar"
	stack := NewStack()
	stack.Push(ourItem)
	item := stack.Pop()
	if item != ourItem {
		t.Fatalf("expected: %s; got: %s", ourItem, item)
	}
}

// TestPushContentCheckMultiples: Test 9
// Push multiple objects, remembering what they are;  call Top, and
// verify that the last item pushed is equal to the one returned by Top.
func TestPushContentCheckMultiples(t *testing.T) {
	ourItem := "foobar"
	stack := NewStack()
	stack.Push("foo")
	stack.Push("bar")
	stack.Push(ourItem)
	item := stack.Pop()
	if item != ourItem {
		t.Fatalf("expected: %s; got: %s", ourItem, item)
	}
}

// TestPushTopNoStackStateChange: Test 10
// Push one object and call Top repeatedly, comparing what is
// returned to what was pushed.
func TestPushTopNoStackStateChange(t *testing.T) {
	ourItem := map[string]string{
		"foo": "bar",
	}
	stack := NewStack()
	stack.Push(ourItem)
	for i := 0; i < 10; i++ {
		item := stack.Top()
		if !reflect.DeepEqual(item, ourItem) {
			t.Fatalf("expected: %v; got: %v", ourItem, item)
		}
	}
}

// TestTopEmptyStack: Test 11
// Call Top on a Stack that has no elements.
func TestTopEmptyStack(t *testing.T) {
	stack := NewStack()
	item := stack.Top()
	if item != nil {
		t.Fatalf("expected: %v; got: %v", nil, item)
	}
}

// TestPushNull: Test 12
// Push nil onto the Stack and verify that IsEmpty is false.
func TestPushNull(t *testing.T) {
	stack := NewStack()
	stack.Push(nil)
	isEmpty := stack.IsEmpty()
	if isEmpty {
		t.Fatalf("expected: %t; got: %t", false, isEmpty)
	}
}

// TestPushNullCheckPop: Test 13
// Push nil onto the Stack, Pop the Stack, and verify
// that the value returned is nil.
func TestPushNullCheckPop(t *testing.T) {
	stack := NewStack()
	stack.Push(nil)
	item := stack.Pop()
	if item != nil {
		t.Fatalf("expected: %v; got: %v", nil, item)
	}
}

// TestPushNullCheckTop: Test 14
// Push nil onto the Stack, call Top, and verify that the
// value returned is nil
func TestPushNullCheckTop(t *testing.T) {
	stack := NewStack()
	stack.Push(nil)
	item := stack.Top()
	if item != nil {
		t.Fatalf("expected: %v; got: %v", nil, item)
	}
}
