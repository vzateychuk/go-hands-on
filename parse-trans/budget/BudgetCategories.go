package budget

import (
	"errors"
	"strings"
)

var CategoryNotFoundError = errors.New("budget category is not found.")

type Category string

const (
	Fuel       Category = "fuel"
	Food       Category = "food"
	Mortgage   Category = "mortgage"
	Repairs    Category = "repairs"
	Insurance  Category = "insurance"
	Utilities  Category = "utilities"
	Retirement Category = "retirement"
)

func GetCategoryBy(name string) (Category, error) {
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
		return "", CategoryNotFoundError
	}
}
