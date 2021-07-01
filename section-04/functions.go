package main

import "fmt"

func main(){
	a := 5
	b := 10

	sum := add(a, b)
	printNumber(sum)
}

func add(a int, b int) int {
	return a + b
}

func printNumber(num int){
	fmt.Printf("The number is %v", num)
}