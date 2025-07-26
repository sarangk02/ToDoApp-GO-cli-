# 📝 ToDoApp (Go CLI)

A simple Command-Line To-Do application written in Go. Supports adding, listing, and filtering tasks by status, and stores data in a local JSON file.

---

## 📁 Project Structure

```
ToDoApp/
│
├── cmd/
│   └── main.go              # Entry point
│
├── internal/
│   ├── task/               # Task-related business logic
│   │   └── task.go
│   ├── models/              # Shared structs (e.g. Task)
│   │   └── task.go
│   └── storage/             # File read/write logic
│       └── json.go
│
├── pkg/
│   └── utils/               # Reusable utilities (e.g. ID generation)
│       └── id.go
│
├── tasks.json              # Auto-created, stores task data
└── go.mod
```

---

## ⚙️ Features

- ✅ Add tasks with title
- 📋 List all tasks
- 🎯 Filter tasks by status: \`pending\`, \`done\`
- 💾 Persistent task storage in \`tasks.json\`

---

## 🚀 Getting Started

### 1. Clone the repo

### 2. Build & Run

```bash
go mod tidy
go run cmd/main.go
```

---

## 🔧 Tech Stack

- [Go 1.21+](https://golang.org/)
- JSON File Storage (no DB needed)

---

## 📌 Notes

- Task IDs are randomly generated (0–9999).
- All tasks are stored in \`tasks.json\` in the root directory.

---

## 🧠 TODO (optional enhancements)

- [ ] Add task completion functionality
- [ ] Delete tasks
- [ ] Edit tasks
- [ ] Sort by created time or priority

---

## 📄 License

This project is licensed under the MIT License.
`