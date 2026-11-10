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

## Statefull Auth

![[Pasted image 20261110104209.png]]


## Stateless Auth

![[Pasted image 20261110104328.png]]


![[Pasted image 20261110104342.png]]



## API KEY
![[Pasted image 20261110104717.png]]


## Delegation Problem

Then a new class of problem appeared: **one website needing access to another website’s resources**. A travel/booking app needs to scan your Gmail for flight tickets. A social app wants to import your contacts from Google or another network. One platform wants another platform’s resource, programmatically. This is the **delegation problem**![[Pasted image 20261110104839.png]]

![[Pasted image 20261110104901.png]]


# OAUTH


![[Pasted image 20261110104937.png]]![[Pasted image 20261110105040.png]]

![[Pasted image 20261110105053.png]]


Around **2014**, **OpenID Connect (OIDC)** was built _on top of_ OAuth 2.0’s security mechanisms to fill the authentication gap. Its core addition is the **ID token**: typically a **JWT** (this is exactly why JWTs were worth understanding early; they reappear everywhere).

The ID token carries identity info: the user ID, when they logged in (the `iat` field), the **issuing authority** (which platform issued the token), and often the user’s name and email, whatever the relying party needs. Because of OIDC, virtually every platform now offers **“Sign in with Google / Facebook / Discord.”** When you do, OIDC works behind the scenes to take your identity (email, profile picture, name) from Google, and the client either stores it or simply uses it to authenticate you, without building its own auth system. OIDC added a layer of **identification** on top of OAuth 2.0

![[Pasted image 20261110105333.png]]



![[Pasted image 20261110105422.png]]


### How RBAC works

Define **roles** (user, admin, moderator, or your own custom roles), and assign each role a set of **permissions** over **resources**. You can go as granular as you like.


```
The runtime flow: a user signs up and the server assigns a role (user, admin, …). On each subsequent request the user sends its identification (a session ID for stateful, or a JWT for stateless). Early in the request cycle the server deduces the role, either from the token’s claims or via a DB lookup, and attaches it so that later middleware/logic can decide access. If the role is `admin`, the dead-zone endpoint is allowed; if the role is `user`, the server returns **403 Forbidden**: meaning you don’t have enough permission for this resource.
```
