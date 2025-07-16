# vibrox-core

`vibrox-core` is the user management microservice in the **Vibrox** suite.  
It provides REST APIs for user operations and acts as a gRPC client to other services.

---

## ✨ Features

- User CRUD via REST APIs (Go + Gin)
- gRPC client integration with:
  - [`vibrox-auth`](https://github.com/your-username/vibrox-auth) – authentication
  - [`vibrox-echo`](https://github.com/your-username/vibrox-echo) – centralized logging

---

## 🚀 Getting Started

### Prerequisites

- Go 1.24+
- Docker (for local DB, etc.)
- PostgreSQL

### Run Locally

```bash
git clone https://github.com/your-username/vibrox-core.git
cd vibrox-core

go mod tidy
go run main.go
```
