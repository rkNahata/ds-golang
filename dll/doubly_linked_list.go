package dll

import "fmt"

type Node struct {
	value interface{}
	next  *Node
	prev  *Node
}

type LinkedList struct {
	head *Node
}

func New() *LinkedList {
	return &LinkedList{nil}
}

func (n *Node) Next() *Node {
	return n.next
}

func (n *Node) Prev() *Node {
	return n.prev
}

func (list *LinkedList) Prepend(value interface{}) *Node {
	newNode := &Node{value, nil, nil}
	newNode.next = list.head
	if list.head != nil {
		list.head.prev = newNode
	}
	list.head = newNode
	return newNode
}

func (list *LinkedList) Append(value interface{}) {
	if list.head != nil {
		newNode := &Node{value: value, next: nil, prev: nil,}
		tail := list.Tail()
		tail.next = newNode
		newNode.prev = tail
	} else {
		list.Prepend(value)
	}
}

func (list *LinkedList) DeleteNode(node *Node) {
	if node.prev != nil {
		//node is not head
		node.prev.next = node.next
	} else {
		list.head = list.head.next
	}
	if node.next != nil {
		node.next.prev = node.prev
	}
}

func (list *LinkedList) DeleteLast() *Node {
	if list.head == nil {
		fmt.Println("empty list!!!")
		return nil
	}
	temp := list.head
	for temp.Next() != nil {
		temp = temp.Next()
	}
	list.DeleteNode(temp)
	return temp
}

func (list *LinkedList) Size() int {
	var size int
	temp := list.head
	if temp == nil {
		return 0
	}
	for temp != nil {
		size++
		temp = temp.Next()
	}
	return size
}

func (list *LinkedList) Tail() *Node {
	temp := list.head
	if temp.Next() == nil {
		return temp
	}
	for temp.Next() != nil {
		temp = temp.Next()
	}
	return temp
}

func (list *LinkedList) LookUp() []int {
	temp := list.head
	var result []int
	for temp != nil {
		result = append(result, temp.value.(int))
		temp = temp.next
	}
	return result
}
