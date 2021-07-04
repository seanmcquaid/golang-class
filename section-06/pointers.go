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

	doubledAge := double(age)
	fmt.Println(doubledAge)
	fmt.Println(age)
}

func double(num int) int {
	result := num * 2
	num = 100
	return result
}