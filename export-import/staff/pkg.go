package staff

var OverPaymentLimit = 100_100
var UnderPaymentLimit = 15_000

type Employee struct {
	FirstName  string
	LastName   string
	Salary     int
	IsFullTime bool
}

type Office struct {
	Employee []Employee
}

func (o *Office) All() []Employee {
	return o.Employee
}

func (o *Office) OverPaid() []Employee {
	var overPaidEmployee []Employee
	for _, emp := range o.All() {
		if emp.Salary > OverPaymentLimit {
			overPaidEmployee = append(overPaidEmployee, emp)
		}
	}
	return overPaidEmployee
}

func (o *Office) UnderPaid() []Employee {
	var underPaidEmployee []Employee
	for _, emp := range o.All() {
		if emp.Salary < UnderPaymentLimit {
			underPaidEmployee = append(underPaidEmployee, emp)
		}
	}
	return underPaidEmployee
}
