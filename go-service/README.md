# 🧠 Go Service - Student Report Microservice

This microservice, built with **Go (Golang)** using the **Gin** framework, exposes an endpoint to generate student reports in PDF format. It is part of a larger ecosystem including a **PostgreSQL** database and a **Node.js** backend API for authentication and student management.

---

## 📦 Project Structure

```
.
├── backend/               # Node.js API (authentication & management)
├── go-service/            # Go microservice (student report generation)
├── seed_db/               # SQL scripts for tables and mock data
├── docker-compose.yml     # Service orchestration
└── README.md              # This file
```

---

## 🚀 Services Overview

| Service       | Port | Description                                  |
| ------------- | ---- | -------------------------------------------- |
| `go-service`  | 8081 | Go service that returns student PDF reports  |
| `backend-api` | 5007 | Node.js API for login and student management |
| `school-db`   | 5440 | PostgreSQL with student and user data        |

---

## 📄 Main Endpoint (`go-service`)

### `GET /api/v1/students/:id/report`

Generates and returns a PDF report for the student identified by `:id`.

#### ❗ Important:

You **do not need to manually send authentication cookies** in the request.  
The Go service authenticates itself using credentials defined in its environment variables.

#### Example:

```http
GET http://localhost:8081/api/v1/students/1/report
```

Returns: `application/pdf`

---

## ⚙️ How to Run Locally

### 1. Clone the repository

```bash
git clone https://github.com/your-username/your-repo-name.git
cd your-repo-name
```

### 2. Configure environment variables

- For the Node.js service: create `backend/.env` based on `.env.example`.
- For the Go service: create `go-service/.env`, example:

```env
LOG_LEVEL=debug
USERNAME=admin@school-admin.com
PASSWORD=3OU4zn3q6Zh9
SERVER_PORT=8081
API_SERVICE_URL=http://backend-api:5007
```

### 3. Start services with Docker Compose

```bash
docker-compose up --build -d
```

This will launch:

✅ PostgreSQL  
✅ Migration runner with seed data  
✅ Node.js backend  
✅ Go microservice

---

## 🧪 Testing with Postman

To test the report endpoint:

```
GET http://localhost:8081/api/v1/students/1/report
```

Headers and cookies are managed internally by the service.  
You can directly hit the endpoint without extra authentication.

---

## 🧰 Tech Stack

- **Go 1.24.4+**
- **PostgreSQL 16**
- **Node.js 20**
- **Docker & Docker Compose**

---

## 🤝 Contributions

Ideas for improving the PDF formatting, extending endpoints, or adding auth layers?  
Contributions are welcome! 🚀

---

## 📌 TODOs

- [ ] Add unit and integration tests
- [ ] Improve request validation
- [ ] Enhance PDF layout and formatting
