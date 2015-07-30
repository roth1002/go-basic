package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var a = [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	i := 0
GoTag:
	fmt.Println(a[i])
	i++
	if i < len(a) {
		goto GoTag
	}

	for j := 1; j <= 100; j++ {
		if j%3 == 0 {
			if j%5 == 0 {
				fmt.Println("FizzBuzz")
			} else {
				fmt.Println("Fizz")
			}
		} else if j%5 == 0 {
			fmt.Println("Buzz")
		}
	}

	for a := 1; a <= 100; a++ {
		for b := 0; b < a; b++ {
			fmt.Printf("A")
		}
		fmt.Println()
	}

	str := "dsjkdshdjsdh....js"
	fmt.Printf("String %s\nLength: %d, Runes: %d\n", str,
		len([]byte(str)), utf8.RuneCount([]byte(str)))

	s := "GGGGGGGGGGG"
	r := []rune(s)
	copy(r[4:4+3], []rune("abc"))
	fmt.Printf("Before: %s\n", s)
	fmt.Printf("After : %s\n", string(r))

	foobar := "foobar"
	c := []rune(foobar)
	for i, j := 0, len(c)-1; i < j; i, j = i+1, j-1 {
		c[i], c[j] = c[j], c[i]
	}
	fmt.Printf("%s\n", string(c))

	value, err := abc(100, 200)
	fmt.Println(value, err)
}

func abc(a int, b int) (c int, d error) {
	c = a
	d = nil
	return
}
