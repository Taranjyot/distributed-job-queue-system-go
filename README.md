# 🚀 Distributed Job Queue System (Go)

A production-grade, distributed job queue built with Go — inspired by AWS SQS. Built for reliability, observability, and extensibility.

---

## ✨ Features

- 🔁 Reliable job queuing with retry logic
- ☠️ Dead Letter Queue for failed jobs
- 🧠 In-memory + persistent storage (Redis/BadgerDB)
- 📊 Prometheus metrics for queue health
- 🔌 REST/gRPC API for producers and workers
- 🐳 Docker-ready deployment

---

## 📐 Architecture

![Architecture Diagram](./docs/architecture.png)

---

## 🗂️ Project Structure

```
.
├── cmd/                # Server and worker binaries
│   ├── server/
│   └── worker/
├── internal/
│   ├── api/            # REST/gRPC handlers
│   ├── queue/          # Job queue logic
│   └── storage/        # Redis/BadgerDB
├── pkg/
│   └── models/         # Job model, constants
├── deployments/
│   └── docker-compose.yml
├── metrics/
│   └── prometheus.yml
├── docs/
│   └── architecture.png
└── README.md
```

---

## 🚀 Quick Start

```bash
# Start server
go run ./cmd/server

# Start a worker
go run ./cmd/worker
```

Or use Docker Compose:

```bash
docker-compose up --build
```

---

## 📡 API Endpoints

| Method | Endpoint     | Description            |
|--------|--------------|------------------------|
| POST   | /jobs        | Enqueue a new job      |
| GET    | /jobs        | Fetch a job to process |
| POST   | /ack         | Acknowledge job result |

---

## 📊 Metrics

Exposed on `/metrics` (Prometheus format):
- `queue_length`
- `job_retries_total`
- `jobs_processed_total`
- `jobs_failed_total`

---

## 🧱 Tech Stack

- Go 1.21+
- Redis or BadgerDB
- gRPC or HTTP
- Docker + Docker Compose
- Prometheus

---

## 📌 Future Plans

- Multi-priority queues
- Rate-limited workers
- Scheduling / delayed jobs
- TLS + mTLS between services
