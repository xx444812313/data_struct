/**
跳表实现（两层索引）
 */
package list_v2

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"time"
)

//跳表对象(有两层索引）
type SkipList struct {
	SecondIndex linkedList //二级索引（顶层索引）
	FirstIndex  linkedList //一级索引
	List        linkedList //数据链表
}

//链表对象
type linkedList struct {
	Head   *skipNode
	Tail   *skipNode
	Length int
	Level  int
}

//数据节点对象
type skipNode struct {
	data      int //数据
	next      *skipNode
	prev      *skipNode
	nextLevel *skipNode //指向下一层
}

var ARROW_CHAR = "--"
var RAND *rand.Rand

//打印跳表
func (sl *SkipList) PrintAll() {
	twoList := make([]string, 0) //二级索引
	oneList := make([]string, 0) //一级索引
	list := make([]string, 0)    //数据层

	//数据层
	current := sl.List.Head
	for current != nil {
		list = append(list, current.value())
		current = current.next
		list = append(list, ARROW_CHAR)
	}

	//一级索引
	current = sl.List.Head
	a := sl.FirstIndex.Head
	for current != nil {
		if a != nil && current == a.nextLevel {
			oneList = append(oneList, a.value())
			a = a.next
		} else {
			oneList = append(oneList, ARROW_CHAR)
		}
		current = current.next
		oneList = append(oneList, ARROW_CHAR)
	}

	//二级索引
	current = sl.List.Head
	b := sl.SecondIndex.Head
	for current != nil {
		if b != nil && current == b.nextLevel.nextLevel {
			twoList = append(twoList, b.value())
			b = b.next
		} else {
			twoList = append(twoList, ARROW_CHAR)
		}
		current = current.next
		twoList = append(twoList, ARROW_CHAR)
	}

	fmt.Println(twoList)
	fmt.Println(oneList)
	fmt.Println(list)
}

func NewSkipList() SkipList {
	sl := SkipList{}
	sl.List = initList()
	sl.FirstIndex = initList()
	sl.FirstIndex.Head.nextLevel = sl.List.Head
	sl.SecondIndex = initList()
	sl.SecondIndex.Head.nextLevel = sl.FirstIndex.Head
	RAND = rand.New(rand.NewSource(time.Now().Unix()))
	return sl
}

func initList() linkedList {
	res := linkedList{}
	res.Head = &skipNode{data: math.MinInt32}
	res.Tail = res.Head
	return res
}

//查找
func (sl *SkipList) Find(x int) bool {
	_, _, val := sl.find(x)
	if val.data == x {
		return true
	} else {
		return false
	}
}

//查找(二级索引、一级索引、数据节点）
func (sl *SkipList) find(x int) (*skipNode, *skipNode, *skipNode) {
	var a, b, c *skipNode
	a = sl.SecondIndex.Head
	for {
		if x > a.data {
			if a.next == nil {
				break
			}
			a = a.next
		} else if x < a.data {
			a = a.prev //退到上一个小于x的节点
			break
		} else {
			return a, a.nextLevel, a.nextLevel.nextLevel
		}
	}
	b = a.nextLevel
	for {
		if x > b.data {
			if b.next == nil {
				break
			}
			b = b.next
		} else if x < b.data {
			b = b.prev
			break
		} else {
			return a, b, b.nextLevel
		}
	}
	c = b.nextLevel
	for {
		if x > c.data {
			if c.next == nil {
				break
			}
			c = c.next
		} else if x < c.data {
			c = c.prev
			break
		} else {
			break
		}
	}
	return a, b, c
}

func (sl *SkipList) Add(x int) bool {
	a, b, c := sl.find(x)
	if c.data == x {
		return false
	}

	newNode := &skipNode{data: x, prev: c, next: c.next}
	if c.next != nil {
		c.next.prev = newNode
	}
	c.next = newNode

	indexNew := RAND.Intn(21)
	if indexNew < 10 {
		indexNode := &skipNode{data: x, prev: b, next: b.next, nextLevel: newNode}
		if b.next != nil {
			b.next.prev = indexNode
		}
		b.next = indexNode
		if indexNew < 5 {
			twoIndexNode := &skipNode{data: x, prev: a, next: a.next, nextLevel: indexNode}
			if a.next != nil {
				a.next.prev = twoIndexNode
			}
			a.next = twoIndexNode
		}
	}
	return true
}

func (sl *SkipList) Del(x int) bool {
	a, b, c := sl.find(x)
	if c.data != x {
		return false
	}
	c.prev.next = c.next
	if c.next != nil {
		c.next.prev = c.prev
	}

	if b.data == x { //删除一级索引
		b.prev.next = b.next
		if b.next != nil {
			b.next.prev = b.prev
		}
	}

	if a.data == x { //删除二级索引
		a.prev.next = a.next
		if a.next != nil {
			a.next.prev = a.prev
		}
	}
	return true
}

//链表插入新节点（尾插法）
func (list *linkedList) addDataNode(t *skipNode) {
	list.Tail.next = t
	t.prev = list.Tail
	list.Tail = t
	list.Length++
}

func (n skipNode) value() string {
	val := strconv.Itoa(n.data)
	if 0 <= n.data && n.data < 10 {
		return "0" + val
	}
	return val
}

//func main() {
//	data := []int{11, 12, 13, 19, 21, 31, 33, 42, 43, 51, 62}
//	sl := list_v2.NewSkipList()
//	for _,d := range data{
//		sl.Add(d)
//	}
//	sl.PrintAll()
//
//	sl.Add(1)
//	sl.Add(2)
//	sl.Add(3)
//	sl.PrintAll()
//
//
//	fmt.Println(sl.Del(21))
//	fmt.Println(sl.Del(3))
//	sl.PrintAll()
//
//}
