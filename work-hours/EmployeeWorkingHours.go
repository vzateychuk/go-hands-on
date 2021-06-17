/*
Пример использования enum (Weekday) и методов структур
*/
package main

import "fmt"

type employee struct {
	id        int
	firstName string
	lastName  string
}

type Developer struct {
	employee
	hourlyRate int
	workHours  [7]int
}

func (dev *Developer) getFullName() string {
	return dev.firstName + " " + dev.lastName
}
func (dev *Developer) LogHours(day Weekday, hours int) *Developer {
	dev.workHours[day] = hours
	return dev
}
func (dev *Developer) HoursWorked() (hours int) {
	totalHours := 0
	for _, hourPerDay := range dev.workHours {
		totalHours += hourPerDay
	}
	return totalHours
}

/*
Calculate the weekly pay, taking into consideration overtime pay.
Overtime calculates as twice the hourly rate for hours greater than 40.

Returns:
- totalHours - total hours worked;
- totalPay - total pay;
- overtimeFlag - true if the pay includes overtime;
*/
func (dev *Developer) PayDay() (totalHours int, totalPay int, overtimeFlag bool) {
	hours := dev.HoursWorked()
	hoursOver := hours - 40
	isOvertime := hoursOver > 0
	if isOvertime {
		return hours, (40 + 2*hoursOver) * dev.hourlyRate, isOvertime
	} else {
		return hours, hours * dev.hourlyRate, isOvertime
	}
}

/* Выводит на экран workHours за каждый день */
func (dev *Developer) PrintWorkHourDetails() {
	for idx, hours := range dev.workHours {
		fmt.Printf("Hours worked: %v on %v: \n", hours, Weekday(idx).String())
	}
}

/* Create function which will calculate the hours of the employee that have not been logged. */
func funcNonLoggedHours() func(int) int {
	total := 0
	return func(hours int) int {
		total += hours
		return total
	}
}

func main() {
	// создаем developer с именем Tony Stark
	tonyStark := employee{0, "Tony", "Stark"}
	dev := Developer{tonyStark, 20, [7]int{}}

	getNonLoggedHours := funcNonLoggedHours()
	fmt.Println("Tracking hours worked thus far today: ", getNonLoggedHours(2))
	fmt.Println("Tracking hours worked thus far today: ", getNonLoggedHours(3))
	fmt.Println("Tracking hours worked thus far today: ", getNonLoggedHours(5))
	fmt.Println()

	// заполняем табель рабочего времени
	dev.LogHours(Monday, 8)
	dev.LogHours(Tuesday, 10)
	dev.LogHours(Wednesday, 10)
	dev.LogHours(Thursday, 10)
	dev.LogHours(Friday, 6)
	dev.LogHours(Saturday, 8)

	hours, pay, isOvertime := dev.PayDay()
	// Выводим часы по дням
	dev.PrintWorkHourDetails()

	fmt.Printf("Total: %v worked: %v hours, isOvertime? %v, pay: $%v\n", dev.getFullName(), hours, isOvertime, pay)
}
