package main

import "fmt"

func main() {
	var intSlice = []int{1, 2, 3, 4}
	fmt.Println("SUM", sumSlice(intSlice))

	var float32Slice = []float32{1.1, 2.2, 3.3, 4.4}
	fmt.Println("SUM", sumSlice(float32Slice))
}

func sumSlice[T int | float32 | float64](slice []T) T {

	var sum T
	for _, value := range slice {
		sum += value
	}

	return sum
}
