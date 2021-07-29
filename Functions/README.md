# Variadic Functions
- define a function that accepts a variable number of arguments
- cannot be used at the last

Example - This will show error
```go
func addNums(nums ...int, total int) {
    // do tasks
}
```

# Predefined Functions
- The `map()` function allows you to “map” items from one collection into another collection.
- The `reduce()` function returns a single value based on the collection you pass in.

Go doesn't provide in-built `filter()` function. The `filter()` function takes in a collection of items and returns another collection containing the items you want.