package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/seanmcquaid/bmi/info"
)

var reader = bufio.NewReader(os.Stdin)

func getUserMetrics() (weight float64, height float64){
	// Prompt for user input (Height + weight)
	fmt.Println(info.WeightPrompt)

	// when enter is pressed, grab the input
	weightInput, _ := reader.ReadString('\n')

	fmt.Println(info.HeightPrompt)
	heightInput, _ := reader.ReadString('\n')

	// Save that user input in variables
	weightInput = strings.Replace(weightInput, "\n", "", -1)
	heightInput = strings.Replace(heightInput, "\n", "", -1)

	weight, _ = strconv.ParseFloat(weightInput, 64)
	height, _ = strconv.ParseFloat(heightInput, 64)

	return
}