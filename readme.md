# 📬 Notification Service (Go + Gin + MongoDB + Docker)

A modular, production-ready **Notification Service** built with **Go**, supporting **Email**, **SMS**, and (soon) **Push notifications**.  
It follows a clean layered architecture (`Handler → Service → Repository`) for clarity and scalability.

---

## 🚀 Features

- Send notifications via **Email** and **SMS**
- Store sent notifications in **MongoDB**
- RESTful API using **Gin**
- Configurable via environment variables
- **Dockerized** for easy deployment
- Future support for **Push notifications (FCM)**

---

## 🏗️ Project Structure

```
notification-service-go/
├── cmd/
│   └── main.go                # Entry point
├── internal/
│   ├── config/                # Configuration loading
│   ├── handler/               # HTTP handlers (API endpoints)
│   ├── model/                 # Data models (Notification struct)
│   ├── repository/            # Database logic (MongoDB CRUD)
│   ├── service/               # Business logic (email/sms)
│   └── utils/                 # Utility helpers (SMTP, etc.)
├── Dockerfile
├── docker-compose.yml
├── go.mod
├── go.sum
└── README.md
```

---

## ⚙️ Setup Instructions

### 1️⃣ Clone Repository

```bash
git clone https://github.com/yourusername/notification-service-go.git
cd notification-service-go
```

---

### 2️⃣ Environment Variables

Create a `.env` file in your root directory:

```bash
PORT=8080

# MongoDB
MONGO_URI=mongodb://mongo:27017
MONGO_DB_NAME=notification_db

# Email SMTP Config
SMTP_HOST=smtp.gmail.com
SMTP_PORT=587
EMAIL_FROM=your_email@gmail.com
EMAIL_PASSWORD=your_app_password

# SMS Provider (for future)
TWILIO_ACCOUNT_SID=
TWILIO_AUTH_TOKEN=
TWILIO_PHONE_NUMBER=
```

> 💡 **Tip:** Use Gmail App Passwords for `EMAIL_PASSWORD` (not your normal login password).

---

### 3️⃣ Run Locally (without Docker)

Make sure MongoDB is running:

```bash
mongod --dbpath /path/to/mongodb/data
```

Then start the server:

```bash
go mod tidy
go run cmd/main.go
```

Your API will be available at: [http://localhost:8081](http://localhost:8081) or your assigned port

---

### 4️⃣ Run with Docker 🐳

---

## 🔌 API Endpoints

### ➤ Send Notification

**POST** `/api/notifications`

#### Request Body
```json
{
  "recipient": "user@example.com",
  "type": "email",
  "subject": "Welcome to our platform!",
  "message": "Hi there, thanks for joining us!"
}
```

#### Response
```json
{
  "status": "success",
  "message": "Notification sent and saved to database"
}
```

---

### ➤ Get All Notifications

**GET** `/api/notifications`

#### Response
```json
[
  {
    "id": "652f123abc...",
    "type": "email",
    "recipient": "user@example.com",
    "subject": "Welcome",
    "message": "Hi there!",
    "sentAt": "2025-10-06T10:20:00Z"
  }
]
```

---

## 🧠 Tech Stack

| Layer | Technology |
|-------|-------------|
| Language | Go (Golang) |
| Framework | Gin |
| Database | MongoDB |
| Email | SMTP (Gmail) |
| Container | Docker |
| SMS (Future) | Twilio |
| Push (Future) | Firebase Cloud Messaging (FCM) |

---

## 🧱 Folder Responsibilities

| Folder | Purpose |
|---------|----------|
| `internal/config` | Loads `.env` configuration |
| `internal/model` | Defines Notification struct |
| `internal/repository` | Handles MongoDB operations |
| `internal/service` | Contains business logic (email/sms sending) |
| `internal/handler` | API routes and controllers |
| `internal/utils` | Utility functions (SMTP, Twilio, etc.) |

---

## 🧪 Testing

To test email sending:

```bash
curl -X POST http://localhost:8080/api/notifications -H "Content-Type: application/json" -d '{
  "recipient": "recipient@example.com",
  "type": "email",
  "subject": "Test Email",
  "message": "This is a test email from Notification Service."
}'
```

---

## 🧱 Future Enhancements

- [ ] Add **Push Notifications (FCM)**
- [ ] Integrate **Kafka / RabbitMQ** for async notification queue
- [ ] Add **Retry mechanism** for failed sends
- [ ] Add **Admin dashboard** to view notification logs

---

## 🧑‍💻 Author

**Piyush Garg**  
---
