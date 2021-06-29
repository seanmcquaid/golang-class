package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func main(){
	fmt.Println("Starting app")
	// Output information
	fmt.Println("BMI Calculator")
	fmt.Println("---------------------")
	// Prompt for user input (Height + weight)
	fmt.Println("Please enter your weight (kg): ")

	// when enter is pressed, grab the input
	weightInput, _ := reader.ReadString('\n')

	fmt.Println("Please enter your height (m): ")
	heightInput, _ := reader.ReadString('\n')

	weightInput = strings.Replace(weightInput, "\n", "", -1)
	heightInput = strings.Replace(heightInput, "\n", "", -1)

	weight, _ := strconv.ParseFloat(weightInput, 64)
	height, _ := strconv.ParseFloat(heightInput, 64)


	// Save that user input in variables
	// Calculate the BMI (weight / (height * height))
	// Output the calculated BMI
}