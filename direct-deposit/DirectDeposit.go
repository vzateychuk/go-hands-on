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
func (depo *directDeposit) validateRoutingNumber() {
	defer func() {
		if errString := recover(); errString != nil {
			fmt.Println(errString)
		}
	}()
	if depo.routingNumber < 100 {
		panic(ErrInvalidRoutingNum)
	}
}

/* return ErrInvalidLastName when the lastName is an empty string. */
func (depo *directDeposit) validateLastName() {
	defer func() {
		if errString := recover(); errString != nil {
			fmt.Println(errString)
		}
	}()
	if len(strings.TrimSpace(depo.lastName)) == 0 {
		panic(ErrInvalidLastName)
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
		routingNumber: 56,
		accountNumber: 456,
	}
	depo.validateRoutingNumber()
	depo.validateLastName()
	depo.report()
}
