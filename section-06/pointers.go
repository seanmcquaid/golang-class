package main

import "fmt"

func main() {
	age := 32

	fmt.Println(age)

	myAge := &age

	*myAge = 33

	fmt.Println(myAge)
	fmt.Println(*myAge)
	fmt.Println(age)
}