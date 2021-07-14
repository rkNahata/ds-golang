package dfs

import "fmt"

type Stack struct {
	Items []int
}

func (s *Stack) Push(item int) {
	s.Items = append(s.Items, item)
}

func (s *Stack) Pop() int {
	if !s.IsEmpty() {
		item := s.Items[len(s.Items)-1]
		s.Items = s.Items[:len(s.Items)-1]
		return item
	}
	return -1
}

func (s *Stack) IsEmpty() bool {
	if len(s.Items) == 0 {
		return true
	}
	return false
}

func New() Stack {
	return Stack{Items: []int{}}
}

//given a number of courses and their schedule, find if the all the courses can be completed
func Traverse(prerequisites [][]int) {
	visited := make([]bool, len(prerequisites))
	s := New()
	graph := createGraph(prerequisites)
	for k, v := range graph {
		s.Push(k)
		for !s.IsEmpty() {
			course := s.Pop()
			if !visited[course] {
				fmt.Print(course)
				visited[course] = true
				for i := range v {
					if visited[v[i]] {
						s.Push(course)
						visited[v[i]] = true
						s.Push(v[i])
						break
					}

				}
			}
		}
	}
	return
}

func createGraph(prerequisites [][]int) map[int][]int {
	graph := make(map[int][]int)
	for i, p := range prerequisites {
		graph[i] = append(graph[i], p...)
	}
	return graph
}
