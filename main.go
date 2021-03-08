package main

import (
	"data_struct/list_v2"
	"fmt"
)

func main() {
	data := []int{11, 12, 13, 19, 21, 31, 33, 42, 51, 62}
	sl := list_v2.SkipList{}
	sl.InitSkip(data)
	fmt.Println(sl.Find(31))
}
