# Encoding Data using JSON
- Encoding a JSON is used to send your data to a web service as a JSON string.
- Encoding is done using `json.Marshal()` function without identation.
- To indent the output and format it nicely, `json.MarshalIndent()` function is being used.
- use `interfaces` instead of `structs` to be able to add additional data as and when you need to.
- Example:
```go
type Customer map[string]interface{}
type Name     map[string]interface{}
type Address  map[string]interface{}
```