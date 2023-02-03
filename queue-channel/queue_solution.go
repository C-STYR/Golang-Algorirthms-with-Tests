package main

import "fmt"

type Queue struct {
	name  string
	items chan item
}

type item string

func newSolutionQueue(name string) *Queue {
	return &Queue{
		name:  name,
		items: make(chan item, 10),
	}
}

func (q *Queue) addItemtoSolutionQueue(item item) {
	q.items <- item
	fmt.Printf("New item %s added to %s queue", item, q.name)
}

func (q *Queue) getNextSolutionItem() item {
	var popped item
	if len(q.items) > 0 {
		popped = <-q.items
		fmt.Printf("Retrieving next item: %s", popped)
		return popped
	} else {
		fmt.Println("Queue is empty.")
		return ""
	}
}

func (q *Queue) getSolutionQueueLength() int {
	return len(q.items)
}


