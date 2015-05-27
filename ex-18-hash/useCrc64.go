package main

import (
	"fmt"
	"hash/crc64"
)

func main() {
	var ECMATable = crc64.MakeTable(crc64.ECMA)
	h := crc64.New(ECMATable)
	h.Write([]byte("test"))
	v := h.Sum64()
	fmt.Println(v)
}
