package cart

import (
	"fmt"
	"practice_go/methods/product"
	"practice_go/methods/user"
	"time"

	"github.com/Rhymond/go-money"
)

type Item struct {
	product.Product
	Quantity uint8
}

type Cart struct {
	ID        string
	CreatedAt time.Time
	UpdatedAt time.Time
	lockedAt  time.Time
	user.User
	Items        []Item
	CurrencyCode string
	isLocked     bool
}

func (c *Cart) TotalPrice() (*money.Money, error) {
	total := money.New(0, c.CurrencyCode)
	var err error

	for _, item := range c.Items {
		itemSubtotal := item.Product.Price.Multiply(int64(item.Quantity))
		total, err = total.Add(itemSubtotal)
		if err != nil {
			return nil, err
		}
	}

	return total, nil
}

func (c *Cart) Lock() error {
	if c.isLocked == true {
		return fmt.Errorf("the cart is already locked")
	}

	c.lockedAt = time.Now()
	c.isLocked = true

	return nil
}

func (c *Cart) delete() error {
	return nil
}
