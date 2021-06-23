package budget

import "fmt"

type Transaction struct {
	ID       string
	Payee    string
	Spent    string
	Category Category
}

func (t *Transaction) String() string {
	return fmt.Sprintf("id: %v, payee: %v, spent: %v, category: %v", t.ID, t.Payee, t.Spent, t.Category)
}
