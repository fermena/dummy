package aggregate

import (
	"dummy/entity"
	"errors"

	"github.com/google/uuid"
)

var ErrMissingValues = errors.New("missing values")

type Product struct {
	item     *entity.Item
	price    float64
	quantity int
}

func deadFunction() {
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
}

func NewProduct(name, description string, price float64) (Product, error) {
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
