package budget

import (
	"errors"
	"strings"
)

var CategoryNotFoundError = errors.New("budget category is not found.")

type Category int

const (
	Fuel = iota
	Food
	Mortgage
	Repairs
	Insurance
	Utilities
	Retirement
	Unknown
)

func GetByName(name string) (Category, error) {
	lower := strings.ToLower(strings.TrimSpace(name))
	switch lower {
	case "fuel", "gas":
		return Fuel, nil
	case "food":
		return Food, nil
	case "mortgage":
		return Mortgage, nil
	case "repairs":
		return Repairs, nil
	case "car insurance", "life insurance":
		return Insurance, nil
	case "utilities":
		return Utilities, nil
	case "retirement":
		return Retirement, nil
	default:
		return Unknown, CategoryNotFoundError
	}
}

func (c Category) String() string {
	cats := []string{"Fuel", "Food", "Mortgage", "Repairs", "Insurance", "Utilities", "Retirement", "Unknown"}
	return cats[c]
}
