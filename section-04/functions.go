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

func add(a int, b int) (sum int) {
	sum = a + b
	return
}

func generateRandomNumbers() (randomNumberOne int, randomNumberTwo int) {
	randomNumberOne = rand.Intn(10)
	randomNumberTwo = rand.Intn(10)
	return
}

func printNumber(num int){
	fmt.Printf("The number is %v", num)
}