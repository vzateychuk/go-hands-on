package main

import (
	"fmt"
)

type Item struct {
	Name    string
	Cost    float64
	TaxRate float64
}

var sales = []Item{
	{"Cake", 0.99, 7.5},
	{"Milk", 2.65, 1.5},
	{"Butter", 0.87, 2},
}

func calcTax(cost float64, rate float64) float64 {
	return cost * rate
}

func main() {
	var taxTotal float64 = 0.0
	for index, item := range sales {
		tax := calcTax(item.Cost, item.TaxRate)
		taxTotal += tax
		fmt.Printf("%d: %v, cost: %v, rate: %v, tax: %v\n", index, item.Name, item.Cost, item.TaxRate, tax)
	}
	fmt.Println("Total: ", taxTotal)
}
