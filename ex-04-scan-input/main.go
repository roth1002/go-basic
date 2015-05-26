package main

import (
	"fmt"
)

func main() {
	fmt.Println("Enter a number: ")
	var input float64
	fmt.Scanf("%f", &input) // input string will be 0

	output := input * 2
	fmt.Println(output)
}
