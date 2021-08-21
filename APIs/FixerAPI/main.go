package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"encoding/json"
)

/*
type Result struct {
	Success   bool
	TimeStamp int
	Base      string
	Date      string
	Rates     map[string]float64
}

type Error struct {
	Success bool
	Error   struct {
		Code int
		Type string
		Info string
	}
}

func main() {
	url := "http://data.fixer.io/api/latest?access_key=<api_key>"
	if resp, err := http.Get(url); err == nil {
		defer resp.Body.Close()
		if body, err := ioutil.ReadAll(resp.Body); err == nil {
			var result Result
			json.Unmarshal([]byte(body), &result)
			if result.Success {
				// create an array to store all keys
				keys := make([]string, 0, len(result.Rates))

				// get all the keys
				for k := range result.Rates {
					keys = append(keys, k)
				}
				// sort the keys
				sort.Strings(keys)

				// print the keys and values in alphabetical order
				for _, k := range keys {
					fmt.Println(k, result.Rates[k])
				}
				// Print only the Rates value
				// for i, v := range result.Rates {
				// 	fmt.Println(i, v)
				// }
			} else {
				var err Error
				json.Unmarshal([]byte(body), &err)
				fmt.Println(err.Error.Info)
			}
			// to print all the info 
			// fmt.Println(string(body))
		} else {
			log.Fatal(err)
		}
	} else {
		log.Fatal(err)
	}
	// fmt.Println("Done")
}
*/

/*
// Refactoring the code for decoding JSON data
var apis map[int]string

func fetchData(API int)  {
	url := apis[API]									// function to reference the URL
	if resp, err := http.Get(url); err == nil {
		defer resp.Body.Close()
		if body, err := ioutil.ReadAll(resp.Body);
		err == nil {
			var result map[string]interface{}			// variable for unmarshal the JSON into a mao of strings keys with values
			json.Unmarshal([]byte(body), &result)
			switch API {
			case 1: // for the Fixer API
			if result["success"] == true {
				fmt.Println(result["rates"].(
					map[string]interface{})["USD"])		// type assertion and extracf the value bsaed on a specific key
			} else {
				fmt.Println(result["error"].(
					map[string]interface{})["info"])
			}
			case 2: // for the openweathermap.org API
			if result["main"] != nil {
				fmt.Println(result["main"].(
					map[string]interface{})["temp"])
			} else {
				fmt.Println(result["message"])
			}
			}
		} else {
			log.Fatal(err)
		}
	} else {
		log.Fatal(err)
	}
}

func main() {
	apis = make(map[int]string)

	apis[1] = "http://api.openweathermap.org/data/2.5/weather?" + "q=<city_name>&appid=<api_key>"
	apis[2] = "http://api.openweathermap.org/data/2.5/weather?q=Talcher&appid=" + "<api_key>"

	go fetchData(1)
	go fetchData(2)

	fmt.Scanln()
}
*/

// Channels returning results

// Channel to store map[int]interface{}
var ch chan map[int]interface{}

var apis map[int]string

func fetchData(API int)  {
	url := apis[API]									// function to reference the URL
	if resp, err := http.Get(url); err == nil {
		defer resp.Body.Close()
		if body, err := ioutil.ReadAll(resp.Body);
		err == nil {

			var result map[string]interface{}			// variable for unmarshal the JSON into a mao of strings keys with values
			json.Unmarshal([]byte(body), &result)

			var re = make(map[int]interface{})
			
			switch API {
			case 1: // for the Fixer API
			if result["success"] == true {
				re[API] = result["rates"].(
					map[string]interface{})["USD"]		// type assertion and extracf the value bsaed on a specific key
			} else {
				re[API] = result["error"].(
					map[string]interface{})["info"]
			}

			// store the result in the channel
			ch <- re 
			fmt.Println("Result of API 1 is passed!!!")

			case 2: // for the openweathermap.org API
			if result["main"] != nil {
				re[API] = result["main"].(
					map[string]interface{})["temp"]
			} else {
				re[API] = result["message"]
			}

			// store the result in the channel
			ch <- re 
			fmt.Println("Result of API 2 is passed!!!")

			}
		} else {
			log.Fatal(err)
		}
	} else {
		log.Fatal(err)
	}
}

func main() {

	// creates a channel to store the result
	ch = make(chan map[int]interface{})

	apis = make(map[int]string)

	apis[1] = "http://api.openweathermap.org/data/2.5/weather?" + "q=SINGAPORE&appid=622c08dd02664b0fa2578d72ca00e04e"
	apis[2] = "http://api.openweathermap.org/data/2.5/weather?q=Talcher&appid=" + "622c08dd02664b0fa2578d72ca00e04e"

	go fetchData(1)
	go fetchData(2)

	// results from two channels
	for i := 0; i < 2; i++ {
		fmt.Println(<- ch)
	}
	fmt.Println("Done!!")
	fmt.Scanln()
}