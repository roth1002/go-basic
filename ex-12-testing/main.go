package main

import (
	"fmt"
	"github.com/roth1002/go-basic/ex-12-testing/math"
)

func main() {
	xs := []float64{1, 2, 3, 4}
	avg := math.Average(xs)
	fmt.Println(avg)
}
