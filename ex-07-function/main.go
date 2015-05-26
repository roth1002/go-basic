package main

import (
	"fmt"
)

func main() {
	// average
	scores := []float64{98, 93, 77, 82, 83}
	fmt.Println(average(scores))

	//add
	fmt.Println(add(1, 2, 3, 4))

	//also call add
	xs := []int{1, 2, 3, 4, 5}
	fmt.Println(add(xs...))

	//closure
	x := 0
	increment := func() int {
		x++
		return x
	}
	fmt.Println(increment())
	fmt.Println(increment())

	//closure II
	nextEven := makeEvenGenerator()
	fmt.Println(nextEven())
	fmt.Println(nextEven())
	fmt.Println(nextEven())
	fmt.Println(nextEven())

	//Recursion
	fmt.Println(factorial(5))

	//Defer
	/*defer is often used when resources need to be freed in some way.
	For example when we open a file we need to make sure to close it later.
	With defer:

		f, _ := os.Open(filename)
		defer f.Close()

	*/
	defer first()
	second()

	//panic : run-time error usually use with defer save it dose not stop execution of the function
	defer func() {
		str := recover()
		fmt.Println(str)
	}()
	panic("PANIC")
}

func average(scores []float64) (float64, float64) {
	total := 0.0
	for _, value := range scores {
		total += value
	}
	return total / float64(len(scores)), total / float64(len(scores))
}

//Variadic Function
func add(arguments ...int) int {
	total := 0
	for _, value := range arguments {
		total += value
	}
	return total
}

func makeEvenGenerator() func() uint {
	i := uint(0)
	fmt.Println("this only be called once")
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

func factorial(x uint) uint {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}

func first() {
	fmt.Println("first")
}

func second() {
	fmt.Println("second")
}
