package main

import (
	"fmt"
	"math/rand"
)

func main(){
	a, b := generateRandomNumbers()

	sum := add(a, b)
	printNumber(sum)
}

func add(a int, b int) int {
	return a + b
}

func generateRandomNumbers() (int, int) {
	randomNumberOne := rand.Intn(10)
	randomNumberTwo := rand.Intn(10)
	return randomNumberOne, randomNumberTwo
}

func printNumber(num int){
	fmt.Printf("The number is %v", num)
}