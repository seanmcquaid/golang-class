package main

import (
	"github.com/seanmcquaid/bmi/info"
)

func main(){
	info.PrintWelcome()

	weight, height := getUserMetrics()

	bmi := calcBmi(weight, height)

	printBmi(bmi)
}

func calcBmi(weight float64, height float64) float64 {
	return weight / (height * height)
}