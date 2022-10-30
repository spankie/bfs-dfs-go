package main

import "fmt"

// slice of 20 elements
// partition(slice, 4)
func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7}
	num := 2
	res := partition(s, num)
	for i := range s {
		s[i] = 0
	}
	fmt.Printf("Slices: %+v", res)
}

func partition(s []int, num int) [][]int {
	result := [][]int{}
	for i := 0; i < len(s); i = i + num {
		end := i + num
		if end > len(s) {
			end = len(s)
		}
		temp := s[i:end]
		result = append(result, temp)
	}
	return result
}
