package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to our Page!!!!")
	fmt.Fprintf(w, "\n")
	fmt.Fprintf(w, r.Method)
}

func allcars(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "List of all cars!!!!")

	// Passing in query string
	kv := r.URL.Query()

	for k, v := range kv {
		fmt.Println(k, v)
	}
}

func cars(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	fmt.Fprintf(w, "Detail for car "+params["carid"])
}



func main() {
	router := mux.NewRouter()						// registers routes
	router.HandleFunc("/api/v1/", home).Methods(	// map it to the home() function
		"GET", "POST", "PUT", "DELETE",				// others methods like PATCH, HEAD, OPTIONS wont work
	)												
	router.HandleFunc("/api/v1/cars", allcars)
	router.HandleFunc("/api/v1/cars/{carid}", cars)

	fmt.Println("Listening to port 5000")
	log.Fatal(http.ListenAndServe(":5000", router))
}