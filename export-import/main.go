package main

import (
	"expImp/staff"
	"fmt"
)

func main() {
	employee := []staff.Employee{
		{"John", "Cena", 40000, true},
		{"John", "Wick", 35000, true},
		{"Bob", "Ulnar", 15000, false},
		{"Sten", "Back", 75000, true},
		{"Hose", "Tortilla", 800000, true},
	}
	allStuff := staff.Office{Employee: employee}
	fmt.Println("All Employees: ", allStuff.All())

	staff.UnderPaymentLimit = 40_000
	staff.OverPaymentLimit = 150_000

	fmt.Println("Under Paid Employees:", allStuff.UnderPaid())
	fmt.Println("Over Paid Employees:", allStuff.OverPaid())
}
