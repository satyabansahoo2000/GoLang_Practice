# Getting Started with GoLang

## hello-world program

1. Create a folder named "folder_name" 
2. `cd folder_name`
3. Enable Dependancy tracking by running `go mod init example.com/helloworld`
4. Code the [helloworld](hello.go)
5. Run `go mod tidy` to update the dependencies for [rsc.io/quote](https://pkg.go.dev/rsc.io/quote)
6. To run the code, execute `go run .` which will run the only .go file present inside the folder. 