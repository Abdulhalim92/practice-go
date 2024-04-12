package main

import (
	"github.com/Rhymond/go-money"
	"log"
	"practice_go/methods/cart"
	"practice_go/methods/product"
	"time"
)

func main() {
	c := cart.Cart{
		ID:        "115552221",
		CreatedAt: time.Now(),
	}

	cartPtr := &c
	cartPtr.Items = []cart.Item{
		{
			Product: product.Product{
				ID:    "1234",
				Name:  "First",
				Price: money.New(1000, money.EUR),
			},
			Quantity: 2,
		},
		{
			Product: product.Product{
				ID:    "1235",
				Name:  "Second",
				Price: money.New(2000, money.EUR),
			},
			Quantity: 3,
		},
	}
	log.Println(c.Items)

	(*cartPtr).Items = []cart.Item{
		{
			Product: product.Product{
				ID:    "1236",
				Name:  "Third",
				Price: money.New(1000, money.EUR),
			},
			Quantity: 2,
		},
		{
			Product: product.Product{
				ID:    "1237",
				Name:  "Fourth",
				Price: money.New(2000, money.EUR),
			},
			Quantity: 3,
		},
	}
	log.Println("after:")
	log.Println(c.Items)
}
