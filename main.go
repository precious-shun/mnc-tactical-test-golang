package main

import (
	"fmt"
)

func main() {

	// for loop with a single condition
	// x := 0
	// for x < 5 {
	// 	fmt.Println("x is", x)
	// 	x++
	// }

	// for loop with a post statement
	// for i := 0; i < 5; i++ {
	// 	fmt.Println("i is", i)
	// }

	// for loop with a range statement
	names := []string{"Alice", "Bob", "Charlie"}
	// for i, name := range names {
	// 	fmt.Println("Name", i, "is", name)
	// }

	// for loop with a range statement and a blank identifier
	// for i, name := range names {
	// 	fmt.Printf("Name %d is %s\n", i, name)
	// }

	// _ is a blank identifier in Go, used to ignore values
	for _, name := range names {
		fmt.Printf("Name is %s\n", name)
	}
}
