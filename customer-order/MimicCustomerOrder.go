/*
In this activity, we are going to mimic a customer Order. An online e-commerce portal
needs to accept customer orders over its web application. As the customer browses
through the site, the customer will add items to their Order. This web application will
need to be able to take the JSON and add orders to the JSON.
*/
package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Customer struct {
	UserName string  `json:"username"`
	ShipTo   Address `json:"shipto"`
	Order    Order   `json:"order"`
}

type Address struct {
	Street  string `json:"street"`
	City    string `json:"city"`
	State   string `json:"state"`
	Zipcode int    `json:"zipcode"`
}

type Item struct {
	Name        string `json:"itemname"`
	Description string `json:"desc,omitempty"`
	Quantity    int    `json:"qty"`
	Price       int    `json:"price"`
}

type Order struct {
	IsPaid      bool   `json:"paid"`
	Fragile     bool   `json:"fragile,omitempty"`
	OrderDetail []Item `json:"orderdetail"`
}

func (order *Order) TotalPrice() int {
	total := 0
	for _, item := range order.OrderDetail {
		total += item.Quantity * item.Price
	}
	return total
}

var jsonData = []byte(`
{
	"username" :"blackhat",
	"shipto":
		{
			"street": "Sulphur Springs Rd",
			"city": "Park City",
			"state": "VA",
			"zipcode": 12345
		},
	"order": {
			"paid":true,
			"orderdetail" : [
				{
					"itemname":"A Guide to the World of zeros and ones",
					"desc": "book",
					"qty": 3,
					"price": 50
				},
				{
					"itemname":"Blue and Green socks",
					"desc": "socks",
					"qty": 17,
					"price": 3
				},
				{
					"itemname":"Spaceship Morningstar",
					"desc": "spaceship",
					"qty": 1,
					"price": 158
				}
			]
		}
}
`)

func main() {
	customer := Customer{}
	err := json.Unmarshal(jsonData, &customer)
	if err != nil {
		log.Fatal(err)
	}
	amount := len(customer.Order.OrderDetail)
	total := customer.Order.TotalPrice()
	fmt.Printf("%v bought %v items, total spent: %d\n", customer.UserName, amount, total)
}
