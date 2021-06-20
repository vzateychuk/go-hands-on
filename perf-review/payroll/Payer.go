package payroll

import "fmt"

type Payer interface {
	Pay() (fullname string, pay float64)
}

func PrintPayerDetails(payer Payer) {
	fullname, pay := payer.Pay()
	fmt.Printf("Payer: %v, pay: %0.2f\n", fullname, pay)
}
