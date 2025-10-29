package main

import "fmt"

func main() {
	fmt.Println("Slice:")
	p := []int{1, 2, 3, 4, 5} // slice of integers
	fmt.Println("p=", p)

	fmt.Println("Array:")
	var o [5]int = [5]int{1, 2, 3, 4, 5} // array of five integers
	fmt.Println("o=", o)
}
