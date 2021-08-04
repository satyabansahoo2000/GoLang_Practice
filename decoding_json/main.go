package main

import (
	"fmt"
	"encoding/json"
)

type People struct {
	FirstName string
	LastName string
	Details struct {
		Height int
		Weight int
	}
}

// Mapping custom attribute names
type Rates struct {
	Price int `json:"base price"`					// base price maps to Price
	Symbol string `json:"symbol currency"`			// Symbol currency maps to Symbol
}

func main() {
	// decoding JSON to struct
	var person People

	person1 := `{"firstname":"test1", "lastname":"test2"}`
	err := json.Unmarshal([]byte(person1), &person)			// unmarshal the JSON and store in person
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(person.FirstName)
		fmt.Println(person.LastName)
	}

	// decoding JSON to array
	var persons []People
	person2 := 
	`
	[
		{
			"firstname": "test3",
			"lastname": "test4"
		},
		{
			"firstname": "example1",
			"lastname": "example2"
		}
	]
	`
	json.Unmarshal([]byte(person2), &persons)
	for _, p := range persons {
		fmt.Println(p.FirstName)
		fmt.Println(p.LastName)
	}

	// decoding an embedded objects
	var personsDetails []People
	person3 := 
	`
	[
		{
			"firstname":"example3",
			"lastname":"example4",
			"details": {
				"height":190,
				"weight":75
			}
		},
		{
			"firstname":"example5",
			"lastname":"example6",
			"details": {
				"height":185,
				"weight":67
			}
		}
	]
	`
	json.Unmarshal([]byte(person3), &personsDetails)
	for _, pd := range personsDetails {
		fmt.Println(pd.FirstName)
		fmt.Println(pd.LastName)
		fmt.Println(pd.Details.Height)
		fmt.Println(pd.Details.Weight)
	}

	var rate Rates
	person4 := 
	`
	{
		"base price":75,
		"symbol currency":"INR"
	}
	`
	json.Unmarshal([]byte(person4), &rate)
	fmt.Println(rate.Price)
	fmt.Println(rate.Symbol)

	// Mapping Unstructured Data
	var result map[string]interface{}
	priceRate := 
	`
	{
		"success": true,
		"timestamp": 1628085220,
		"base": "EUR",
		"date": "2021-08-04",
		"rates": {
			"AUD": 1.683349,
			"CAD": 1.528643,
			"GBP": 0.874757,
			"SGD": 1.534513,
			"USD": 1.080054
		}
	}
	`
	json.Unmarshal([]byte(priceRate), &result)
	fmt.Println(result["success"])
	
	rates := result["rates"]
	currencies := rates.(map[string]interface{})
	USD := currencies["USD"]
	fmt.Println(USD)
}