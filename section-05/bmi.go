package main

import (
	"fmt"

	"github.com/seanmcquaid/bmi/info"
)

func main(){
	info.PrintWelcome()

	weight, height := getUserMetrics()

	// Calculate the BMI (weight / (height * height))
	bmi := weight / (height * height)

	// Output the calculated BMI
	fmt.Printf("Your BMI is %.2f", bmi)
}