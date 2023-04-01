package main

import (
	"fmt"
	"sort"
)

type Employee struct {
	name        string
	id          int
	designation string
	salary      float64
}

func main() {
	animals := []string{"dog", "fish", "cat", "horse"}
	for _, animal := range animals {
		animals = append(animals, animal)
	}
	fmt.Println(animals[0:2])
	sortedAnimals := sortAnimals(animals)
	fmt.Println(sortedAnimals)
}

func sortAnimals(arr []string) []string {
	if sort.StringsAreSorted(arr) {
		return arr
	}
	sort.Strings(arr)
	return arr
}
