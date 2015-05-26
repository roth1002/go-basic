package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x string = "Hello World"
	fmt.Println(x)

	var y string
	y = "Hello World"
	fmt.Println(y)

	var z string
	z = "first"
	fmt.Println(z)
	z = "second"
	fmt.Println(z)
	z = z + " third"
	fmt.Println(z)

	fmt.Println(x == y)
	fmt.Println(z == x)

	//Auto type assign
	a := 5
	fmt.Println(reflect.TypeOf(a))
	b := "Hello World"
	fmt.Println(reflect.TypeOf(b))

	// Const
	const c string = "const string"

	// multiple define
	var (
		d = 5
		e = 10
		f = 15
	)
	fmt.Println("c = ", c)
	fmt.Println("d = ", d)
	fmt.Println("e = ", e)
	fmt.Println("f = ", f)

}
