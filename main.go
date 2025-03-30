package main

import "fmt"

func main() {

	//strings
	var nameOne string = "Shun"
	var nameTwo = "Arita"
	var nameThree string

	fmt.Println(nameOne, nameTwo, nameThree)

	nameOne = "Shun Arita"
	nameThree = "Arita Shun"

	fmt.Println(nameOne, nameThree)

	nameFour := "Shun"

	fmt.Println(nameFour)

	//integers
	var ageOne int = 20
	var ageTwo = 30
	ageThree := 40

	fmt.Println(ageOne, ageTwo, ageThree)

	//bits and memory
	var numOne int8 = 25
	var numTwo int8 = -128
	var numThree uint8 = 255
	fmt.Println(numOne, numTwo, numThree)

	//floats
	var scoreOne float32 = 25.98
	var scoreTwo float64 = 238417098751.238417098751
	scoreThree := 238417098751.238417098751
}
