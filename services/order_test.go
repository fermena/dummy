package services

import (
	"dummy/aggregate"
	"testing"

	"github.com/google/uuid"
)

func initProducts(t *testing.T) []aggregate.Product {
	beer, err := aggregate.NewProduct("beer", "Healthy beverage", 1.99)
	if err != nil {
		t.Error(err)
	}

	wine, err := aggregate.NewProduct("wine", "Healthy beverage too", 2.00)
	if err != nil {
		t.Error(err)
	}

	peanuts, err := aggregate.NewProduct("peanuts", "Food for squirrels", 0.54)
	if err != nil {
		t.Error(err)
	}

	products := []aggregate.Product{beer, wine, peanuts}
	return products
}

func TestOrderNewOrderService(t *testing.T) {
	products := initProducts(t)

	os, err := NewOrderService(
		WithMemoryCustomerRepository(),
		WithMemoryProductRepository(products),
	)
	if err != nil {
		t.Error(err)
	}

	cust, err := aggregate.NewCustomer("Filete")
	if err != nil {
		t.Error(err)
	}

	err = os.Customers.Add(cust)
	if err != nil {
		t.Error(err)
	}

	order := []uuid.UUID{
		products[0].GetID(),
	}

	_, err = os.CreateOrder(cust.GetID(), order)
	if err != nil {
		t.Error(err)
	}
}
