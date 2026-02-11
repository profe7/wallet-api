# Wallet API

A RESTful wallet service built in Go using only the standard library (no frameworks). It supports balance inquiries and withdrawals, backed by SQLite.

## Architecture

The project follows a clean layered architecture inspired by Spring Boot conventions:

```
main/main.go            → Entry point & route registration
handler/                 → HTTP handlers (Controller layer)
service/                 → Business logic (Service layer)
repository/              → Database access (Repository layer)
model/                   → Data structures (DTOs & entities)
middleware/              → Cross-cutting concerns (error handling)
middleware/httperrors/    → Typed HTTP error definitions
utils/                   → Shared utilities (JSON writer)
db/                      → Database initialisation & migrations
```

Dependency flow: **Handler → Service → Repository → DB**

## Tech Stack

- **Go** : standard library `net/http` for routing and server
- **SQLite** via `modernc.org/sqlite`
- **Manual dependency injection**

## API Endpoints

### Check Balance

```
GET /balance?user_id=evaristeGalois
```

**Expected Response `200 OK`**
```json
{
  "success": true,
  "data": {
    "user_id": "evaristeGalois",
    "balance": 322322322
  }
}
```

### Withdraw

```
POST /withdraw
Content-Type: application/json

{
  "user_id": "evaristeGalois",
  "amount": 100
}
```

**Expected Response `200 OK`**
```json
{
  "success": true,
  "message": "Withdrawal successful",
  "data": {
    "user_id": "evaristeGalois",
    "withdrawn": 100,
    "new_balance": 322322222
  }
}
```

### Error Responses

All errors follow a consistent template:
```json
{
  "success": false,
  "message": "[Source] description of the problem"
}
```

| Status | Scenario |
|--------|----------|
| `400`  | Missing/invalid parameters, insufficient funds |
| `404`  | Account not found |
| `405`  | Wrong HTTP method |
| `500`  | Unexpected server error |

## Running

```bash
go run main/main.go
```

Server starts on `http://localhost:8080`. A demo account (`evaristeGalois`) is seeded automatically on first run.

## Design Decisions

- **No external router** : uses `net/http.HandleFunc` to keep dependencies minimal.
- **Typed errors** : `APIError` struct carries HTTP status code, source tag, and message; the `ErrorHandler` middleware maps these to JSON responses automatically.
- **Transactional withdrawals** : the repository uses `BEGIN` / `COMMIT` / `ROLLBACK` to ensure atomicity.
