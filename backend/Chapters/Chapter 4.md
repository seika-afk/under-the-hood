# Authentication



who are you (authentication) and what can you do (authorization).

![[Pasted image 20261109101043.png]]

## Session
### How a session works

1. **Session creation.** When a user logs in, the server creates a unique **session ID** and stores it alongside the relevant user data (role, name, cart items, whether the user is authenticated) in a **persistent store**: a database, or an in-memory store like Redis.
2. **Delivery.** The session ID, just a key, is sent to the client (the browser) as a **cookie**. Every subsequent request includes that cookie, so the server can use the session ID to fetch all the user data back out of the store.
3. **Expiry.** Sessions are short-lived. If the expiry is 15 minutes, after that the session is invalid; the server creates a fresh session and issues a new session ID. This loop repeats.

![[Pasted image 20261109101236.png]]

The accepted choices are **Argon2id** (preferred today), **scrypt**, and **bcrypt**

Salt: defeat precomputation
 Pepper: a secret the database never holds
 
## JSON WEB TOKEN

![[Pasted image 20261109101636.png]]


## Cookie
A **cookie** is a way to store a piece of information, any string value, in a user’s browser, _from the server side_.


- A cookie is only accessible to the server (domain) that set it, one server cannot see another server’s cookies. This scoping is a built-in security feature of the browser.
- A cookie set by a server is automatically attached to **every subsequent request** back to that same server.


CSRF (Cross-Site Request Forgery) is a web security vulnerability where an attacker tricks an authenticated user's browser into performing an unwanted action on another website.

A CSRF token is a secret, unpredictable value used to verify that a request genuinely originated from the legitimate website's interaction with the user.