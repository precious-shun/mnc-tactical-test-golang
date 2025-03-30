package main

import "fmt"

func main() {
	name := "Shun"
	age := 18

	// Print
	fmt.Print("---------Print---------\n")
	fmt.Print("Hello, \n")
	fmt.Print("World!\n")

	// Println
	fmt.Println("---------Println---------")
	fmt.Println("Hello, Shun")
	fmt.Println("Goodbye, Shun!")
	fmt.Println("My age is", age, "years old. And my name is", name, ".")

	// Printf (formatted print)
	fmt.Println("---------Printf---------")
	fmt.Printf("Hello my name is %v, and my age is %v\n", name, age)
	fmt.Printf("Hello my name is %q, and my age is %q\n", name, age)
	fmt.Printf("Hello my name is %T, and my age is %T\n", name, age)
	fmt.Printf("My score is %f\n", 3.14159265358979323846)

	// Sprintf (save formatted string)
	fmt.Println("---------Sprintf---------")
	var str = fmt.Sprintf("Hello my name is %v, and my age is %v", name, age)
	fmt.Printf("The saved string is %q", str)
}
