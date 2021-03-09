package main

import (
	"data_struct/list_v2"
	"fmt"
)

func main() {
	data := []int{11, 12, 13, 19, 21, 31, 33, 42, 43, 51, 62}
	sl := list_v2.NewSkipList()
	for _, d := range data {
		sl.Add(d)
	}
	sl.PrintAll()

	sl.Add(1)
	sl.Add(2)
	sl.Add(3)
	sl.PrintAll()

	fmt.Println(sl.Del(21))
	fmt.Println(sl.Del(3))
	sl.PrintAll()

}
