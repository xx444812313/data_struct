package main

import "data_struct/list_v2"

func main() {
	data := []int{11, 12, 13, 19, 21, 31, 33, 42, 43, 51, 62}
	sl := list_v2.SkipList{}
	sl.InitSkip(data)
	sl.PrintAll()

	sl.Add(1)
	sl.Add(2)
	sl.Add(3)
	sl.PrintAll()

}
