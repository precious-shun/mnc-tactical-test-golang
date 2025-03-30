package main

import (
	"fmt"
	"math"
)

// Function that takes a string and prints a greeting message
func sayGreeting(name string) {
	fmt.Println("Hello, " + name + "!")
}

// Function that takes a string and prints a goodbye message
func sayGoodbye(name string) {
	fmt.Println("Goodbye, " + name + "!")
}

// Function that takes a slice of strings and a function as arguments
func cycle(n []string, f func(string)) {
	// Iterate over the slice of names and call the function f for each name
	for _, name := range n {
		f(name)
	}
}

func circleArea(r float64) float64 {
	return math.Pi * r * r
}

func main() {
	sayGreeting("Alice")
	sayGoodbye("Bob")

	names := []string{"Vaan", "Penelo", "Balthier", "Fran", "Basch", "Ashe", "Vayne"}
	cycle(names, sayGreeting)
	cycle(names, sayGoodbye)

	a1 := circleArea(10.5)
	a2 := circleArea(15)

	fmt.Println(a1, a2)
	fmt.Printf("Circle 1 is %0.3f and Circle 2 is %0.3f\n", a1, a2)
}
