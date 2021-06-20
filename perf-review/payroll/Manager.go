package payroll

type Manager struct {
	Employee
	Salary         float64
	CommissionRate float64
}

func (man Manager) FullName() string {
	return man.FirstName + " " + man.LastName
}

func (man Manager) Pay() (string, float64) {
	payment := man.Salary + (man.Salary * man.CommissionRate)
	return man.FullName(), payment
}
