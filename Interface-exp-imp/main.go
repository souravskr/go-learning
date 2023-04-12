package main

import (
	"fmt"
	"interface/Vehicle"
	"math/rand"
)

type Car struct {
	Model      string
	Year       int
	IsElectric bool
	IsHybrid   bool
	Vehicle    Vehicle.Vehicle
}

func (c Car) carInfo() {
	fmt.Println("Model:", c.Model)
	fmt.Println("Year", c.Year)
	fmt.Println("IsElectric", c.IsElectric)
	fmt.Println("IsHybrid", c.IsHybrid)
	c.Vehicle.VehicleDetails()
}

func main() {
	suv := Vehicle.Vehicle{NumberOfWheels: 4, NumberOfSeats: 6}
	sedan := Vehicle.Vehicle{NumberOfWheels: 4, NumberOfSeats: 4}

	toyota := Car{"Corolla Cross", 2023, false, true, suv}
	tesla := Car{"X", 2022, true, false, sedan}
	toyota.carInfo()
	fmt.Println("=================")
	tesla.carInfo()

	x := 0
	for x <= 10 {
		fmt.Println(x)
		x = rand.Intn(10) + 2
		if x > 10 {
			fmt.Println("x is greater than 10")
			fmt.Println("Quiting...")
		}
	}
}
