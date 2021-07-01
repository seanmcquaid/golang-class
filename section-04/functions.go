package main

import "fmt"

func main(){
	a := 5
	b := 10

	sum := add(a, b)
	fmt.Println(sum)
}

func add(a int, b int) int{
	return a + b
}