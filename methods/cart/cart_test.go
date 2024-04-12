package cart

import (
	"github.com/Rhymond/go-money"
	"github.com/stretchr/testify/assert"
	"practice_go/methods/product"
	"practice_go/methods/user"
	"testing"
	"time"
)

func TestTotalPrice(t *testing.T) {
	items := []Item{
		{
			Product: product.Product{
				ID:    "p-1254",
				Name:  "Product test",
				Price: money.New(1000, money.EUR),
			},
			Quantity: 2,
		},
		{
			Product: product.Product{
				ID:    "p-1255",
				Name:  "Product test 2",
				Price: money.New(2000, money.EUR),
			},
			Quantity: 1,
		},
	}

	c := Cart{
		ID:           "1254",
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
		User:         user.User{},
		Items:        items,
		CurrencyCode: "EUR",
	}
	actual, err := c.TotalPrice()
	assert.NoError(t, err)
	assert.Equal(t, money.New(4000, money.EUR), actual)
}

func TestLock(t *testing.T) {
	c := Cart{
		ID: "1254",
	}

	err := c.Lock()
	assert.NoError(t, err)
	assert.True(t, c.isLocked)
	assert.True(t, c.lockedAt.Unix() > 0)
}

func TestLockAlreadyLocked(t *testing.T) {
	c := Cart{
		ID:       "1254",
		isLocked: true,
	}

	err := c.Lock()
	assert.Error(t, err)
}
