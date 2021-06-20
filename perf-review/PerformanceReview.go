package main

import (
	"fmt"
	"os"
	"perf-review/payroll"
)

var (
	dev     payroll.Developer
	manager payroll.Manager
)

func init() {
	fmt.Println("Welcome to the Employee Pay and Performance Review\n++++++++++++++++++++++++++++++++++++++++++++++++++")
}

func init() {
	fmt.Println("Initializing variables")
	var employeeReview map[string]interface{} = map[string]interface{}{
		"WorkQuality":     5,
		"TeamWork":        2,
		"Communication":   "Poor",
		"Problem-solving": 4,
		"Dependability":   "Unsatisfactory",
	}

	vzateychuk := payroll.Employee{
		Id:        10,
		FirstName: "Vladimir",
		LastName:  "Zateychuk",
	}
	dev = payroll.Developer{vzateychuk, 35, 2400, employeeReview}

	var ivanovich = payroll.Employee{
		Id:        20,
		FirstName: "Ivan Ivanovich",
		LastName:  "Ivanov",
	}
	manager = payroll.Manager{
		Employee:       ivanovich,
		Salary:         150000.00,
		CommissionRate: 0.7,
	}
}

func main() {
	err := dev.ReviewRating()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	payroll.PrintPayerDetails(dev)
	payroll.PrintPayerDetails(manager)
}
