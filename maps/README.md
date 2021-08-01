# Maps 
- uses key-value pair 
- keys are a set of unqiue-identifiers
- syntax
```go
map[keyValue] valueType
```
- maps are reference type (stores the address where value is stored)
- if the key is not present, the value will return 0
- `delete()` function to delete a key-value pair
- as `delete()` function doesn't return anything, it's better to use `if-statement` to check for the key existance
- syntax
```go
delete(map, key)
```