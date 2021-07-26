package main

import (
	"fmt"
)

// constant variable
const College_ID = 762159

// METHODS
// struct type definition
type Student struct {
	id float64
	first_name, last_name, address string
	email string
	
}

// Creating an interface
type College interface {
	Student_Name() string
}

// Function to use method
// Student_Name -> method
// s Student -> Receiver
func (s Student) Student_Name() string {
	return s.first_name + " " + s.last_name
}

// POINTERS
func (s *Student) Students_Details() {
	fmt.Printf("Address : %v\nEmail : %v\n", s.address, s.email)
}
// Main Function
func main() {
	fmt.Printf("College ID : %v\n", College_ID)

	// Object of Student
	stu1 := Student{45, "Satyaban", "Sahoo", "Pune, India", "test@test.com"}
	name := stu1.Student_Name()
	fmt.Printf("Name : %v\n", name)

	// Pointer pointing to the Student struct
	stu2 := &stu1
	stu2.Students_Details()
}