package main

import (
	"fmt"
)

func main() {
	fmt.Println("1.0 + 1.0 = ", 1.0+1.0) //float + float = Integer if after '.' is 0
	fmt.Println("1.0 + 1 = ", 1.0+1)
	fmt.Println(len("Hello World"))                                      //String Length
	fmt.Println("Hello World"[1])                                        // char 'e' in ascii
	fmt.Println("Hello " + "World")                                      // String concat
	fmt.Println(true && true)                                            //true
	fmt.Println(true && false)                                           //false
	fmt.Println(true || true)                                            // true
	fmt.Println(true || false)                                           //true
	fmt.Println(!true)                                                   //false
	fmt.Println((true && false) || (false && true) || !(false && false)) //true
}
