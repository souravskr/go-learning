package main

import "fmt"

func main() {
	person := map[string]string{
		"firstName": "John",
		"lastName":  "Doe",
	}
	printMap(person)
}

func printMap(c map[string]string) {
	for key, value := range c {
		fmt.Printf("%s : %s", key, value)
		fmt.Println()
		//if key == "lastName" {
		//	delete(c, key)
		//}
	}
}
