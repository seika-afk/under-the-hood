# Validations & Transformations

- data integrity
- Security

Incoming client data
![[Pasted image 20261110110157.png]]



- Repositery layer
- Service layer
- Controller layer
![[Pasted image 20261110110349.png]]

# Request Life Cycle

![[Pasted image 20261111081236.png]]

![[Pasted image 20261111081311.png]]


![[Pasted image 20261111081512.png]]



# Types of Validation


### 1 Type Validation

```
numberField: expected number, received string
arrayField:  expected array,  received string
boolField:   expected boolean, received string
```


### 2.  Syntactic Validation

Here the type is already correct (a string) but the string must follow a **specific structural pattern**. The validation algorithm checks the _shape_ of the value:

- **Email**: a local part, then the `@` character, then a domain made of a name plus a top-level domain (`.com`, `.co.in`, `.org`, …). There are thousands of valid TLDs; the structure is what gets checked.
- **Phone number**: a country code followed by the right number of digits for that country (e.g. a country code then a 10-digit number).
- **Date**: must match the expected structure, e.g. `YYYY-MM-DD` (year, then month, then day).


### 3. Semantic Validation

The type is right and the syntax is right, but does the value **make sense in the real world?** Semantic checks encode domain logic:

- A **date of birth** cannot be in the future, if today is `2025-01-11`, a DOB of `2025-01-13` is nonsense even though it is a perfectly valid date.
- An **age** of `365` (or `430`) does not make sense, at least not yet. A reasonable human age range is roughly `1` to `120`.


### 4. Complex / Dependent Validation

The most powerful flavor: a field’s rules depend on **other fields**. The pipeline can encode arbitrary cross-field logic:

- **Password confirmation**: `passwordConfirmation` must exactly match `password`; and `password` must be at least 8 characters. Two different strings yield _“passwords don’t match”_; a 6-character password yields _“string must contain at least 8 characters”_.
- **Conditional requirement**: `partner` is optional when `married` is `false`, but the moment `married` is `true` the `partner` name becomes required: _“partner name is required when married is true”_.


![[Pasted image 20261111082127.png]]