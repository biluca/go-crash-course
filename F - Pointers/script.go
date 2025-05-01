package main

import "fmt"

func main() {
	var p *int
	var i int = 4

	p = &i

	fmt.Println("Value of i:", i)
	fmt.Println("Address of i:", p)
	fmt.Println("Derreference of p (which is = i):", *p)

	// Changing the Value of i using the pointer p
	*p = 5
	fmt.Println("Value of i:", i)
	fmt.Println("Address of i:", p)
	fmt.Println("Derreference of p (which is = i):", *p)

}
