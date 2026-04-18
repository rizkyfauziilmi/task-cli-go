```md
# Task CLI (Go)

Simple command-line application to manage tasks using Go.  
Project ini dibuat sebagai latihan dari roadmap.sh.

## 📌 Project Reference
Latihan ini berdasarkan project berikut:  
https://roadmap.sh/projects/task-tracker

---

## 🚀 Features

- Add task
- Update task description
- Delete task
- Mark task as in-progress
- Mark task as done
- List tasks (filter by status)

---

## ⚙️ Installation

Clone repository:

```bash
git clone https://github.com/USERNAME/task-cli.git
cd task-cli
````

Init dependency:

```bash
go mod tidy
```

---

## ▶️ Usage

### Add task

```bash
go run main.go add "Belajar Golang"
```

### Update task

```bash
go run main.go update 1 "Belajar Go Coroutines"
```

### Delete task

```bash
go run main.go delete 1
```

### Mark in progress

```bash
go run main.go mark-in-progress 1
```

### Mark done

```bash
go run main.go mark-done 1
```

### List tasks

```bash
go run main.go list
```

Filter by status:

```bash
go run main.go list done
```

---

## 💾 Data Storage

Task disimpan dalam file lokal:

```
tasks.json
```

---

## 🧠 Learning Goals

* Basic Go CLI application
* File handling (JSON storage)
* Struct & modular architecture
* Simple CRUD operations

---

## 📚 Reference

* [https://roadmap.sh/projects/task-tracker](https://roadmap.sh/projects/task-tracker)

```
