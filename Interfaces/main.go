package main

import "fmt"

type Animal interface {
	says() string
	numberOfLegs() int8
}

type Dog struct {
	Name         string
	Sounds       string
	NumberOfLegs int8
}

type Cat struct {
	Name         string
	Sounds       string
	NumberOfLegs int8
	HasTail      bool
}

func main() {
	dog := Dog{"Dog", "woof", 4}
	cat := Cat{"Cat", "meow", 4, true}
	riddle(&dog)
	riddle(&cat)
}

func riddle(a Animal) {
	riddle := fmt.Sprintf(`The animal has %d and it says %s, what animal is it?`, a.numberOfLegs, a.says)
	fmt.Println(riddle)
}

func (d *Dog) says() string {
	return d.Sounds
}

func (d *Dog) numberOfLegs() int8 {
	return d.NumberOfLegs
}

func (c *Cat) says() string {
	return c.Sounds
}

func (c *Cat) numberOfLegs() int8 {
	return c.NumberOfLegs
}
