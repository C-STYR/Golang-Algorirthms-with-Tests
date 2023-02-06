package main

import (
	"testing"
)

func Test_Stack(t *testing.T) {
	// // construct a new stack
	// s := NewStack(3)
	// cap := s.capacity
	// if cap != 3 {
	// 	t.Errorf("Expected capacity of %d, but got %d", 3, cap)

	// }

	// // test Push() and Size() and IsFull()
	// s.Push("one")
	// s.Push("two")
	// s.Push(3)

	// full := s.IsFull()
	// if full != true {
	// 	t.Errorf("Expected IsFull() to return true when stack full, but got %v", full)
	// }

	// size := s.Size()
	// if size != 3 {
	// 	t.Errorf("Expected size of %d, but got %d\n", 3, size)
	// }

	// err := s.Push(4)
	// if err == nil {
	// 	t.Error("Expected Push() to return an error when Stack full, but got nil")
	// }

	// // test Peek()
	// peek1, _ := s.Peek()
	// peek2, _ := s.Peek()

	// if peek1 != peek2 {
	// 	t.Error("Expected two consecutive Peek() calls to have the same value")
	// }

	// // test Pop() and IsEmpty()
	// s.Pop()
	// pop2, _ := s.Pop()

	// if pop2 != "two" {
	// 	t.Errorf("Expected popped value to equal %s, but got %v", "two", pop2)
	// }

	// s.Pop()
	// pop4, err := s.Pop()
	// if pop4 != nil {
	// 	t.Errorf("Expected popped value to be nil, but got %v", pop4)
	// }
	// if err == nil {
	// 	t.Error("Expected IsEmtpy() to return error for Pop() on an empty Stack.")
	// }
}

func Test_StackSolution(t *testing.T) {
	// construct a new stack
	s := NewStackSolution(3)
	cap := s.capacity
	if cap != 3 {
		t.Errorf("Expected capacity of %d, but got %d", 3, cap)
	}

	// test Push() and Size() and IsFull()
	s.PushSolution("one")
	s.PushSolution("two")
	s.PushSolution(3)

	full := s.IsFullSolution()
	if full != true {
		t.Errorf("Expected IsFull() to return true when stack full, but got %v", full)
	}

	size := s.SizeSolution()
	if size != 3 {
		t.Errorf("Expected size of %d, but got %d\n", 3, size)
	}

	err := s.PushSolution(4)
	if err == nil {
		t.Error("Expected Push() to return an error when Stack full, but got nil")
	}

	// test Peek()
	peek1, _ := s.PeekSolution()
	peek2, _ := s.PeekSolution()

	if peek1 != peek2 {
		t.Error("Expected two consecutive Peek() calls to have the same value")
	}

	// test Pop() and IsEmpty()
	s.PopSolution()
	pop2, _ := s.PopSolution()

	if pop2 != "two" {
		t.Errorf("Expected popped value to equal %s, but got %v", "two", pop2)
	}

	s.PopSolution()
	pop4, err := s.PopSolution()
	if pop4 != nil {
		t.Errorf("Expected popped value to be nil, but got %v", pop4)
	}
	if err == nil {
		t.Error("Expected IsEmtpy() to return error for Pop() on an empty Stack.")
	}
}
