package main

import (
	"container/list"
	"fmt"
)

type Stack []string

func (s *Stack) isEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(str string) {
	*s = append(*s, str)
}

func (s *Stack) Pop() (string, bool) {
	if s.isEmpty() {
		return "", false
	}
	index := len(*s) - 1
	value := (*s)[index]
	*s = (*s)[:index]
	return value, true
}

var graph = map[string][]string{
	"a": []string{"c", "b"},
	"b": []string{"d"},
	"c": []string{"e"},
	"d": []string{"f"},
	"e": []string{},
	"f": []string{},
}

func deptFirstSearch(g map[string][]string, source string) {
	stack := Stack{source}
	// fmt.Println(source)
	for len(stack) > 0 {
		current, ok := stack.Pop()
		if !ok {
			break
		}
		fmt.Println(current)
		for _, v := range g[current] {
			stack.Push(v)
		}
	}
	// for _, v := range g[source] {
	// stack = append(stack, v)
	// source = v
	// }
}

func breadthFirstSearch(g map[string][]string, source string) {
	q := list.New()
	q.PushFront(source)

	for e := q.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
		for _, v := range g[e.Value.(string)] {
			q.PushBack(v)
		}
	}
}

func main() {
	// deptFirstSearch(graph, "a")
	breadthFirstSearch(graph, "a")
	/*
		l := list.New()
		l.PushFront("spankie")
		l.PushBack(2)
		for e := l.Front(); e != nil; e = e.Next() {
			fmt.Printf("name: %v\n", e.Value)
		}*/
}
