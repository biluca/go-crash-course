package main

import "fmt"

func main() {

	var myString = []rune("Résumé São João")
	fmt.Println(myString)

	for i, v := range myString {
		fmt.Println(":", i, v)
	}

}
