package main

import (
	"fmt"
	"time"
)

func main() {
	//Ques8
	currentTime := time.Now()
	var (
		year  int        = currentTime.Year()
		month time.Month = currentTime.Month()
		day   int        = currentTime.Day()
	)
	fmt.Println("Current Time in String: ", currentTime.String())
	fmt.Println("The year is", currentTime.Year())
	fmt.Println("The month is", currentTime.Month())
	fmt.Println("The day is", currentTime.Day())
	fmt.Println("The date in DD-MM-YYYY: ", currentTime.Day(), "-", currentTime.Month(), "-", currentTime.Year())
	fmt.Printf("%v-%v-%v", day, month, year)
}
