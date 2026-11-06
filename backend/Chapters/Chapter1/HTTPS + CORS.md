
Application layer” means it’s the topmost layer, the one that defines the _meaning_ of a message (this is a GET for `/users/5`)

![[Pasted image 20260925215658.png]]

1. Statelessness
	1. The server keeps **no session memory between two requests**
	2. ==The protocol has amnesia by design.
	3. Why statelessness?
		1. **Simplicity**: no need to store memory per session
		2. **Horizontal scalability**: because no server “owns” a conversation, any request can be routed to _any_ server in a fleet
		3. **Crash resilience** : The client simply retries, the load balancer sends it elsewhere
	4. The tradeoff is that **continuity must be rebuilt on every request**. Since the protocol itself forgets, we re-establish identity each time using cookies, server-side sessions
2. Client - Server Model
	1. The server is fundamentally passive : hosts resources, listens, and only ever speaks in _response_ to a request

#### WHY TCP?

HTTP doesn’t move bytes itself; it delegates that to a transport protocol.
HTTP cannot tolerate silently dropped or reordered bytes, because a half-delivered header or body is meaningless

```
How its handled in TCP : 

TCP provides reliable, ordered delivery of a byte stream. Lost data is retransmitted when necessary

Each TCP segment has a sequence number-> If bytes arrive later or early -> it knows the sequence

```


### TCP - UDP

**TCP** is connection-based and reliable: it performs a 3-way handshake (SYN -> SYN-ACK -> ACK) to establish a connection, numbers every byte, acknowledges receipt, retransmits losses, and delivers data in order. This is what HTTP/1 and HTTP/2 ride on
![[Pasted image 20260925221216.png]]
The TCP handshake and any TLS negotiation happen below Layer 7, they’re transport/network-engineering concerns


```
SAFE : request is read only 
IDEMPOTENT : same final State
Cacheable : the response may be stored and reused to satisfy a later equivalent request
```
![[Pasted image 20260925222049.png]]

HTTP/3 - Abandons TCP for **QUIC over UDP**.- 2022

```
QUIC Quick UDP Internet Connections
runs over UDP but provides many of the reliability and security features you associate with TCP + TLS.
```
![[Pasted image 20260925222202.png]]

QUIC implements reliable transport, congestion control, stream multiplexing, and encryption using TLS 1.3. It is specified by the IETF.


**HTTP/2’s multiplexing** solves this at the application layer: many independent _streams_ share one connection, and frames from different streams interleave, so a slow response no longer blocks others. HTTP/2 also switched from verbose text to a compact **binary framing** layer and added **HPACK** header compression

![[Pasted image 20261103175411.png]]


```
Across every version the mental model is identical: **a client and server establish a connection, then messages flow back and forth.** The versions only optimize the plumbing, connection reuse, parallelism, compression, and which transport carries it all.
```

![[Pasted image 20261103181545.png]]![[Pasted image 20261103181609.png]]


# Why Headers exists

![[Pasted image 20261103182255.png]]


separating headers from payload :

- extensibility , can add more things on headers itself
- Separates  parcel info from the payload ,like package
**Remote control**: headers let the client send _instructions and preferences_ that steer the server’s behavior, and let the server send directives back. The client says “I’d prefer JSON” (`Accept`), “cache this for an hour” is what the server replies (`Cache-Control`), “here’s who I am” (`Authorization`), “I’m talking to this domain” (`Host`).

# Type of Headers

- request- by user
- general - both user and client
- representation
- security![[Pasted image 20261104095618.png]]

```go
func securityHeaders(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        h := w.Header()
        h.Set("Strict-Transport-Security", "max-age=31536000; includeSubDomains")
        h.Set("Content-Security-Policy", "default-src 'self'")
        h.Set("X-Frame-Options", "DENY")
        h.Set("X-Content-Type-Options", "nosniff")
        next.ServeHTTP(w, r) // headers MUST be set before this call (see sec 3)
    })
}
// http.ListenAndServe(":8080", securityHeaders(mux))
```

