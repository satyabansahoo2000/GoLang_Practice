package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type carsInfo struct {
	Model string `json:"Model"`
}

var allCars map[string]carsInfo

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to our Page!!!!")
}

func allcars(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "List of all cars!!!!")

	// Passing in query string
	kv := r.URL.Query()

	for k, v := range kv {
		fmt.Println(k, v)
	}
	// returns all the models in JSON
	json.NewEncoder(w).Encode(allCars)
}

func cars(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	// fmt.Fprintf(w, "Detail for car "+params["carid"])
	// fmt.Fprintf(w, "\n")
	// fmt.Fprintf(w, r.Method)

	// GET is for retrieving cars
	if r.Method == "GET" {
		if _, ok := allCars[params["carid"]]; ok {
			json.NewEncoder(w).Encode(allCars[params["carid"]])
		} else {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("Error 404 : No Model Found"))
		}
	}

	// DELETE is for deleting cars
	if r.Method == "DELETE" {
		if _, ok := allCars[params["carid"]]; ok {
			delete(allCars, params["carid"])
			w.WriteHeader(http.StatusNoContent)
		} else {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("Error 404 : No Model Found"))
		}
	}

	if r.Header.Get("Content-type") == "application/json" {
		
		// POST is for creating new car
		if r.Method == "POST" {
			// read the string sent to the service
			var newCar carsInfo
			reqBody, err := ioutil.ReadAll(r.Body)

			if err == nil {
				// Convert JSON to Object
				json.Unmarshal(reqBody, &newCar)

				if newCar.Model == "" {
					w.WriteHeader(http.StatusUnprocessableEntity)
					w.Write([]byte(
						"Error 422 : Please provide the Car Model" + "information " + "in JSON Format"))
					return
				}
				// check if the model exists; add only if course does not exist
				if _, ok := allCars[params["carid"]]; !ok {
					w.WriteHeader(http.StatusCreated)
					w.Write([]byte(
						"Success 201 : Course Added: " + params["carid"]))
				} else {
					w.WriteHeader(http.StatusConflict)
					w.Write([]byte(
						"Error 409 : Duplicate Course ID found!!!"))
				}
			} else {
				w.WriteHeader(http.StatusUnprocessableEntity)
				w.Write([]byte("Error 422 : Please provide the course information" + "in JSON format..."))
			}
		}

		// PUT is for creating or updating existing car
		if r.Method == "PUT" {
			var newCar carsInfo
			reqBody, err := ioutil.ReadAll(r.Body)

			if err == nil {
				json.Unmarshal(reqBody, &newCar)

				if newCar.Model == "" {
					w.WriteHeader(http.StatusUnprocessableEntity)
					w.Write([]byte(
						"Error 422 : Please provide the Car Model" + "information " + "in JSON Format"))
					return
				}

				// check if course exists; add only if car does not exist
				if _, ok := allCars[params["carid"]]; !ok {
					w.WriteHeader(http.StatusCreated)
					w.Write([]byte(
						"Success 201 : Course Added: " + params["carid"]))
				} else {
					// update course
					allCars[params["carid"]] = newCar
					w.WriteHeader(http.StatusNoContent)
				}
			} else {
				w.WriteHeader(http.StatusUnprocessableEntity)
				w.Write([]byte("Error 422 : Please provide the course information" + "in JSON format..."))
			}
		}
	}
}

func main() {
	allCars = make(map[string]carsInfo)

	router := mux.NewRouter()						// registers routes
	router.HandleFunc("/api/v1/", home).Methods(	// map it to the home() function
		"GET", "POST", "PUT", "DELETE",				// others methods like PATCH, HEAD, OPTIONS wont work
	)												
	router.HandleFunc("/api/v1/cars", allcars)
	router.HandleFunc("/api/v1/cars/{carid}", cars)

	fmt.Println("Listening to port 5000")
	log.Fatal(http.ListenAndServe(":5000", router))
}