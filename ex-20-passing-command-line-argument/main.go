package main

import (
	"flag"
	"fmt"
)

func main() {
	// Define flag
	maxf := flag.Int("max", 6, "the max value")

	// Parse
	flag.Parse()

	fmt.Println(*maxf)

	fmt.Println(flag.Args())
}
