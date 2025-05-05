package api

import (
	"encoding/json"
	"net/http"

	"../../pkg/models"
	"../queue"
	"github.com/google/uuid"
)

type Server struct {
	Q *queue.JobQueue
}

func (s *Server) EnqueHandler(w http.ResponseWriter, r *http.Request) {
	var job models.Job
	json.NewDecoder(r.Body).Decode(&job)
	job.ID = uuid.NewString()
	s.Q.Enqueue(job)
	json.NewEncoder(w).Encode(job)
}

func (s *Server) DequeueHandler(w http.ResponseWriter, r *http.Request) {
	job := s.Q.Dequeue()
	if job == nil {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	json.NewEncoder(w).Encode(job)
}
