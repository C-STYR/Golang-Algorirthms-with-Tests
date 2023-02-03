package main

import "testing"

func TestQueue(t *testing.T) {
	//
	// nq := newQueue("test_queue")

	// if nq.name != "test_queue" {
	// 	t.Error("queue name not properly assigned")
	// }	
	
	// nq.addItemtoQueue("apples")
	// nq.addItemtoQueue("bananas")
	// nq.addItemtoQueue("carrots")

	// if nq.getQueueLength() != 3 {
	// 	t.Error("items not properly added to queue")
	// }

	// popped := nq.getNextItem()
	// if popped != "apples" {
	// 	t.Errorf("expected %s but got %v", "apples", popped)
	// }

	// nq.getNextItem()
	// nq.getNextItem()

	// popped = nq.getNextItem()
	// if popped != "" {
	// 	t.Errorf("expected empty string, but got %v", popped)
	// }
}

func TestQueueSolution(t *testing.T) {
	nq := newSolutionQueue("test_queue")

	if nq.name != "test_queue" {
		t.Error("queue name not properly assigned")
	}	
	
	nq.addItemtoSolutionQueue("apples")
	nq.addItemtoSolutionQueue("bananas")
	nq.addItemtoSolutionQueue("carrots")

	if nq.getSolutionQueueLength() != 3 {
		t.Error("items not properly added to queue")
	}

	popped := nq.getNextSolutionItem()
	if popped != "apples" {
		t.Errorf("expected %s but got %v", "apples", popped)
	}

	nq.getNextSolutionItem()
	nq.getNextSolutionItem()

	popped = nq.getNextSolutionItem()
	if popped != "" {
		t.Errorf("expected empty string, but got %v", popped)
	}
}