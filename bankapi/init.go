package main

import "bankcore"

var accounts = map[float64]*CustomAccount{}

func init() {
	accounts[1001] = &CustomAccount{
		Account: &bankcore.Account{
			Customer: bankcore.Customer{
				Name:    "John",
				Address: "Los Angeles, California",
				Phone:   "(213) 555 0147",
			},
			Number: 1001,
		},
	}

	accounts[1002] = &CustomAccount{
		Account: &bankcore.Account{
			Customer: bankcore.Customer{
				Name:    "Mark",
				Address: "Irvine, California",
				Phone:   "(949) 555 0198",
			},
			Number: 1002,
		},
	}
}
