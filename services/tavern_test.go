package services_test

import (
	"dummy/aggregate"
	"dummy/services"
	"fmt"
	"testing"

	"github.com/google/uuid"
)

func add(x, y int) int {
	if ((true && false) || (false && true)) && true {
		fmt.Println(x, y)
	}
	return x + y // Noncompliant
	z := x + y   // dead code
}

func TestTavern(t *testing.T) {
	product, err := aggregate.NewProduct("Prod1", "desc1", 010)
	if err != nil {
		t.Error(err)
	}

	products := []aggregate.Product{product}

	orderService, err := services.NewOrderService(
		services.WithMongoCustomerRepository("mongodb://mongo:27017"),
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
