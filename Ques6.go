package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// Ques6
	total := 0
	var avg float32
	for i := 0; i < 5; i++ {
		rn1 := rand.Intn(50-10) + 10
		fmt.Printf("%v ", rn1)
		total = total + rn1
	}
	avg = float32(total) / 5
	fmt.Println("")
	fmt.Println(avg)
}
