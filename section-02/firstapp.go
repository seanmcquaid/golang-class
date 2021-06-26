package main

import "fmt"

var greetingText = "First app here"

func main(){
	// var greetingText string
	// greetingText  = "First app here"

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

	var firstRune rune = 'a'
	fmt.Println(firstRune)
	fmt.Println(string(firstRune))

	var firstByte byte = 'a'
	fmt.Println(firstByte)
	fmt.Println(string(firstByte))

	firstName := "Sean"
	lastName := "McQuaid"
	age := 28

	fmt.Printf("Hi I am %v %v and I am %v years old - %T", firstName, lastName, age, age)
	// fmt.Println(int("9"))
}