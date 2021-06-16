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

func (dev *Developer) HoursWorked() (int, int) {
	totalHours := 0
	for idx, hourPerDay := range dev.workHours {
		totalHours += hourPerDay
		if hourPerDay > 0 {
			fmt.Printf("Hours worked on %v: %v\n", Weekday(idx), hourPerDay)
		}
	}
	return totalHours, dev.hourlyRate * totalHours
}

func main() {
	vlad := employee{0, "Vlad", "Zateychuk"}
	dev := Developer{vlad, 20, [7]int{}}
	// заполняем табель рабочего времени
	dev.LogHours(Weekday(Monday), 5).LogHours(Weekday(Tuesday), 7).LogHours(Weekday(Wednesday), 8)
	totalHours, salary := dev.HoursWorked()
	fmt.Printf("Total: %v worked: %v hours, salary: $%v\n", dev.getFullName(), totalHours, salary)
}
