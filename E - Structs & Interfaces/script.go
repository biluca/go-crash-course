package main

import "fmt"

type engine interface {
	milesLeft() uint8
}
type gasEngine struct {
	milesPerGallon uint8
	gallons        uint8
}

type eletricEngine struct {
	milesPerKilowatts uint8
	kilowatts         uint8
}

func (g gasEngine) milesLeft() uint8 {
	return g.milesPerGallon * g.gallons
}

func (e eletricEngine) milesLeft() uint8 {
	return e.milesPerKilowatts * e.kilowatts
}

func main() {
	var myGasEngine gasEngine
	myGasEngine.milesPerGallon = 25
	myGasEngine.gallons = 10
	fmt.Println(myGasEngine)

	fmt.Println("Total miles in Gas:", myGasEngine.milesLeft())

	var myEletricEngine eletricEngine
	myEletricEngine.milesPerKilowatts = 34
	myEletricEngine.kilowatts = 9
	fmt.Println(myEletricEngine)

	fmt.Println("Total miles in Eletric:", myEletricEngine.milesLeft())
}
