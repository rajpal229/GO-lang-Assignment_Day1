package main

import (
	"fmt"
)

func main() {
	// Ques1
	var s1, s2, s3 string
	fmt.Println("Enter your favourite below:")
	fmt.Scanln(&s1)
	fmt.Scanln(&s2)
	fmt.Scanln(&s3)
	fmt.Printf("My three Favourite places are: %s, %s, %s\n", s1, s2, s3)
	fmt.Printf("My three Favourite places are:\n%s\n%s\n%s\n", s1, s2, s3)
}
