package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Name struct {
	FirstName string
	LastName  string
}

type Address struct {
	Line1 string
	Line2 string
	Line3 string
}

type Customer struct {
	Name    Name
	Email   string
	Address Address
	DOB     time.Time
}

func main() {
	// Encoding Structs to JSON
	layout := "2021-01-20"
	dob, _ := time.Parse(layout, "2000-02-28")
	customer1 := Customer{
		Name: Name{
			FirstName: "Test",
			LastName:  "Example",
		},
		Email: "test@example.com",
		Address: Address{
			Line1: "road 123",
			Line2: "street 45",
			Line3: "near shop",
		},
		DOB: dob,
	}

	output, err := json.MarshalIndent(customer1, "", "    ")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(output))
	}

	// Encoding interfaces to JSON
	
}
