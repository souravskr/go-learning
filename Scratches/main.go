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

type EmpMap map[string]string

func main() {
	animals := []string{"dog", "fish", "cat", "horse"}
	for _, animal := range animals {
		animals = append(animals, animal)
	}
	fmt.Println(animals[0:2])
	sortedAnimals := sortAnimals(animals)
	fmt.Println(sortedAnimals)
	animals = deleteFromSlice(animals, 3)
	fmt.Println(sortedAnimals)
	emp := EmpMap{"name": "Joh Cena", "id": "id_321456", "designation": "Software Engineer", "salary": "80_000"}
	printMap(emp)
	delete(emp, "salary")
	printMap(emp)

}

func sortAnimals(arr []string) []string {
	if sort.StringsAreSorted(arr) {
		return arr
	}
	sort.Strings(arr)
	return arr
}

func deleteFromSlice(arr []string, i int) []string {
	l := len(arr)
	arr[i], arr[l-1] = arr[l-1], arr[i]
	arr[l-1] = ""
	sort.Strings(arr)
	return arr[0:l]
}

func printMap(myMap map[string]string) {
	myMap["id"] = "S252550"
	for _, value := range myMap {
		fmt.Println(value)
	}
}
