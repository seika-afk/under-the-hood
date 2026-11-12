![[Pasted image 20261111082515.png]]


#### Entry point

When a request arrives, the operating system forwards the HTTP request to whichever **port** your server is listening on, `:3000`, `:4000`, or any port you configured. The point where your server first receives that request is its **entry point**. Your server is always _listening_ on that port, and that listening is how it picks up incoming requests.

#### Routing

Immediately after the entry point comes **routing**. A server exposes many routes, `/users`, `/users/123`, dynamic routes like `/users/:id`, and so on. The **routing algorithm** inspects the incoming request’s method and path and **maps it to a particular handler**



## Controller / Handler layer

The handler is the entry point for a route once routing has matched it

`**control the flow of data**`

- Data Extraction
    
    Pull the relevant data out of the request object. _What_ you pull depends on the method: **GET** → query parameters; **POST / PUT / PATCH** → the request body; **DELETE** → usually nothing, sometimes a body.
    
- Binding (Deserialization)
    
    The body arrived as a **JSON** string because JSON is a serializable format that travels well across the front-end/back-end boundary. The handler **deserializes** it into the language’s native type, a Go struct, a Python dict/class, a Rust struct. Frameworks call this **binding**. If it fails, halt immediately and return 400 Bad Request.
    
- Validation
    
    Confirm the data matches the expected shape: mandatory fields present, types correct, no malicious payloads. Validate _everything_ from an external client, path params, query params, and body, and be as specific as possible.
    
- Transformation
    
    Optionally reshape the validated data for the convenience of downstream layers, for example, injecting default values.
- Delegation
    
    Pass the clean, validated, transformed data (plus context like the authenticated user ID) **down to the Service layer**.
    
- Sending the Response
    
    When the service returns, choose the correct status code and send the final response back to the client


```go
func ListBooksHandler(w http.ResponseWriter, r *http.Request) {
    sort := r.URL.Query().Get("sort") // "" if absent

    // VALIDATION: if present, it must be one of the allowed values.
    if sort != "" && sort != "name" && sort != "date" {
        http.Error(w, "sort must be 'name' or 'date'", http.StatusBadRequest)
        return
    }

    // TRANSFORMATION: query params are optional -> inject a default.
    if sort == "" {
        sort = "date" // downstream layers never see an empty value
    }

    books, err := bookService.ListBooks(r.Context(), sort) // delegate
    if err != nil {
        http.Error(w, "could not fetch books", http.StatusInternalServerError) // 500
        return
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(books) // 200 with the array of books
}
```
## The Service Layer

The service layer should know **absolutely nothing about HTTP**. No request/response objects, no status codes, no validation concerns

it calls separate repositery layer functions
```go
// Notice: no http.Request, no ResponseWriter, no status codes.
// You cannot tell from this signature that it serves an API.
func (s *BookService) ListBooks(ctx context.Context, sort string) ([]Book, error) {
    // Orchestration: call the repository for the data it needs.
    books, err := s.repo.FindAllBooks(ctx, sort)
    if err != nil {
        return nil, err // bubble the error up; the handler decides the code
    }
    // Could also: enrich, merge other repo calls, send notifications...
    return books, nil
}

// A service that needs no database at all is perfectly valid:
func (s *BookService) NotifyOwner(email string) error {
    return s.mailer.Send(email, "Your book was added")
}
```

## The Repositery Layer

The repository (or _database_) layer has a single concern: **talking to the database**. It receives data from the service, **constructs the database query**: for inserting, filtering, or sorting, runs it, and returns the raw result back up to the service.

```go
// ONE method, ONE shape of result: all books, sorted.
func (r *BookRepo) FindAllBooks(ctx context.Context, sort string) ([]Book, error) {
    // Build the query from the data the service handed down.
    query := fmt.Sprintf("SELECT id, title, author FROM books ORDER BY %s", sort)
    rows, err := r.db.QueryContext(ctx, query)
    if err != nil { return nil, err }
    defer rows.Close()

    var books []Book
    for rows.Next() {
        var b Book
        rows.Scan(&b.ID, &b.Title, &b.Author)
        books = append(books, b)
    }
    return books, nil
}
```

![[Pasted image 20261112120513.png]]

## Middlewares

![[Pasted image 20261112120550.png]]
```go
// In Go, middleware wraps the "next" handler. Calling next.ServeHTTP

// is the equivalent of next(): pass execution along the chain.
func LoggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        log.Printf("%s %s", r.Method, r.URL.Path) // do work

        // EARLY EXIT example (short-circuit, never calls next):
        if r.Header.Get("X-Blocked") == "yes" {
            http.Error(w, "forbidden", http.StatusForbidden)
            return // request stops here
        }

        next.ServeHTTP(w, r) // === next(): continue the chain ===
    })
}
```

![[Pasted image 20261112120859.png]]