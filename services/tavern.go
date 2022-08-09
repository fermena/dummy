package services

import (
	"log"

	"github.com/google/uuid"
)

type TavernConfiguration func(ts *Tavern) error

type Tavern struct {
	OrderService   *OrderService
	BillingService interface{}
}

func WithOrderService(os *OrderService) TavernConfiguration {
	return func(ts *Tavern) error {
		ts.OrderService = os
		return nil
	}
}

func NewTavern(cfgs ...TavernConfiguration) (*Tavern, error) {
	t := &Tavern{}

	for _, cfg := range cfgs {
		err := cfg(t)
		if err != nil {
			return nil, err
		}
	}
	return t, nil
}

func (t *Tavern) Order(customer uuid.UUID, products []uuid.UUID) error {
	price, err := t.OrderService.CreateOrder(customer, products)
	if err != nil {
		return err
	}
	log.Printf("Bill the customer: %0.0f", price)

	return nil
}
