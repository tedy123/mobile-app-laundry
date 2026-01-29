# Laundry App Backend 🧺

A robust RESTful API built with Go to power the Laundry App management system. This backend handles user authentication, transaction tracking, and token-based security.

## 🚀 Tech Stack

- **Language:** Go (Golang) 1.23+
- **Database:** PostgreSQL
- **Security:** 
    - JWT (JSON Web Tokens) for session management
    - bcrypt for password hashing
    - **Direct Google Token Verification**: Uses `oauth2.googleapis.com/tokeninfo` for secure, server-side Google ID Token validation (no Firebase Service Account JSON required).
- **Framework:** Net/HTTP (Standard library) with custom CORS middleware.

## 📋 Features

- **User Authentication:** 
    - Standard Email/Password login.
    - Secure Google Sign-In verification via direct Google API.
    - Automatic user registration upon first Google login.
- **Transaction Management:**
    - Record daily laundry intake (`qty_in`) and outtake (`qty_out`).
    - Track revenue and expenses (`amount_in` / `amount_out`).
    - **Smart Querying**: Fuzzy date matching (`DATE(tanggal)`) to handle timezone discrepancies between mobile devices and the database.
- **CORS Support**: Pre-configured to allow requests from local development environments and mobile webviews.

## 🛠️ Project Structure

```text
backend/
├── cmd/api/          # Application entry point (main.go)
├── internal/
│   ├── db/           # Database connections (Postgres)
│   ├── handler/      # HTTP Request handlers & routing logic
│   ├── service/      # Business logic layer (Auth, Transactions)
│   ├── repository/   # Data access layer (Postgres queries)
│   ├── model/        # Data structures (User, Transaction)
│   └── utils/        # Utilities (JWT generator, Token parser)
└── go.mod            # Dependency management
```

## ⚙️ Environment Variables

Create your own environment setup or set these variables:

| Variable | Description | Example |
| :--- | :--- | :--- |
| `DATABASE_URL` | PostgreSQL connection string | `host=localhost user=postgres password=root dbname=laundry port=5432 sslmode=disable` |
| `PORT` | API Port | `4000` |

## 📥 Getting Started

1.  **Clone the repository**.
2.  **Database Setup:**
    - Create a PostgreSQL database named `laundry`.
    - Create the `users` and `transactions` tables.
    - **Transaction Table Columns**: `id`, `email`, `nama`, `alamat`, `tanggal` (DATE/TIMESTAMP), `qty_in`, `qty_out`, `amount_in`, `amount_out`.
3.  **Install Dependencies:**
    ```bash
    go mod tidy
    ```
4.  **Run the Server:**
    ```bash
    go run cmd/api/main.go
    ```
    The API will be available at `http://localhost:4000`.

## 🔍 Debugging & Logs

The backend includes detailed logging to help you track data flow:
- `[DEBUG] Fetching Transactions...`: Shows which email and date the mobile app is requesting.
- `[DB-RESULT] Found for...`: Shows the summary of data retrieved from PostgreSQL.
- `[API] Response 404/200`: Shows the status of incoming requests.

---
🐾 *Created for Ayocuci Laundry System*
