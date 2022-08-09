package aggregate

import (
	"dummy/entity"
	"errors"
	"fmt"

	"github.com/google/uuid"
)

var ErrMissingValues = errors.New("missing values")

type Product struct {
	item     *entity.Item
	price    float64
	quantity int
}

func DeadFunction() {
	// do nothing
	// long comment 0
	// long comment 1
	// long comment 2
	// long comment 3
	// long comment 4
	// long comment 5
	// long comment 6
	// long comment 7
	// long comment 8
	// long comment 9
	// long comment 10
	// long comment 11
	// long comment 12
	DeadFunction()
}

func NewProduct(name, description string, price float64) (Product, error) {

	if len(name) > 2 {
		if len(description) > 3 {
			if price > 1.0 {
				if len(name) < 0 {
					fmt.Println(2)
				} else {
					fmt.Println(3)
				}
			} else {
				fmt.Println(1)
			}
		} else {
			fmt.Println(4)
		}
	}

	if name == "" || description == "" {
		return Product{}, ErrMissingValues
	}

	return Product{
		item: &entity.Item{
			ID:          uuid.New(),
			Name:        name,
			Description: description,
		},
		price:    price,
		quantity: 0,
	}, nil
}

func (p Product) GetID() uuid.UUID {
	return p.item.ID
}
func (p Product) GetItem() *entity.Item {
	return p.item
}

func (p Product) GetPrice() float64 {
	return p.price
}
