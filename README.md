# 🅿️ Golang Parking System

This is a take-home assessment project implementing a scalable, concurrent **multi-gate parking lot system** using Go and PostgreSQL.

## ✅ Features

- 🚗 **Park a vehicle**
- 🛻 **Unpark a vehicle**
- 📍 **Search vehicle by plate**
- 📊 **Check available spots**

## ⚙️ Tech Highlights

### 🧱 Architecture

- **Clean Architecture**: Domain → Usecase → Handler separation
- **OpenAPI**: Swagger documentation auto-generated for all endpoints
- **Context-based Transactions**: Passed from handler to domain using `pkg.GetTransactionFromCtx()`
- **Unit Testable**: Each layer is fully testable with mocked interfaces

### 🔐 Concurrency & Data Integrity

- **SQL-backed storage** for production-readiness and scaling
- **Database Transactions** with `BEGIN`, `COMMIT`, `ROLLBACK`
- **Row-Level Locking**: `SELECT ... FOR UPDATE` to prevent race conditions
- **Unique Constraints**: Ensures only one active parking record per vehicle (`spot_id`, `unparked_at IS NULL`)
- **Spot indexing** for fast lookups and integrity

### 📦 Deployment & Environment

- **Dockerized** for local and production environments
- **Traefik Load Balancer** for multi-instance simulation
- **.env config** for environment variables

### 📈 Load Testing

- **k6** used to simulate real-world load across multiple app instances
- Metrics tracked:
  - Latency (P95)
  - Throughput (RPS)
  - Error rates

## 🚀 Getting Started

```bash
cd parking-system

# Copy and configure environment
cp .env.example .env

# Seed DB
make seed

# Start the app
make start
```

## 🔗 Access the App

- **App:** [http://localhost:8080](http://localhost:8080)
- **Swagger:** [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)

---

## 🧪 Run Unit Tests

```bash
make test
make coverage
```

## 🔬 Load Test with k6

This project uses [k6](https://k6.io/docs/getting-started/installation/) for simulating real-world concurrency.

### ▶️ Run Load Test

```bash
make run-load-test
```

## 🧩 Areas for Improvement

- 🔄 **Pagination** for vehicle history
- 🔍 **Observability**: logging, tracing, and metrics
- 📊 **Queue** if there is usecase where No spots available now, but a car can wait for one to be freed. You'll need a waitlist queue.

---

## 🙏 Final Notes

This project was a great opportunity to implement:

- ✅ **Clean architecture**
- ✅ **SQL-based concurrency control**
- ✅ **Real-world scalable backend patterns**

---

**Made with 💻 by Umam**
