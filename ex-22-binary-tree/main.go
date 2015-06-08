package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	fmt.Println("tree=", t)
	left := t.Left
	right := t.Right
	if (left == nil) && (right == nil) {
		fmt.Println("child=", t.Value)
		ch <- t.Value
		return
	}

	if left != nil {
		Walk(left, ch)
	}

	if (left != nil) || (right != nil) {
		fmt.Println("parent=", t.Value)
		ch <- t.Value
	}

	if right != nil {
		Walk(right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)

	depth := 10
	for i := 0; i < depth; i++ {
		x1 := <-ch1
		x2 := <-ch2
		if x1 != x2 {
			return false
		}
	}

	return true
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(tree.New(1), ch1)
	go Walk(tree.New(2), ch2)

	for i := 0; i < 10; i++ {
		fmt.Println(<-ch1)
	}
	for i := 0; i < 10; i++ {
		fmt.Println(<-ch2)
	}
	fmt.Println(Same(tree.New(1), tree.New(2)))
}
