package main

import "fmt"

func main() {
	x := 10

	fmt.Println("\nBasic if:")
	if x > 5 {
		fmt.Println("x is greater than 5")
	}

	fmt.Println("\nif..else:")

	if x > 5 {
		fmt.Println("x is greater than 5")
	} else {
		fmt.Println("x is not greater than 5")
	}

	fmt.Println("\nif-else if-else:")
	if x > 10 {
		fmt.Println("x is greater than 10")
	} else if x == 10 {
		fmt.Println("x is exactly 10")
	} else {
		fmt.Println("x is less than 10")
	}

	fmt.Println("\nWith Initialization Statement:")
	if y := 20; y > 10 {
		fmt.Println("y is greater than 10\n")
	}
	// Note: y is not accessible outside the if block

}
