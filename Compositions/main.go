package main

import "fmt"

type Vehicle struct {
	NumberOfWheels     int
	NumberOfPassengers int
}

type Car struct {
	Model      string
	Year       int
	IsElectric bool
	IsHybrid   bool
	Vehicle    Vehicle
}

func (v Vehicle) showDetails() {
	fmt.Println("Number of Wheels:", v.NumberOfWheels)
	fmt.Println("Number of Passengers:", v.NumberOfPassengers)
}

func (c Car) carDetails() {
	fmt.Println("Model:", c.Model)
	fmt.Println("Year:", c.Year)
	fmt.Println("IsElectric:", c.IsElectric)
	fmt.Println("IsHybrid", c.IsHybrid)
	c.Vehicle.showDetails()
}

func main() {
	suv := Vehicle{4, 6}
	sedan := Vehicle{4, 4}
	toyotaCorolla := Car{"Toyota Corolla Cross", 2023, false, true, suv}
	tesla := Car{"Tesla X", 2022, true, false, sedan}
	toyotaCorolla.carDetails()
	fmt.Println("===========================")
	tesla.Vehicle.NumberOfPassengers = 5
	tesla.carDetails()
}