**GET** is idempotent, safe  

POST and PATCH are cacheable only when the response explicitly includes freshness info and a matching `Content-Location`, rare in practice, so treat them as effectively non-cacheable


**PUT VS POST** : Post creates new resource , put creates if doesnt exist else updates

**DELETE** is idempotent for a subtle but important reason: deleting once removes the resource, and deleting again changes nothing because it’s already gone



**OPTIONS** asks the server which methods and headers a resource supports without acting on it; the server answers with an `Allow` header (e.g. `Allow: GET, POST, OPTIONS`). Its starring real-world role is the **CORS preflight** (sec 9): before a browser sends a “non-simple” cross-origin request, it automatically fires an OPTIONS to ask permission, and only proceeds if the server’s response approves the method, headers, and origin. You almost never call OPTIONS by hand, but you’ll see it constantly in the browser’s Network tab paired with your real requests.


OPTIONS asks the server which methods and headers are supported. Mainly used by browsers for CORS preflight, checking permission before sending certain cross-origin requests.

CONNECT ,used in forward proxies handling HTTPS

HEAD:  works like GET but returns only the response headers, without the response body.
Used to check resource information (size, content type, last modified, etc.) without downloading the actual content

### Need for idempotency


==Idempotency is the property that makes a request **safe to retry**==
and retries are not optional in a distributed system, they’re inevitable.Networks drop packets, a client times out without knowing whether the request reached the server, a user impatiently double-clicks, a mobile app re-sends when the radio reconnects.

GET,PUT,DELETE are idempotent by design,POST is not, which is the root cause of duplicate-record and double-charge bugs.


The client generates a unique key (typically a UUID) for the _logical operation_ and sends it as an `Idempotency-Key` header. The server stores the key alongside the result of the first successful execution. If the same key ever arrives again, because the client retried, the server recognizes it and returns the **stored result** without re-running the work. The operation thus executes _at most once_ regardless of how many times it’s submitted. (MDN now documents `Idempotency-Key` as a standardizing header, and payment APIs like Stripe popularized the pattern.)


Idempotency-Key prevents duplicate operations when a client retries a request.

- Client sends a unique key (usually UUID) in the `Idempotency-Key` header.
    
- Server stores the key and the result of the first execution.
    
- If the same key arrives again, the server returns the stored result instead of executing again.
    

Example: Prevents a payment from being charged twice due to network retries



# Cookies,Sessions
The entire cookie mechanism is two headers. The server attaches a `Set-Cookie` header to a response; the browser stores the value and then **automatically** includes it in a `Cookie` header on every subsequent request to that domain, without any code on the page asking it to
its headers:
```
HttpOnly; Secure; SameSite=Strict; Max-Age=3600; Path=/
```

`HttpOnly`, the cookie is invisible to JavaScript (`document.cookie` can’t read it).
`Secure`, the cookie is only ever transmitted over HTTPS, so it can’t leak

`SameSite`, controls whether the cookie rides along on cross-site requests, the core defense against **CSRF**

Max-Age -> expires when


#### Authentication
 Session cookies 
 or 
 Tokens
 (JWT)
 When the modified JWT reaches the server:

1. Server reads the header and payload.
    
2. Recalculates the signature using its secret.
    
3. Compares it with the signature attached to the JWT.
    
4. Signatures don't match.
    
5. Rejects the token.![[Pasted image 20261104103830.png]]



# CORS

Browsers enforce the **Same-Origin Policy**: JavaScript on a page from one origin may not read responses from a _different_ origin. An “origin” is the triple of **scheme + host + port**: `https://example.com` and `https://example.com:8443` and `http://example.com` are all different origins

**CORS protects the user’s already-trusted, cookie-carrying session from being read by an origin it never agreed to share it with.**

# Status Code

![[Pasted image 20261105180428.png]]![[Pasted image 20261105180439.png]]![[Pasted image 20261105180445.png]]


301..303 : may drop post + resource
307,308 dont
