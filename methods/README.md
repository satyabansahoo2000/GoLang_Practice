# GO Basics - 3
## Errors
- `error` type is a built-in interface
- returns an error message if a code fails
- handled errors by testing whether the error returns `nil` or error

## Readers
- `io` package provides `io.Reader` interface to read end of a stream of data
- `Read()` interface reads a slice and returns number of bytes assigned and an error
- `Reader` has implementations like `Len`, `Size`, `Read`, `Seek` etc.