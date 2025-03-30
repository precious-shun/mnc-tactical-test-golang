package main

import (
	"fmt"
)

func main() {
	age := 25

	// Comparison operators
	fmt.Println(age < 50)
	fmt.Println(age >= 50)
	fmt.Println(age == 50)
	fmt.Println(age != 50)

	// if-else statement
	if age < 50 {
		fmt.Println("You are young!")
	} else if age == 50 {
		fmt.Println("You are 50 years old!")
	} else {
		fmt.Println("You are old!")
	}

	// skip the iteration at index 1
	names := []string{"Alice", "Bob", "Charlie"}
	for index, name := range names {
		if index == 1 {
			fmt.Println("Continuing at pos", index)
			continue
		}
		fmt.Printf("The value pos %v is %v\n", index, name)
	}

}
