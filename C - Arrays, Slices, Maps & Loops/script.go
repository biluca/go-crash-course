package main

import (
	"errors"
	"fmt"
)

func main() {
	var intArray = [5]int{0, 1, 2, 3, 4}
	fmt.Println(intArray)

	var intSlice = []int{0, 1, 2, 3, 4}
	fmt.Println(intSlice)
	intSlice = append(intSlice, 5)
	fmt.Println(intSlice)

	var myMap map[string]int = make(map[string]int)
	myMap["price_tag"] = 500
	myMap["price"] = 599
	fmt.Println(myMap)

	var val1, err1 = retrieveValue(myMap, "price_tag")
	fmt.Println(val1, err1)

	var val2, err2 = retrieveValue(myMap, "not-exist")
	fmt.Println(val2, err2)

	for key, value := range myMap {
		fmt.Println(key, ".", value)
	}

	loopFor(9)

}

func loopFor(limit int) {
	for i := 0; i < limit; i++ {
		fmt.Println(i)
	}
}

func retrieveValue(myMap map[string]int, key string) (int, error) {

	var value, ok = myMap[key]
	if ok {
		return value, nil
	} else {
		return -1, errors.New("key not found")
	}

}
