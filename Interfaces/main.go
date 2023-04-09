package main

import "fmt"

type Animal interface {
	says() string
	numberOfLegs() int
}

type Dog struct {
	Name         string
	Sounds       string
	NumberOfLegs int
}

type Cat struct {
	Name         string
	Sounds       string
	NumberOfLegs int
	HasTail      bool
}

func (d *Dog) says() string {
	return d.Sounds
}
func (d *Dog) numberOfLegs() int {
	return d.NumberOfLegs
}

func (c *Cat) says() string {
	return c.Sounds
}

func (c *Cat) numberOfLegs() int {
	return c.NumberOfLegs
}

func riddle(a Animal) {
	riddle := fmt.Sprintf(`This animal has %d legs and says %s. What animal is it?`, a.numberOfLegs(), a.says())
	fmt.Println(riddle)
}

func main() {
	cat := Cat{"Cat", "meow", 4, true}
	dog := Dog{"Dog", "woof", 4}
	riddle(&cat)
	riddle(&dog)
}
