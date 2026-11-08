A JavaScript client and a Rust server have nothing in common at the level of data types. Serialization is how they talk anyway


Serialization

Converting **native data** (a JavaScript object, a Rust struct, a Go struct, a Python object) **into the common standard format** before sending it over the network (or storing it).

Deserialization

The exact reverse: taking data _in_ the common format and **parsing it back into the machine’s own native data type**, so it can run its business logic.

![[Pasted image 20261108101653.png]]

![[Pasted image 20261108101924.png]]