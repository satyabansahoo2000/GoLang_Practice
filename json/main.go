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

}