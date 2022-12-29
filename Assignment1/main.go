package main

import "fmt"

func main() {
	sliceOfInts := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, item := range sliceOfInts {
		if item%2 == 0 {
			fmt.Println(item, "even")
		} else {
			fmt.Println(item, "odd")
		}
	}
}
