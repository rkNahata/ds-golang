package linkedlist

import (
	"errors"
	"fmt"
)

type Node struct {
	Value interface{}
	Next  *Node
}

type SinglyLinkedList struct {
	Head *Node
}

func New()*SinglyLinkedList{
	return &SinglyLinkedList{nil}
}

func (s *SinglyLinkedList) IsEmpty() bool {
	if s.Head == nil {
		return true
	}
	return false
}

//add an element at the end
func (s *SinglyLinkedList) Append(val interface{}) {
	newNode := Node{
		Value: val,
		Next:  nil,
	}

	if s.Head == nil {
		s.Head = &newNode
	} else {
		lastNode := s.Head
		for {
			if lastNode.Next == nil {
				break
			}
			lastNode = lastNode.Next
		}
		lastNode.Next = &newNode
	}
}

//insert an element at the given index
func (s *SinglyLinkedList) InsertAt(i int, val interface{}) error {
	if i < 0 || i > s.Length() {
		return errors.New("index out of range")
	}
	newNode := Node{
		Value: val,
		Next:  s.Head,
	}
	if i == 0 {
		s.Head = &newNode
		return nil
	}

	j := 0
	currentNode := s.Head
	for j < i-2 {
		currentNode = currentNode.Next
		j++
	}
	tmp := currentNode.Next
	currentNode.Next = &newNode
	newNode.Next = tmp
	return nil
}

func (s *SinglyLinkedList) Length() int {
	var count int
	currentNode := s.Head
	for {
		if currentNode == nil {
			break
		}
		count++
		currentNode = currentNode.Next
	}
	return count
}

func (s *SinglyLinkedList) IndexOf(val interface{}) (bool, int) {
	if !s.IsEmpty() {
		currentNode := s.Head
		index := 1
		for currentNode != nil {
			if currentNode.Value == val {
				return true, index
			}
			currentNode = currentNode.Next
			index++
		}
	}
	return false, 0
}

func (s *SinglyLinkedList) ShowAll() {
	currentNode := s.Head
	for i := 0; i < s.Length(); i++ {
		fmt.Printf("%v ", currentNode.Value)
		currentNode = currentNode.Next
	}
}

func (s *SinglyLinkedList) DeleteAtTop() {
	if s.Head == nil {
		return
	}
	s.Head = s.Head.Next
}

func (s *SinglyLinkedList) RemoveAt(i int) error {
	if i < 0 || i > s.Length() {
		return errors.New("index out of range")
	}
	if i == 0 {
		s.Head = s.Head.Next
		return nil
	}
	j := 0
	currentNode := s.Head
	for j < i-2 {
		currentNode = currentNode.Next
		j++
	}
	currentNode.Next = currentNode.Next.Next
	return nil
}



func (s *SinglyLinkedList) BulkInsert(count int) {
	for count > 0 {
		s.Append(count)
		count--
	}
}
