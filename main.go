package main

import (
	"data_struct/list_v2"
)

func main() {
	lru, _ := list_v2.NewLRU(5)
	lru.Set("a1", 1)
	lru.Set("a2", 1)
	lru.Set("a3", 1)
	lru.Set("a4", 1)
	lru.Set("a5", 1)
	lru.PrintAll()
	lru.Get("a3")
	lru.Get("a4")
	lru.Set("a6", 1)
	lru.PrintAll()
}
