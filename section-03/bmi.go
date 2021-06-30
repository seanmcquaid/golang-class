package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/seanmcquaid/bmi/info"
)

func main(){
	// Output information
	fmt.Println(info.MainTitle)
	fmt.Println(info.Separator)
	// Prompt for user input (Height + weight)
	fmt.Println(info.WeightPrompt)

	// when enter is pressed, grab the input
	weightInput, _ := reader.ReadString('\n')

	fmt.Println(info.HeightPrompt)
	heightInput, _ := reader.ReadString('\n')

	// Save that user input in variables
	weightInput = strings.Replace(weightInput, "\n", "", -1)
	heightInput = strings.Replace(heightInput, "\n", "", -1)

	weight, _ := strconv.ParseFloat(weightInput, 64)
	height, _ := strconv.ParseFloat(heightInput, 64)

	// Calculate the BMI (weight / (height * height))
	bmi := weight / (height * height)

	// Output the calculated BMI
	fmt.Printf("Your BMI is %.2f", bmi)
}