# Vibrox Core

`vibrox-core` is the user management microservice of the **Vibrox** suite.  
It exposes REST APIs for user operations and communicates with other services via gRPC.

---

## ✨ Features

- Create, read, update, delete (CRUD) users
- REST API interface using Go + Gin
- gRPC client integration with:
  - `vibrox-auth` for authentication
  - `vibrox-log` for logging events

---

## 🚀 Getting Started

### Prerequisites

- Go 1.21+
- Docker (for running dependencies like Redis or DB)
- (Optional) Postgres / MongoDB, depending on your persistence layer

### Run Locally

```bash
git clone https://github.com/your-username/vibrox-core.git
cd vibrox-core

go mod tidy
go run main.go
```
