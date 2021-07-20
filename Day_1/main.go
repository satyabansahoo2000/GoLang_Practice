package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
	"math"
	"time"
)

// Function to read from user
func readName() string {
	
	// define a reader and ask for the input
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Name: ")

	// reader will return a string and error
	// assigning the string to input
	// and error to _ to ignore the variable
	input, _ := reader.ReadString('\n')
	return input
}

// Function to take an string and convert it to integer or float
func convertAge() float64 {
	// define a reader and ask for the input
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter age: ")

	// reader will return a string and error
	// assigning the string to age
	// and error to _ to ignore the variable
	age, _ := reader.ReadString('\n')

	// TrimSpace trims all the spaces from both end
	// ParseFloat then parses the string without spaces
	// and converts it to Float64
	ageFloat, err := strconv.ParseFloat(strings.TrimSpace(age), 64)
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
	return ageFloat
}

// Function to get all the marks and average marks
func arrayMarks() (float64, float64) {

	// making an array of 5 elements 
	marks := make([]float64, 5)
	var totalMarks float64

	// user input of marks
	for i := 0; i < 5; i++ {
		fmt.Printf("Enter marks for Subject %d : ", i+1)
		fmt.Scanln(&marks[i])
	}

	// total marks calculation
	for i := 0; i < 5; i++ {
		totalMarks += marks[i]
	}

	// average marks calculation
	// math.Round to round up the marks to a whole number
	avg := totalMarks / float64(5)
	return totalMarks, math.Round(avg)
}

// Function to set the datetime
func dateAdmission() string {
	// declaring the time Date function to August 3, 2020 9:30:0:0 AM
	// time.Local will set the time for local timezone
	date := time.Date(2020, time.August, 3, 9, 30, 0, 0, time.Local)
	return date.Format(time.ANSIC)
}

// Driver Code
func main() {
	name := readName()
	age := convertAge()
	total, avergae := arrayMarks()
	admission := dateAdmission()

	fmt.Println("\nDETAILS OF THE STUDENT!!!!")
	fmt.Println("--------------------------")
	fmt.Print("Name : ", name)
	fmt.Println("Age : ", age)
	fmt.Println("Total Marks : ", total)
	fmt.Println("Average Marks : ", avergae)
	fmt.Print("Date of Admission : ", admission)
}