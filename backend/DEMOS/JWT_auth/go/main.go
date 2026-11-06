package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// ==========================================
// 1. Configuration & Models
// ==========================================

var jwtSecretKey = []byte("super-secret-key-change-me-in-production")

// User represents stored user data.
type User struct {
	Username     string `json:"username"`
	PasswordHash string `json:"-"` // Never expose hash in JSON
}

// Credentials is used for sign up and sign in request payloads.
type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Claims defines the JWT payload structure.
type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// Context key type to safely pass authenticated user info through request context.
type contextKey string

const userContextKey contextKey = "username"

// ==========================================
// 2. In-Memory User Store (Demo Purpose)
// ==========================================

type UserStore struct {
	sync.RWMutex
	users map[string]User
}

var store = &UserStore{
	users: make(map[string]User),
}

func (s *UserStore) Create(username, password string) (*User, error) {
	s.Lock()
	defer s.Unlock()

	if _, exists := s.users[username]; exists {
		return nil, errors.New("user already exists")
	}

	// 1. Hash password with bcrypt before storing
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := User{
		Username:     username,
		PasswordHash: string(hashedPassword),
	}
	s.users[username] = user
	return &user, nil
}

func (s *UserStore) Authenticate(username, password string) (*User, error) {
	s.RLock()
	defer s.RUnlock()

	user, exists := s.users[username]
	if !exists {
		return nil, errors.New("invalid username or password")
	}

	// 2. Compare hashed password with incoming plain password
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return nil, errors.New("invalid username or password")
	}

	return &user, nil
}

// ==========================================
// 3. JWT Helpers (Sign & Verify)
// ==========================================

// generateToken creates a signed JWT string valid for 1 hour.
func generateToken(username string) (string, error) {
	expirationTime := time.Now().Add(1 * time.Hour)

	claims := &Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Subject:   username,
		},
	}

	// Sign the token with HMAC-SHA256 (HS256)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecretKey)
}

// validateToken parses and validates the JWT string.
func validateToken(tokenString string) (*Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		// Ensure the signing method matches expectations
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return jwtSecretKey, nil
	})

	if err != nil || !token.Valid {
		return nil, errors.New("invalid or expired token")
	}

	return claims, nil
}

// ==========================================
// 4. Authentication Middleware
// ==========================================

// authMiddleware inspects the "Authorization: Bearer <token>" header,
// validates the JWT, and attaches the username to the request context.
func authMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			writeJSON(w, http.StatusUnauthorized, map[string]string{"error": "Authorization header missing"})
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			writeJSON(w, http.StatusUnauthorized, map[string]string{"error": "Invalid Authorization header format. Expected 'Bearer <token>'"})
			return
		}

		tokenString := parts[1]
		claims, err := validateToken(tokenString)
		if err != nil {
			writeJSON(w, http.StatusUnauthorized, map[string]string{"error": err.Error()})
			return
		}

		// Inject user identity into request context for downstream handlers
		ctx := context.WithValue(r.Context(), userContextKey, claims.Username)
		next(w, r.WithContext(ctx))
	}
}

// ==========================================
// 5. HTTP Handlers
// ==========================================

func signupHandler(w http.ResponseWriter, r *http.Request) {
	var creds Credentials
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
		return
	}

	if strings.TrimSpace(creds.Username) == "" || strings.TrimSpace(creds.Password) == "" {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "Username and password are required"})
		return
	}

	user, err := store.Create(creds.Username, creds.Password)
	if err != nil {
		writeJSON(w, http.StatusConflict, map[string]string{"error": err.Error()})
		return
	}

	// Generate JWT immediately upon signup
	token, err := generateToken(user.Username)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": "Failed to generate token"})
		return
	}

	writeJSON(w, http.StatusCreated, map[string]interface{}{
		"message": "User registered successfully",
		"token":   token,
		"user": map[string]string{
			"username": user.Username,
		},
	})
}

func signinHandler(w http.ResponseWriter, r *http.Request) {
	var creds Credentials
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
		return
	}

	user, err := store.Authenticate(creds.Username, creds.Password)
	if err != nil {
		writeJSON(w, http.StatusUnauthorized, map[string]string{"error": err.Error()})
		return
	}

	// Generate JWT upon valid sign in
	token, err := generateToken(user.Username)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": "Failed to generate token"})
		return
	}

	writeJSON(w, http.StatusOK, map[string]interface{}{
		"message": "Login successful",
		"token":   token,
		"user": map[string]string{
			"username": user.Username,
		},
	})
}

// profileHandler is a protected route accessible only with a valid JWT
func profileHandler(w http.ResponseWriter, r *http.Request) {
	// Retrieve username injected by authMiddleware
	username, ok := r.Context().Value(userContextKey).(string)
	if !ok {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": "User context not found"})
		return
	}

	writeJSON(w, http.StatusOK, map[string]interface{}{
		"message":  "Welcome to your protected profile!",
		"username": username,
		"timestamp": time.Now(),
	})
}

// Helper to write JSON response
func writeJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(data)
}

// ==========================================
// 6. Server Initialization
// ==========================================

func main() {
	mux := http.NewServeMux()

	// Public routes
	mux.HandleFunc("POST /signup", signupHandler)
	mux.HandleFunc("POST /signin", signinHandler)

	// Protected route wrapped by authMiddleware
	mux.HandleFunc("GET /profile", authMiddleware(profileHandler))

	port := ":8080"
	fmt.Printf("Go JWT Server running on http://localhost%s\n", port)
	if err := http.ListenAndServe(port, mux); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
