# GoTasker 🗂️

A simple CLI-based task manager built in **Go** with **MySQL** backend support.  
It allows users to register, log in, manage tasks, get reminders, and export task data to CSV.

---

## 🚀 Features

- ✅ User registration & login system
- 📝 Create, view, delete, and complete tasks
- ⏰ Background reminders for upcoming tasks using Goroutines
- 🧾 Export your tasks to a CSV file
- 🔐 MySQL-based persistent storage
- 🧵 Uses Go structs, maps, and goroutines (concurrency)
- 📁 Clean modular code with basic CLI interface

---

## 🛠️ Tech Stack

- **Language:** Go (Golang)
- **Database:** MySQL (local)
- **Concurrency:** Go routines
- **CSV Export:** Go's `encoding/csv`

---

## 📦 Setup Instructions

### 1. Clone the repo

```bash
git clone https://github.com/yourusername/GoTasker.git
cd GoTasker
