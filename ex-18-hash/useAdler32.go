package main

import (
	"fmt"
	"hash/adler32"
)

func main() {
	h := adler32.New()
	h.Write([]byte("test"))
	v := h.Sum32()
	fmt.Println(v)
}
