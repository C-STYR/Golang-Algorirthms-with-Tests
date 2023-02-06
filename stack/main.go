package main

import "fmt"

func main() {
	s := NewStackSolution(3)
	cap := s.capacity
	fmt.Println("cap:", cap)

	// test Push() and Size() and IsFull()
	s.PushSolution("one")
	// log.Println("top after one push:", s.top)
	s.PushSolution("two")
	// log.Println("t:op after two push:", s.top)
	s.PushSolution(3)
	fmt.Println("size:", s.SizeSolution())
	fmt.Println("cap:", s.capacity)

	full := s.IsFullSolution()
	fmt.Println("full:", full)

	size := s.SizeSolution()
	fmt.Println("size:", size)
}
