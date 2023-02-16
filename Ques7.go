package main

import "fmt"

func main() {
	//Ques7
	var first string
	var last string
	fmt.Print("Enter First Name: ")
	fmt.Scan(&first)
	fmt.Print("Enter Last Name: ")
	fmt.Scan(&last)
	fmt.Printf("Full name: %s %s", first, last)
}
