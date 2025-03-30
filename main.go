package main

import (
	"fmt"
	"strings"
)

func getInitials(n string) (string, string) {
	s := strings.ToUpper(n)
	names := strings.Split(s, " ")

	var initials []string
	for _, name := range names {
		initials = append(initials, name[0:1])
	}

	if len(initials) > 1 {
		return initials[0], initials[1]
	}

	return initials[0], "_"
}

func main() {
	fn1, sn1 := getInitials("Zhu Yuan")
	fmt.Println(fn1, sn1) // Z Y

	fn1, sn1 = getInitials("Ellen Joe")
	fmt.Println(fn1, sn1) // E J

	fn1, sn1 = getInitials("Burnice")
	fmt.Println(fn1, sn1) // E _
}
