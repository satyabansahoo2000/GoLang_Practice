# Buffered Channels
- allows multiple values to be stored in a channel.
- it will only be blocked when a value in sent to the channel that is full or reading a value from an empty channel
- syntax 
```go
ch := make(chan int, <buffer_size)     // buffered channel
ch := make(chan int, 10)               // buffer channel with buffer size 10, channel can hold upto 10 values before the sending code will block
ch := make(chan int, 0)                // this is same as unbuffered channel
```