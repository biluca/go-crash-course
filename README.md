# 📦 Go REST API Crash Course Project

This project is a hands-on implementation of a RESTful API in **Go (Golang)**, based on the YouTube crash course by [Traversy Media](https://www.youtube.com/@TraversyMedia):  
👉 [Watch the Course on YouTube](https://www.youtube.com/watch?v=8uiZC0l4Ajw)

It follows the structure and concepts presented in the original video and is based on the excellent starter repository by [@avukadin](https://github.com/avukadin):  
📂 [Original GitHub Repo](https://github.com/avukadin/goapi/tree/main)

---

## ✨ Features

- 🌐 RESTful API built with Go and Gorilla Mux
- 📦 Modular folder structure
- 📄 JSON request/response handling
- 🗃️ Simple in-memory data store (slices)
- 📕 Clean code & separation of concerns

---

## 🛠️ Tech Stack

- **Language:** Go (Golang)
- **Framework:** Gorilla Mux
- **JSON Parsing:** `encoding/json`
- **Routing & HTTP Handling:** `net/http`
- **Data Layer:** In-memory slice (can be replaced with DB)
- **Go Modules** for dependency management

---

## 🚀 Getting Started

### Prerequisites

- Go 1.20+ installed: https://golang.org/dl/

### Clone the repo

```bash
git clone https://github.com/your-username/goapi-crash-course.git
cd goapi-crash-course
```

### Install dependencies

```bash
go mod tidy
```

### Run the application

```bash
go run cmd/main.go
```

The API will be available at:  
`http://localhost:8001/account/coins`

You must pass the UserName `username` as a Query Param, alongise the `Header`:`Authorization`

Sample:
`username:  john_doe`
`Authorization: 123-abc`
