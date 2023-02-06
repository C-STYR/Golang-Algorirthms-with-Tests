package main

import (
	"errors"
)

type Stack struct {
	top        int
	capacity   int
	stackArray []any
}

func NewStackSolution(cap int) *Stack {
	s := new(Stack)
	s.top = -1
	s.capacity = cap
	s.stackArray = make([]any, cap)
	return s
}

func (s *Stack) PushSolution(data any) error {
	if s.IsFullSolution() {
		return errors.New("stack overflow")
	}
	s.top++ 
	s.stackArray[s.top] = data
	return nil
}

func (s *Stack) PopSolution() (any, error) {
	if s.IsEmptySolution() {
		return nil, errors.New("stack underflow")
	}
	temp := s.stackArray[s.top]
	s.top--
	return temp, nil
}

func (s *Stack) PeekSolution() (any, error) {
	if s.IsEmptySolution() {
		return nil, errors.New("stack underflow")
	}
	temp := s.stackArray[s.top]
	return temp, nil
}

func (s *Stack) SizeSolution() int {
	return s.top + 1
}

func (s *Stack) IsEmptySolution() bool {
	return s.top == -1
}

func (s *Stack) IsFullSolution() bool {
	return s.top == s.capacity-1
}
