package main

import (
	"errors"
	"fmt"
	"strings"
)

var (
	ErrInvalidRoutingNum = errors.New("invalid routing num")
	ErrInvalidLastName   = errors.New("invalid last name")
)

type directDeposit struct {
	lastName      string
	firstName     string
	bankName      string
	routingNumber int
	accountNumber int
}

// The method  will return ErrInvalidRoutingNum when the routing number is less than 100
func (depo *directDeposit) validateRoutingNumber() error {

	if depo.routingNumber < 100 {
		return ErrInvalidRoutingNum
	} else {
		return nil
	}
}

/* return ErrInvalidLastName when the lastName is an empty string. */
func (depo *directDeposit) validateLastName() error {
	strings.TrimSpace(depo.lastName)
	if len(strings.TrimSpace(depo.lastName)) == 0 {
		return ErrInvalidLastName
	} else {
		return nil
	}
}

func (depo *directDeposit) report() {
	fmt.Println("lastName: ", depo.lastName)
	fmt.Println("firstName: ", depo.firstName)
	fmt.Println("bankName: ", depo.bankName)
	fmt.Println("routingNumber: ", depo.routingNumber)
	fmt.Println("accountNumber: ", depo.accountNumber)
}

func main() {
	depo := directDeposit{
		lastName:      "  ",
		firstName:     "Vladimir",
		bankName:      "ING",
		routingNumber: 99,
		accountNumber: 456,
	}
	err := depo.validateRoutingNumber()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	err = depo.validateLastName()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	depo.report()
}
