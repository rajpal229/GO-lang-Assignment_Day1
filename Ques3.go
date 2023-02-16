package main

import "fmt"

func main() {
	//Ques3
	var c float32
	fmt.Print("Enter Celcius: ")
	fmt.Scan(&c)
	f := 1.8*c + 32
	fmt.Printf("%v degree Celcius is %v degree Fahrenheit", c, f)
}
