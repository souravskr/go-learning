package main

import "fmt"

type person map[string]string

func main() {
	john := person{"firstName": "John", "lastName": "Doe", "age": "20"}
	for key, value := range john {
		fmt.Printf("%s : %s \n", key, value)
	}
	supplements := map[string]string{"Name": "Magnesium", "Dose": "500mg", "Quantity": "60"}
	fmt.Println(supplements)
	nums := []int{5, 10, 15}
	loopNumbers(nums...)
	fmt.Println(sumAll(nums...))
}

func loopNumbers(nums ...int) {
	fmt.Printf("%T", nums)
	for _, num := range nums {
		fmt.Println(num)
	}
}

func sumAll(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}
