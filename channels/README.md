# Channels 
- pipeline that connects goroutines
- pass the values through one goroutine to another goroutine
- syntax
```go
ch := make(chan <type>)         // create a channel
ch <- <value>                   // send a value to the channel
varNew := <- ch                 // Retrieve the value 
```
- unbuffered channel
    - the sender blocks the code to run until the value has been received by the receiver