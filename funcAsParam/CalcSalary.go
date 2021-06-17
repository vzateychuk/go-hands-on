package main

import "fmt"

func managerSalary(baseSalary, bonus int) int {
	return baseSalary + bonus
}

func developerSalary(hourlyRate, hoursWorked int) int {
	return hourlyRate * hoursWorked
}

func salary(x int, y int, pay func(i, j int) int) int {
	return pay(x, y)
}

func main() {
	devSalary := salary(50, 2080, developerSalary)
	bossSalary := salary(150000, 25000, managerSalary)
	fmt.Printf("Boss salary: %d\n", bossSalary)
	fmt.Printf("Developer salary: %d\n", devSalary)
}
