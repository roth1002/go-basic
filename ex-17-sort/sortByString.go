package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

type ByName []Person

/*
The Sort function in sort takes a sort.Interface and sorts it.
The sort.Interface requires 3 methods: Len, Less and Swap.
To define our own sort we create a new type (ByName) and make it equivalent to a slice of what we want to sort.
We then define the 3 methods.
*/
func (this ByName) Len() int {
	return len(this)
}
func (this ByName) Less(i, j int) bool {
	return this[i].Name < this[j].Name
}
func (this ByName) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func main() {
	clouds := []Person{
		{"tim", 18},
		{"roth", 19},
		{"leo", 20},
	}
	sort.Sort(ByName(clouds))
	fmt.Println(clouds)
}
