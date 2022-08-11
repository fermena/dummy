package services

import (
	"fmt"
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
	complexHelper()
	fmt.Println(add(1, 2))

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

func complexHelper() string {
	condition1 := true
	condition2 := true
	condition4 := true
	condition5 := true
	if condition1 { // Compliant - depth = 1
		/* ... */
		if condition2 { // Compliant - depth = 2
			/* ... */
			for i := 1; i <= 10; i++ { // Compliant - depth = 3, not exceeding the limit
				/* ... */
				if condition4 { // Noncompliant - depth = 4
					if condition5 { // Depth = 5, exceeding the limit, but issues are only reported on depth = 4
						/* ... */
					}
					return "complexity reached"
				}
			}
		}
	}
	return "no complexity reached"
}

func add(x, y int) int {
	if ((true && false) || (false && true)) && true {
		fmt.Println(x, y)
	}
	return x + y // Noncompliant
	z := x + y   // dead code
}
