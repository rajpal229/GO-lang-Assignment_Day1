package main

import (
	"fmt"
)

func main() {
	//Ques2
	const pi float32 = 3.14
	var dia float32
	fmt.Println("Enter Diameter:")
	fmt.Scan(&dia)
	perimeter := pi * dia
	fmt.Printf("Perimeter for Diameter %v is %v", dia, perimeter)
}
