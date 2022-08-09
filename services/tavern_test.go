package services_test

import (
	"dummy/aggregate"
	"dummy/services"
	"testing"

	"github.com/google/uuid"
)

func TestTavern(t *testing.T) {
	product, err := aggregate.NewProduct("Prod1", "desc1", 1.99)
	if err != nil {
		t.Error(err)
	}

	products := []aggregate.Product{product}

	orderService, err := services.NewOrderService(
		services.WithMongoCustomerRepository("mongodb://localhost:27017"),
		//services.WithMemoryCustomerRepository(),
		services.WithMemoryProductRepository(products),
	)
	if err != nil {
		t.Error(err)
	}

	tavern, err := services.NewTavern(
		services.WithOrderService(orderService),
	)
	if err != nil {
		t.Error(err)
	}

	customer, err := aggregate.NewCustomer("Filete")
	if err != nil {
		t.Error(err)
	}

	orderService.Customers.Add(customer)

	err = tavern.Order(
		customer.GetID(),
		[]uuid.UUID{product.GetID()})
	if err != nil {
		t.Error(err)
	}

}
