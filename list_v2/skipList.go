//跳表
package list_v2

import "fmt"

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

func showSkipLinkedList(link linkedList, name int) {
	var currentNode *node
	currentNode = link.Head
	for {
		i := 1
		fmt.Print(name, "-node:", currentNode.data)
		if currentNode.next == nil {
			break
		} else {
			currentNode = currentNode.next
		}
		if name == 1 {
			fmt.Print("-------->")
		} else if name == 2 {
			for i <= 3 {
				fmt.Print("-------->")
				i++
			}
		} else {
			for i <= 7 {
				fmt.Print("-------->")
				i++
			}
		}
	}
	fmt.Println("")
}

func (sl *SkipList) InitSkip(list []int) {
	sl.List = linkedList{}
	sl.FirstIndex = linkedList{}
	sl.SecondIndex = linkedList{}
	var currentNode *node
	for i := 0; i < len(list); i++ {
		currentNode = new(node)
		currentNode.data = list[i]
		addNode(sl, currentNode)
		//insertToLink(&sl.List, currentNode)
	}
	showSkipList(*sl)
}

//查找
func (sl *SkipList) Find(x int) bool {
	var current *node
	current = sl.SecondIndex.Head
	if x < current.data {
		return false
	}
	for {
		if x > current.data {
			current = current.next
		} else if x < current.data { //比当前数据节点小，则到下一层查找
			if current.prev.nextLevel == nil { //当前是最底层了，没有找到
				return false
			}
			current = current.prev.nextLevel.next
		} else {
			return true
		}
	}
}

func (sl *SkipList) add(x int) {
	var current *node
	current = sl.SecondIndex.Head
	if current.data == x {
		fmt.Println(x, " Had existed in skipList")
		return
	}
	if x < current.data {
		//
		newNode2 := new(node)
		newNode2.data = x
		newNode2.next = sl.SecondIndex.Head
		sl.SecondIndex.Head.prev = newNode2
		sl.SecondIndex.Head = newNode2

		newNode1 := new(node)
		newNode1.data = x
		newNode1.next = sl.FirstIndex.Head
		sl.FirstIndex.Head.prev = newNode1
		sl.FirstIndex.Head = newNode1

		newNode := new(node)
		newNode.data = x
		newNode.next = sl.List.Head
		sl.SecondIndex.Head.prev = newNode
		sl.List.Head = newNode
		return
	}
	for {
		if x > current.data {
			if current.next == nil {
				if current.nextLevel != nil {
					current = current.nextLevel
				} else {
					//插入
					newNode := new(node)
					newNode.data = x
					current.next = newNode
					newNode.prev = current
					return
				}
			} else {
				current = current.next
			}
		} else if x < current.data {
			//向下去寻找第一个大于x的值
			if current.prev.nextLevel != nil {
				current = current.prev.nextLevel.next
			} else {
				//插入
				newNode := new(node)
				newNode.data = x
				current.prev.next = newNode
				newNode.next = current
				current.prev = newNode
				return
			}
		} else {
			fmt.Println(current.data)
			return
		}
	}
}
func showSkipList(sl SkipList) {
	showSkipLinkedList(sl.SecondIndex, 3)
	fmt.Println("")
	showSkipLinkedList(sl.FirstIndex, 2)
	fmt.Println("")
	showSkipLinkedList(sl.List, 1)
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
	if list.Head == nil { //链表为空
		list.Head, list.Tail = t, t
	} else {
		list.Tail.next = t
		t.prev = list.Tail
		list.Tail = t
	}
	list.Length++
}
