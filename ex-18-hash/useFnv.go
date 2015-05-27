package main

import (
	"fmt"
	"hash/fnv"
)

func main() {
	h := fnv.New64a()
	h.Write([]byte("test"))
	v := h.Sum64()
	fmt.Println(v)
}
