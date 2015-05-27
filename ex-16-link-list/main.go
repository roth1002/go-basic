package main

import (
	"container/list"
	"fmt"
)

func main() {
	var tim list.List
	tim.PushBack("Tim")
	tim.PushBack("是")
	tim.PushBack("一個")
	tim.PushBack("詩人")

	for e := tim.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
