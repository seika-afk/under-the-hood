# ROUTING

![[Pasted image 20261107080104.png]]


dynamic  routing

![[Pasted image 20261107080306.png]]



### Query Parameters

A **query parameter** is a key/value pair attached to the end of the route after a **question mark**. Multiple pairs are joined by an ampersand. For example, `/api/search?query=some+value`, the route is `/api/search`, and after the `?` comes the key `query` with the value `some+value`.![[Pasted image 20261107080523.png]]

An optional `key=value` pair in the query string, used to **filter, sort, paginate, or search**. Typically used with GET requests because they have no body. Optional, repeatable, and order-independent'

![[Pasted image 20261107080910.png]]


##### Nesting in Routing
![[Pasted image 20261107080957.png]]



![[Pasted image 20261108100318.png]]



### Why Versioning in API
Why bother? Suppose v1 returns each product as `{ id, name, price }`. Later, new requirements arrive, say you now serve a mobile app and must change the response shape, renaming `name` to `title`. If you mutate the live route, you break every existing client. Versioning lets both shapes live at once:



### Catch All Routes

![[Pasted image 20261108100620.png]]

