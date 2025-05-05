package models

import "time"

type JobStatus string

const (
	StatusPending  JobStatus = "pending"
	StatusActive   JobStatus = "active"
	StatusFailed   JobStatus = "failed"
	StatusComplete JobStatus = "complete"
)

type Job struct {
	ID        string    `json:"id"`
	Payload   string    `json:"payload"`
	Status    JobStatus `json:"status"`
	Retries   int       `json:"retries"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
