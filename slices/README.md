# Slices
- Slice is like array, but can shrink or grow their size
- The size can change as we append or remove element(s)
- syntax
```go
var_name := make([]data_type, size, capacity)
```
- `len`: The length of the slice
- `cap`: The capacity, or maximum number of allowed elements in the array
    - if we exceed the `cap` value, it will allocate more `cap` value
    - for example - we set it to 10, and we have input 8 elements at first, the `cap` value will be 10 but if we add 3 more elements, total elements will be 11 then the `cap` value will increase to 20.
- `ptr`: A pointer that points to the address of the underlying array
