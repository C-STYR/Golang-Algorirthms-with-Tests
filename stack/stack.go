package main

/* type already declared in solution file, no need to recreate*/

// type Stack struct {
// 	top        int
// 	capacity   int
// 	stackArray []any
// }

/*Implement the following methods*/

// returns a new Stack
func NewStack(cap int) *Stack {
	s := new(Stack)
	// your code here

	return s
}

// adds an item to the stack
func (s *Stack) Push(data any) error {
	// your code here
	return nil
}

// returns the item from the top of the stack (and an error) and removes it
func (s *Stack) Pop() (any, error) {
	// your code here
	return "", nil
}

// returns the item from the top of the stack without removing it from the stack
func (s *Stack) Peek() (any, error) {
	// your code here
	return "", nil
}

// returns the size of the stack
func (s *Stack) Size() int {
	// your code here
	return 1
}

// returns true if the stack is empty
func (s *Stack) IsEmpty() bool {
	return true
}

// returns true if the stack is full
func (s *Stack) IsFull() bool {
	return true
}
