package main

import "fmt"

func main() {
	// Arrays
	fmt.Println("------------Arrays------------")
	var ages [3]int = [3]int{10, 20, 30}
	var shortenedAges = [3]int{10, 20, 30}

	fmt.Println(ages, len(ages), cap(ages))
	fmt.Println(shortenedAges, len(shortenedAges), cap(shortenedAges))

	names := [4]string{"John", "Doe", "Jane", "Smith"}
	fmt.Println(names, len(names), cap(names))

	// Slices
	fmt.Println("------------Slices------------")
	var scores = []int{100, 90, 80}
	scores[2] = 25
	scores = append(scores, 65)

	fmt.Println(scores, len(scores), cap(scores))

	// Slice Range
	fmt.Println("------------Slice Range------------")
	rangeOne := names[1:3]
	fmt.Println(rangeOne, len(rangeOne), cap(rangeOne))
	rangeTwo := names[2:]
	fmt.Println(rangeTwo, len(rangeTwo), cap(rangeTwo))
	rangeThree := names[:3]
	fmt.Println(rangeThree, len(rangeThree), cap(rangeThree))
	rangeOne = append(rangeOne, "New")
	fmt.Println(rangeOne, len(rangeOne), cap(rangeOne))
}
