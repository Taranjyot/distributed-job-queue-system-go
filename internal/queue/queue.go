package queue

import (
	"sync"

	"../../pkg/models"
)

type JobQueue struct {
	jobs []models.Job
	mu   sync.Mutex
}

func NewQueue() *JobQueue {
	return &JobQueue{}
}

func (q *JobQueue) Enqueue(job models.Job) {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.jobs = append(q.jobs, job)
}

func (q *JobQueue) Dequeue() *models.Job {
	q.mu.Lock()
	defer q.mu.Unlock()
	if len(q.jobs) == 0 {
		return nil
	}

	job := q.jobs[0]
	q.jobs = q.jobs[1:]
	return &job
}
