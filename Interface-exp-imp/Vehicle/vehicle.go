package Vehicle

import "fmt"

type Vehicle struct {
	NumberOfWheels int
	NumberOfSeats  int
}

func (v *Vehicle) VehicleDetails() {
	fmt.Println("Number of wheels:", v.NumberOfWheels)
	fmt.Println("Number of seats:", v.NumberOfSeats)
}
