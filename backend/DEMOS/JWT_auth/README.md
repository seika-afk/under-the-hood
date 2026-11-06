# JWT Authentication Demo (Go & TypeScript)

This project contains two minimal, clear implementations of **JWT (JSON Web Token)** authentication:
- **`go/`**: Written in Go using standard library `net/http` and `github.com/golang-jwt/jwt/v5`.
- **`ts/`**: Written in TypeScript using `express`, `jsonwebtoken`, and `tsx`.

Both implementations follow the exact same architecture, endpoints, and flow to make it easy to compare and understand.

---

## What is JWT & How Does it Work?

A **JSON Web Token (JWT)** is a compact, URL-safe means of representing claims to be transferred between two parties. It allows **stateless authentication**, meaning the server doesn't need to query a session database on every single request; it simply validates the cryptographic signature of the token.

A JWT consists of three parts separated by dots (`.`):
```
header.payload.signature
```
1. **Header**: Contains the token type (`JWT`) and signing algorithm (`HS256`).
2. **Payload (Claims)**: Contains user data (e.g., `username`), issue time (`iat`), and expiration time (`exp`).
3. **Signature**: Created by taking the encoded header, encoded payload, and signing them with a secret key:
   ```
   HMACSHA256(base64UrlEncode(header) + "." + base64UrlEncode(payload), secret)
   ```

### Authentication Flow

```
1. Sign Up / Sign In
   Client  ----------------> POST /signup or POST /signin (username, password)
   Server  ----------------> Hashes/checks password, generates signed JWT
   Client  <---------------- Returns { "token": "eyJhbGci..." }

2. Access Protected Route
   Client  ----------------> GET /profile [Header: Authorization: Bearer eyJhbGci...]
   Server  ----------------> Middleware verifies signature & expiration
   Server  ----------------> Injects user info & executes handler
   Client  <---------------- Returns { "user": { "username": "..." }, "message": "..." }
```

---

## Project Structure

```
JWT_auth/
├── go/
│   ├── go.mod
│   ├── go.sum
│   └── main.go       # Complete Go server: store, JWT sign/verify, middleware, routes
├── ts/
│   ├── package.json
│   ├── tsconfig.json
│   └── src/
│       └── index.ts  # Complete TS server: store, JWT sign/verify, middleware, routes
└── README.md
```

---

## 1. Running the Go Version

### Prerequisites
- Go 1.22+ installed

### Start the Server
```bash
cd go
go run main.go
```
The Go server will start on `http://localhost:8080`.

---

## 2. Running the TypeScript Version

### Prerequisites
- Node.js (v18+) and npm installed

### Start the Server
```bash
cd ts
npm install
npm run dev
```
The TypeScript server will start on `http://localhost:3000`.

---

## 3. Testing the Endpoints (cURL)

Both servers implement identical endpoints. Change the port to `8080` for Go, or `3000` for TypeScript.

### Step 1: Sign Up (Register a new user)
```bash
curl -X POST http://localhost:8080/signup \
  -H "Content-Type: application/json" \
  -d '{"username": "alice", "password": "password123"}'
```
**Response:**
```json
{
  "message": "User registered successfully",
  "token": "<YOUR_JWT_TOKEN>",
  "user": {
    "username": "alice"
  }
}
```

### Step 2: Sign In (Login with existing user)
```bash
curl -X POST http://localhost:8080/signin \
  -H "Content-Type: application/json" \
  -d '{"username": "alice", "password": "password123"}'
```
**Response:**
```json
{
  "message": "Login successful",
  "token": "<YOUR_JWT_TOKEN>",
  "user": {
    "username": "alice"
  }
}
```

### Step 3: Access Protected Route (`/profile`)
Replace `<YOUR_JWT_TOKEN>` with the token from Step 1 or 2:

```bash
curl -X GET http://localhost:8080/profile \
  -H "Authorization: Bearer <YOUR_JWT_TOKEN>"
```
**Response:**
```json
{
  "message": "Welcome to your protected profile!",
  "username": "alice",
  "timestamp": "..."
}
```

### Step 4: Verify Protection (Unauthorized Request)
Try accessing without the `Authorization` header:
```bash
curl -X GET http://localhost:8080/profile
```
**Response:**
```json
{
  "error": "Authorization header missing"
}
```
Or with an invalid token:
```bash
curl -X GET http://localhost:8080/profile \
  -H "Authorization: Bearer invalid-token-string"
```
**Response:**
```json
{
  "error": "invalid or expired token"
}
```
