package main

import (
	// "context"
	"fmt"
	"log"
	"time"
)

// Queue holds name, list of jobs, and context with cancel
type Queue struct {
	name string
	jobs chan Job
	// ctx    context.Context
	// cancel context.CancelFunc
}

// Job holds logic to perform some operations during queue execution
type Job struct {
	Name   string
	Action func() error // func to be executed when the job is running
}

// Worker responsible for queue serving
type Worker struct {
	Queue *Queue
}

// NewQueue instantiates a new queue
func NewQueue(name string) *Queue {

	return &Queue{
		jobs: make(chan Job, 100),
		name: name,
	}
}

// QUEUE METHODS

// AddJobs adds jobs to the queue
func (q *Queue) AddJobs(jobs []Job) {
	for _, job := range jobs {
		q.AddJob(job)
	}
}

// AddJob sends job to the job to the channel
func (q *Queue) AddJob(job Job) {
	q.jobs <- job
	log.Printf("New job %s added to %s queue", job.Name, q.name)
}

// JOB METHODS

// Run prints the name of the job and calls the action func associated
func (j Job) Run() error {
	log.Printf("Job running: %s", j.GetName())

	err := j.Action()

	if err != nil {
		return err
	}

	return nil
}

// GetName returns the name of the job
func (j Job) GetName() string {
	return j.Name
}

// WORKER FUNCS AND METHODS

// instantiates a new Worker
func NewWorker(queue *Queue) *Worker {
	return &Worker{
		Queue: queue,
	}
}

// DoWork processes jobs from the queue (jobs channel)
func (w *Worker) DoWork() bool {
	for {
		select {

		// if job received
		case job := <-w.Queue.jobs:
			err := job.Run()
			if err != nil {
				log.Print(err)
				continue
			}

		// if no further jobs received after 10ms, return
		case <-time.After(10 * time.Millisecond):
			log.Printf("Work done in queue %s!", w.Queue.name)
			return true
		}
	}

}

var products = []string{
	"books",
	"computers",
}

func main() {
	newProducts := []string{
		"apples",
		"oranges",
		"wine",
		"bread",
		"orange juice",
	}

	// instantiate a new queue
	productsQueue := NewQueue("NewProducts")

	// create list of jobs to be done
	var jobs []Job

	for _, newProduct := range newProducts {
		product := newProduct

		// define the action for each product (the work to be performed on the product)
		// in this case: add each product to the products queue
		action := func() error {
			products = append(products, product)
			return nil
		}

		// for each product add a job to the queue with a unique action
		jobs = append(jobs, Job{
			Name:   fmt.Sprintf("Importing new product: %s", product),
			Action: action,
		})
	}

	// add the jobs to the queue
	productsQueue.AddJobs(jobs)

	// instantiate a worker for the queue
	worker := NewWorker(productsQueue)

	// perform the work on each job in the channel
	worker.DoWork()

	defer fmt.Print(products)
}
