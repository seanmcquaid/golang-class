package main

import "fmt"

func main(){
	// var greetingText string
	// greetingText  = "First app here"

	greetingText := "First app here"

	luckyNumber := 12
	evenMoreLuckyNumber := luckyNumber + 5

	fmt.Println(greetingText)
	fmt.Println(luckyNumber)
	fmt.Println(evenMoreLuckyNumber)

	var newNumber float64 = float64(luckyNumber) / 3

	var defaultFloat float64 = 1.2323232323232332323232323
	var smallFloat float32 = 1.2323232323232332323232323

	fmt.Println(newNumber)
	fmt.Println(defaultFloat)
	fmt.Println(smallFloat)
}