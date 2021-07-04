package main

import "fmt"

func main() {
	age := 32

	fmt.Println(age)

	myAge := &age

	fmt.Println(myAge)
}