//跳表实现（两层索引）
package list_v2

import (
	"fmt"
	"math"
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
	Head   *node
	Tail   *node
	Length int
	Level  int
}

//数据节点对象
type node struct {
	data      int //数据
	next      *node
	prev      *node
	nextLevel *node //指向下一层
}

var ARROW_CHAR = "--"

func (sl *SkipList) PrintAll() {
	twoList := make([]string, 0) //二级索引
	oneList := make([]string, 0) //一级索引
	list := make([]string, 0)    //数据层

	//数据层
	current := sl.List.Head
	for current != nil {
		list = append(list, fmt.Sprintf("%v", current.data))
		current = current.next
		list = append(list, ARROW_CHAR)
	}

	//一级索引
	current = sl.List.Head
	a := sl.FirstIndex.Head
	for current != nil {
		if a != nil && current == a.nextLevel {
			oneList = append(oneList, fmt.Sprintf("%v", a.data))
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
			twoList = append(twoList, fmt.Sprintf("%v", b.data))
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

func (sl *SkipList) InitSkip(list []int) {
	sl.List = initList()
	sl.FirstIndex = initList()
	sl.FirstIndex.Head.nextLevel = sl.List.Head
	sl.SecondIndex = initList()
	sl.SecondIndex.Head.nextLevel = sl.FirstIndex.Head

	var currentNode *node
	for i := 0; i < len(list); i++ {
		currentNode = new(node)
		currentNode.data = list[i]
		addNode(sl, currentNode)
	}
}

func initList() linkedList {
	res := linkedList{}
	res.Head = &node{data: math.MinInt32}
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

//查找
func (sl *SkipList) find(x int) (*node, *node, *node) {
	var a, b, c *node
	a = sl.SecondIndex.Head
	for {
		if x > a.data {
			if a.next == nil {
				break
			}
			a = a.next
		} else if x < a.data {
			a = a.prev
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

func (sl *SkipList) Add(x int) {
	_, b, c := sl.find(x)
	if c.data == x {
		return
	}

	newNode := &node{data: x, prev: c, next: c.next}
	c.next.prev = newNode
	c.next = newNode

	flag := time.Now().Second() & 1
	if flag == 1 {
		secondNewNode := &node{data: x, prev: b, next: b.next, nextLevel: newNode}
		b.next.prev = secondNewNode
		b.next = secondNewNode
	}
}

func (sl *SkipList) Del(x int) {

}

func addNode(skipList *SkipList, t *node) {
	skipList.List.addDataNode(t)
	if skipList.FirstIndex.Length == 0 || ((skipList.List.Length-1)%2 == 0 && skipList.List.Length > 2) {
		newNode := new(node)
		newNode.data = t.data
		newNode.nextLevel = t
		skipList.FirstIndex.addDataNode(newNode)
		if skipList.SecondIndex.Length == 0 || ((skipList.FirstIndex.Length-1)%2 == 0 && skipList.FirstIndex.Length > 2) {
			newNode2 := new(node)
			newNode2.data = t.data
			newNode2.nextLevel = newNode
			skipList.SecondIndex.addDataNode(newNode2)
		}
	}
}

//链表插入新节点（尾插法）
func (list *linkedList) addDataNode(t *node) {
	list.Tail.next = t
	t.prev = list.Tail
	list.Tail = t
	list.Length++
}
