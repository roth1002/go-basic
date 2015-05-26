package main

import (
	"fmt"
  "reflect"
)

func main() {
  fmt.Println("1 + 1 =", 1 + 1)
  fmt.Println("TypeOf(1 + 1) =", reflect.TypeOf(1 + 1))

	fmt.Println("1.0 + 1.0 =", 1.0 + 1.0) //float + float = Integer if after '.' is 0
  fmt.Println("TypeOf(1.0 + 1.0) =", reflect.TypeOf(1.0 + 1.0))

  fmt.Println("1.0 + 1 =", 1.0 + 1)
	fmt.Println("TypeOf(1.0 + 1) =", reflect.TypeOf(1.0 + 1))

  fmt.Println("reflect.TypeOf(\"Hello World\") =", reflect.TypeOf("Hello World"))
	fmt.Println("len(\"Hello World\") =", len("Hello World")) //String Length
  fmt.Println("\"Hello World\"[1] =", "Hello World"[1])  // char 'e' in ascii
	fmt.Println("reflect.TypeOf(\"Hello World\"[1]) =", reflect.TypeOf("Hello World"[1]))
	fmt.Println("\"Hello \" + \"World\" =", "Hello " + "World") // String concat

  fmt.Println("reflect.TypeOf(true) =", reflect.TypeOf(true))  // true
	fmt.Println("true && true =", true && true)  // true
	fmt.Println("true && false =", true && false)  // false
	fmt.Println("true || true =", true || true)  // true
	fmt.Println("true || false =", true || false)  // true
	fmt.Println("!true =", !true) // false
	fmt.Println("(true && false) || (false && true) || !(false && false) =", (true && false) || (false && true) || !(false && false)) //true
}
