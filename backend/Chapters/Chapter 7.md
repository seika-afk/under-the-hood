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












