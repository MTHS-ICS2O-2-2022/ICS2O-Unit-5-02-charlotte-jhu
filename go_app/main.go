// Created by: Charlotte Jhu
// Created on: May 2023
//
// This program generates a positive or negative number

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// This is the main function
	// variables
	var randomNumber int
	var positiveOrNegative string
	var positive string = "positive"
	var negative string = "negative"

	// input
	fmt.Println("This program generates a positive or negative number")
	fmt.Println()
	fmt.Println("Would you like to generate a positive or negative number? (positive/negative): ")
	fmt.Scanln(&positiveOrNegative)

	// process
	if positiveOrNegative == positive {
	rand.Seed(time.Now().UnixNano())
	min := 1
	max := 6
	randomNumber = rand.Intn(max-min) + min
	} else if positiveOrNegative == negative {
	rand.Seed(time.Now().UnixNano())
	min := -6
	max := -1
	randomNumber = rand.Intn(max-min) + min
	} else {
		fmt.Println("Invalid input")
	}

	// output
	fmt.Println("The random number is: ",randomNumber )
	fmt.Println("\nDone.")
}
