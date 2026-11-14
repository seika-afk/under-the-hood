# API DESIGN

**REST, REpresentational State Transfer**

### Representational

A resource (a piece of data, a user, a book, a cart) can be _represented_ in different formats depending on who’s asking. The same `user` record might be sent as **JSON** to an API client (server-to-server), or as **HTML** to a browser,

### State

State is the _current condition_ of a resource, its present attributes. Think of an Amazon shopping cart: its state is the set of items in it, their quantities, and the total price _right now_


### Transfer

Because we have a client and a server, the whole point is moving these representations between them

![[Pasted image 20261113070519.png]]

**Stateless** is the heavy hitter. The server does not remember your previous request, every request must contain all the information needed to understand and process it. This is what makes scaling possible: put a ==load balancer== in front of three identical servers, and because no server holds your session in memory, _any_ of them can handle _any_ request via round-robin



---

### Anatomy of Route


##### Rule 1: always plural nouns

List all books -> `/books`. Fetch one book -> `/books/123`, _not_ `/book/123`. The resource _type_ (“books”) is a collection; the ID just narrows into it

##### Rule 2 :  No spaces or underscores; use slugs

Replace spaces with hyphens: `harry-potter` -> `/books/harry-potter`

##### Rule 3 : the slash means hierarchy
`/organizations/123/projects`


## Method and Idempotency

An operation is idempotent if performing it many times has the _same effect on the server_ as performing it once![[Pasted image 20261114091802.png]]
#### Get

![[Pasted image 20261114091910.png]]

### PUT and PATCH

- **PATCH**: partial update. Send only the fields you want to change (e.g. just `status`). Best fit for SPA-era JSON apps, where you almost always update a few fields, not the whole record.
- **PUT**: full replacement. You must send the _entire_ representation; the server replaces what it has with your payload. Omit a field and you’ve wiped it.

![[Pasted image 20261114092029.png]]



### Delete : IdemPotent

### POST : 
non idempotent , but can make with idempotent keys


## List APIS

### Pagination: why & how

Returning a thousand records at once is expensive on both ends: JSON serialization/deserialization is heavy work, the network bottlenecks, and the user perceives a multi-second delay, even though on screen they only see the first 10-20 items before scrolling. So the server returns the data in **chunks (pages)**.

![[Pasted image 20261114092444.png]]


A paginated response includes four fields:

- `data`, the array of items for this page.
- `total`, the absolute count of items in the DB (so the UI can say “showing 10 of 50”).
- `page`, which page this response represents.
- `totalPages`, the maximum number of pages (the frontend uses this to know when to stop fetching on infinite scroll: when `page === totalPages`, stop).

### Sorting

Two params: `sortBy` (which field) and `sortOrder` (`ascending` / `descending`). Crucial default: even with no sort params, the server must sort by _something_, otherwise the DB returns rows in arbitrary order and the same call yields different orderings each time. The natural default is `sortBy=createdAt`, `sortOrder=descending` (newest first).

### Filtering

Pass resource fields as query params to narrow the list: `?status=active`, or combine them: `?status=archived&name=org`. The server applies each as a filter on the result set.


![[Pasted image 20261114092605.png]]

# Example:

![[Pasted image 20261114092702.png]]


### Be ruthlessly consistent

- **JSON in `camelCase`, always**: both payloads and responses.
- **No abbreviations**: if a field is `description`, never call it `desc` in another endpoint. You have context the consumer doesn’t; an abbreviation they can’t decode costs them a failed call and a docs hunt.
- **Same shapes across resources**: if `organizations` are plural and paginated with `{data,total,page,totalPages}`, then `projects` and `tasks` are too. Consumers integrate one endpoint, then _assume_ the rest follow the same style. Reward that assumption![[Pasted image 20261114092904.png]]


# EIGHT PRINCIPLES FOR RESTFUL API

### 1. Model resources as nouns, not actions
### 2. Use HTTP methods for their intended semantics
### 3. Return meaningful status codes\
### 4. Stay stateless
### 5. Be consistent
### 6. Version your API
### 7. Handle collections gracefully
### 8. Design clear, structured errors