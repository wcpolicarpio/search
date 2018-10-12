package main

import "fmt"

func linearsearch(itemlist []int, value int) int {
	for i, item := range itemlist {
		if item == value {
			return i
		}
	}
	return 10
}

func main() {
	items := []int{10, 20, 30, 40, 50}
	fmt.Println(linearsearch(items, 30))
}
