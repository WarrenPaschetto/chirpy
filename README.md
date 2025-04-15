# 🐤 Chirpy - Microblogging API Server

Chirpy is a Twitter-like microblogging API server built in Go. This was a guided project from [Boot.dev](https://www.boot.dev). Chirpy allows users to create, read, update, and delete short messages known as "chirps." The app supports user registration, login with JWT-based authentication, refresh token management, and webhook integration. The project was built as part of a hands-on learning experience to deepen understanding of backend development and API design.


## Features

- User registration & login with JWT authentication
- Chirp creation, retrieval, deletion
- Webhooks for upgrading users
- Refresh token creation, renewal, and revocation
- Custom middleware for metrics

---

## Setup

### Prerequisites
- Go 1.21+
- PostgreSQL
- [Goose](https://github.com/pressly/goose) CLI installed for database migrations
- `.env` file with the following variables:
  ```env
  DB_URL=your_database_url
  PLATFORM=development
  JWT_SECRET=your_jwt_secret
  POLKA_KEY=your_webhook_api_key
  ```

### Database Migration
Make sure your database is initialized and up to date with the latest schema:
```bash
goose -dir sql/schema postgres "$DB_URL" up
```

### Run the Server
```bash
go run main.go
```

---

## API Documentation

### Health Check
**GET** `/api/healthz`
- Returns 200 OK if server is healthy

---

### Chirps
#### Create a Chirp
**POST** `/api/chirps`
- Requires Bearer JWT token
- Request Body:
```json
{
  "body": "Hello World!"
}
```
- Response:
```json
{
  "id": "UUID",
  "created_at": "timestamp",
  "updated_at": "timestamp",
  "body": "Hello World!",
  "user_id": "UUID"
}
```

#### Get All Chirps
**GET** `/api/chirps`
- Optional query params: `author_id`, `sort=asc|desc`

#### Get Chirp by ID
**GET** `/api/chirps/{chirpID}`

#### Delete Chirp by ID
**DELETE** `/api/chirps/{chirpID}`
- Requires Bearer JWT token (must be owner)

---

### Users
#### Register User
**POST** `/api/users`
- Request Body:
```json
{
  "email": "test@example.com",
  "password": "password123"
}
```

#### Update User
**PUT** `/api/users`
- Requires Bearer JWT token
- Request Body:
```json
{
  "email": "updated@example.com",
  "password": "newpassword"
}
```

---

### Authentication
#### Login
**POST** `/api/login`
- Request Body:
```json
{
  "email": "test@example.com",
  "password": "password123"
}
```
- Response:
```json
{
  "token": "jwt_token",
  "refresh_token": "refresh_token"
}
```

#### Refresh Token
**POST** `/api/refresh`
- Requires Bearer refresh token

#### Revoke Token
**POST** `/api/revoke`
- Requires Bearer refresh token

---

### Webhook
#### Polka Webhook
**POST** `/api/polka/webhooks`
- Requires API key in header: `Authorization: ApiKey {key}`
- Request Body:
```json
{
  "event": "user.upgraded",
  "data": {
    "user_id": "UUID"
  }
}
```

---

## License
This project was built as a guided learning exercise and is free to use and modify.

---

## Author
Created by [Boot.dev](https://www.boot.dev) as part of their Backend Development Track Course.

