package main

import "fmt"

func main() {
	fmt.Println("hello!")
}

func mapSearch(target int, m map[int]struct{}) bool {
	_, ok := m[target]
	return ok
}

func sliceSearch(target int, s []int) bool {
	for _, i := range s {
		if target == i {
			return true
		}
	}
	return false
}
