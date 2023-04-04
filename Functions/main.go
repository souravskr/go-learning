package main

import "fmt"

type Supplement struct {
	name     string
	dose     []string
	quantity []int
	helps    string
}

type Animal struct {
	name         string
	sound        string
	numberOfLegs int
}

// RECEIVER --> EXTENDS FUNCTIONALITY OF STRUCT
func (a Animal) info() {
	fmt.Printf("%s has %d legs and sounds %s.", a.name, a.numberOfLegs, a.sound)
}

func main() {
	magDoses := []string{"500 mg", "250 mg"}
	magQuantity := []int{60, 80, 120}
	mag := Supplement{
		name:     "Magnesium",
		dose:     magDoses,
		quantity: magQuantity,
		helps:    "Bone Health",
	}
	fmt.Println(mag)
	cat := Animal{"Cat", "meow", 4}
	cat.info()
}
