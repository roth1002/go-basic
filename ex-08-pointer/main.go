package main

import (
	"fmt"
)

/*
The * and & operators
In Go a pointer is represented using the * (asterisk) character followed by the type of the stored value.
In the zeroByRef function xPtr is a pointer to an int.

* is also used to “dereference” pointer variables. Dereferencing a pointer gives us access to the value the pointer points to.
When we write *xPtr = 0 we are saying “store the int 0 in the memory location xPtr refers to”.
If we try xPtr = 0 instead we will get a compiler error because xPtr is not an int it's a *int,
which can only be given another *int.

Finally we use the & operator to find the address of a variable.
&x returns a *int (pointer to an int) because x is an int. This is what allows us to modify the original variable.
&x in main and xPtr in zeroByRef refer to the same memory location.

*/

func main() {
	x := 5
	zero(x)
	fmt.Println(x) // x is still 5

	zeroByRef(&x)
	fmt.Println(x) // x change to 0

	//new
	xPtr := new(int)
	two(xPtr)
	fmt.Println(*xPtr) //xPtr is 2

	y := 1.5
	square(&y)
	fmt.Println(y)

	var (
		a = 2
		b = 1
	)
	swap(&a, &b)
	fmt.Println("a =", a, "b =", b)
}

func zero(x int) {
	x = 0
}

func zeroByRef(xPtr *int) {
	*xPtr = 0
}

func two(xPtr *int) {
	*xPtr = 2
}

func square(x *float64) {
	*x = *x * *x
}

func swap(x *int, y *int) {
	zPtr := new(int)
	*zPtr = *x
	*x = *y
	*y = *zPtr
}
