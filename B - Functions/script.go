package main

import (
	"errors"
	"fmt"
)

func main() {
	printMe("I am Printed!")

	var numerator = 10
	var denominator = 2
	var result, rest, err = intDivision(numerator, denominator)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Result of division: %v with rest: %v", result, rest)
	}
}

func printMe(printValue string) {
	fmt.Println(printValue)
}

func intDivision(numerator int, denominator int) (int, int, error) {
	if denominator == 0 {
		return 0, 0, errors.New("division by zero")
	}
	return numerator / denominator, numerator % denominator, nil
}
