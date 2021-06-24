package main

import "fmt"

func main(){
	firstName := "Sean"
	lastName := "McQuaid"

	fmt.Println(firstName)
	fmt.Println(lastName)

	var currentYear int 
	currentYear = 2021
	birthYear := 1993

	age := currentYear - birthYear

	fmt.Println(age)

	currentYear = currentYear + 1
	age = currentYear - birthYear

	fmt.Println(age)
}