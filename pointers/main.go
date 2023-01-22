package main

import "fmt"

type person struct {
	name string
	age  int
}

func intPerson() *person {
	m := person{"john", 20}
	return &m
}

func main() {
	fmt.Println(*intPerson())
}
