package main

import "fmt"

type Employee struct {
	Name   string
	Age    int
	Salary float64
}

func main() {
	smith := Employee{"John Smith", 29, 50000}
	wick := Employee{"John Wick", 32, 60000}
	var employees []Employee
	employees = append(employees, smith)
	employees = append(employees, wick)
	for _, item := range employees {
		if item.Age >= 30 {
			fmt.Println(item.Name, "is 30 or older")
		} else {
			fmt.Println(item.Name, "is under 30")
		}
		if item.Age > 30 && item.Salary > 50000 {
			fmt.Println(item.Name, "make more than 50000 & older than 30")
		} else {
			fmt.Println(item.Name, "make less than 50000 & under 30")
		}
	}
}
