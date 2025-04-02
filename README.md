
# Backend

## 📌 Tech Stack
- 🏗️ **Go (without a framework)** - Lightweight backend
- 🔀 **`github.com/julienschmidt/httprouter`** - Fast routing library
- 🗄️ **Raw SQL Queries** - Direct database interaction (PostgreSQL/MySQL)

## 🎯 Features
- 🔐 **Authentication**: Implements JWT-based authentication.
- 📜 **History Storage**: Saves conversion history, including text, voice selection, and settings.
- 📡 **RESTful API**: Endpoints for user authentication and history management.

## ⚙️ Installation & Setup
### Prerequisites
- Go
- PostgreSQL/MySQL database

### Setup
1. Set up a PostgreSQL/MySQL database and update the connection details.
2. Run the backend server:
```sh
cd backend
go run main.go
```

## 📡 API Endpoints
### 🔐 Authentication
- `POST /auth/login` - User login
- `POST /auth/register` - User logout

### 🔄 Conversion History
- `GET /histories` - Fetch user's history
- `POST /histories` - Save a new conversion

## ⚖️ Design Decisions & Trade-offs
- **Backend:** Chose Go without a framework for performance and simplicity.
- **Database:** Raw SQL queries for better control and optimization.

## 📝 Author
**Fahem Khakiki Khaya**

