package main

import (
	"fmt"
)

func main() {
	Movie(100, 10, 0.9)
}

// Movie function comment
func Movie(card, ticket int, perc float64) int {
	// Function for getting the number of tickets in which
	// the system B cost less compared to the system A
	var isLess = false
	var priceSystemA float64 = 0
	var priceSystemB float64 = 0
	var exp float64 = 1

	for !isLess {
		priceSystemA = float64(ticket) * exp
		fmt.Print("System A price: ")
		fmt.Println(priceSystemA)
		priceSystemB = float64(card) + float64(ticket)*exp
		fmt.Print("System B price: ")
		fmt.Println(priceSystemB)
		isLess = true
	}

	return 1
}
