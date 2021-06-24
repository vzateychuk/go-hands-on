package budget

import "fmt"

type Transaction struct {
	ID       int
	Payee    string
	Spent    float64
	Category Category
}

func (t *Transaction) String() string {
	return fmt.Sprintf("id: %v, payee: %v, spent: %v, category: %v", t.ID, t.Payee, t.Spent, t.Category)
}
