package main

import (
	"errors"
	"fmt"
	"os"
)

type Employee struct {
	Id        int
	FirstName string
	LastName  string
}
type Developer struct {
	Employee
	HourlyRate        int
	HoursWorkedInYear int
	Review            map[string]interface{}
}

type Payer interface {
	Pay() (fullname string, pay float64)
}

type Manager struct {
	Employee
	Salary         float64
	CommissionRate float64
}

func (dev *Developer) FullName() string {
	return dev.FirstName + " " + dev.LastName
}

func (man Manager) FullName() string {
	return man.FirstName + " " + man.LastName
}

func (dev Developer) Pay() (string, float64) {
	payment := dev.HourlyRate * dev.HoursWorkedInYear
	return dev.FullName(), float64(payment)
}

func (man Manager) Pay() (string, float64) {
	payment := man.Salary + (man.Salary * man.CommissionRate)
	return man.FullName(), payment
}

func printPayerDetails(payer Payer) {
	fullname, pay := payer.Pay()
	fmt.Printf("Payer: %v, pay: %0.2f\n", fullname, pay)
}

func convertReviewToInt(str string) (rate int, err error) {
	switch str {
	case "Excellent":
		return 5, nil
	case "Good":
		return 4, nil
	case "Fair":
		return 3, nil
	case "Poor":
		return 2, nil
	case "Unsatisfactory":
		return 1, nil
	default:
		return 0, errors.New("invalid rating: " + str)
	}
}

func OverallReview(review interface{}) (int, error) {
	switch val := review.(type) {
	case int:
		return val, nil
	case string:
		rating, err := convertReviewToInt(val)
		if err != nil {
			return 0, err
		}
		return rating, nil
	default:
		return 0, errors.New("unknown type")
	}
}

func (dev Developer) ReviewRating() error {
	sum := 0
	for _, review := range dev.Review {
		val, err := OverallReview(review)
		if err != nil {
			return err
		}
		sum += val
	}
	averageRating := float64(sum) / float64(len(dev.Review))
	fmt.Printf("%s got a review rating of %.2f\n", dev.FullName(), averageRating)
	return nil
}

var (
	employeeReview map[string]interface{} = map[string]interface{}{
		"WorkQuality":     5,
		"TeamWork":        2,
		"Communication":   "Poor",
		"Problem-solving": 4,
		"Dependability":   "Unsatisfactory",
	}
	vzateychuk = Employee{
		Id:        10,
		FirstName: "Vladimir",
		LastName:  "Zateychuk",
	}
	dev       = Developer{vzateychuk, 35, 2400, employeeReview}
	ivanovich = Employee{
		Id:        20,
		FirstName: "Ivan Ivanovich",
		LastName:  "Ivanov",
	}
	manager = Manager{
		Employee:       ivanovich,
		Salary:         150000.00,
		CommissionRate: 0.7,
	}
)

func main() {
	err := dev.ReviewRating()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	printPayerDetails(dev)
	printPayerDetails(manager)
}
