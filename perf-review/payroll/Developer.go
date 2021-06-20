package payroll

import (
	"errors"
	"fmt"
)

type Developer struct {
	Employee
	HourlyRate        int
	HoursWorkedInYear int
	Review            map[string]interface{}
}

func (dev *Developer) FullName() string {
	return dev.FirstName + " " + dev.LastName
}

func (dev Developer) Pay() (string, float64) {
	payment := dev.HourlyRate * dev.HoursWorkedInYear
	return dev.FullName(), float64(payment)
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
