package main

import (
	"container/list"
	"fmt"
)

func buildGraph(edges [][]string) map[string][]string {
	graph := map[string][]string{}

	for _, v := range edges {
		a, b := v[0], v[1]
		if _, ok := graph[a]; !ok {
			graph[a] = []string{}
		}
		if _, ok := graph[b]; !ok {
			graph[b] = []string{}
		}
		graph[a] = append(graph[a], b)
		graph[b] = append(graph[b], a)
	}

	return graph
}

func undirectedPath(edges [][]string, src, dst string) bool {
	graph := buildGraph(edges)
	set := map[string]bool{}
	q := list.New()
	q.PushFront(src)
	for e := q.Front(); e != nil; e = e.Next() {
		set[e.Value.(string)] = true
		// fmt.Printf("v: %v", e.Value)
		if e.Value.(string) == dst {
			return true
		}

		for _, v := range graph[e.Value.(string)] {
			if !set[v] {
				q.PushBack(v)
			}
		}
	}
	return false
}

func main() {
	edges := [][]string{
		{"i", "j"},
		{"k", "i"},
		{"m", "k"},
		{"k", "l"},
		{"o", "n"},
	}
	// fmt.Printf("graph: %+v", buildGraph(edges))
	fmt.Printf("is there a path? %v", undirectedPath(edges, "i", "o"))
}
