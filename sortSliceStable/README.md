# SliceStable Function 
- implementing `SliceStable()` Function from `sort()` function
- `SliceStable()` sorts the slice x using the provided less function, keeping equal elements in their original order.
- syntax
```go
func SliceStable(slice interface{}, less func(i, j int) bool)
```
- sorting by DOB