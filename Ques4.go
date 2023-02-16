package main

import (
	"fmt"
	"strings"
)

func main() {
	//Ques4
	var place string
	fmt.Print("Enter the place you like to visit most: ")
	fmt.Scan(&place)
	U_place := strings.ToUpper(place)
	L_place := strings.ToLower(place)
	fmt.Printf("%s in all Uppercase: %s\nin all Lowercase: %s", place, U_place, L_place)
}
