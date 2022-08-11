package services

import (
	"context"
	"dummy/aggregate"
	"dummy/domain/customer"
	"dummy/domain/customer/memory"
	"dummy/domain/customer/mongo"
	"dummy/domain/product"
	"log"

	prodmemory "dummy/domain/product/memory"

	"github.com/google/uuid"
)

type OrderService struct {
	Customers customer.CustomerRepository
	Products  product.ProductRepository
}

type OrderConfiguration func(os *OrderService) error

func WithMongoCustomerRepository(connectionString string) OrderConfiguration {
	return func(os *OrderService) error {
		cr, err := mongo.New(context.Background(), connectionString)
		if err != nil {
			return err
		}
		os.Customers = cr
		return nil
	}
}
func WithMemoryProductRepository(products []aggregate.Product) OrderConfiguration {
	return func(os *OrderService) error {
		pr := prodmemory.New()

		for _, p := range products {
			err := pr.Add(p)
			if err != nil {
				return err
			}
		}

		os.Products = pr
		return nil
	}
}

func WithCustomerRepository(cr customer.CustomerRepository) OrderConfiguration {
	return func(os *OrderService) error {
		os.Customers = cr
		return nil
	}
}

func WithMemoryCustomerRepository() OrderConfiguration {
	cr := memory.New()
	return WithCustomerRepository(cr)
}

func NewOrderService(cfgs ...OrderConfiguration) (*OrderService, error) {
	os := &OrderService{}

	for _, cfg := range cfgs {
		err := cfg(os)
		if err != nil {
			return nil, err
		}
	}

	return os, nil
}

func (o *OrderService) CreateOrder(customerID uuid.UUID, productIDs []uuid.UUID) (float64, error) {
	c, err := o.Customers.Get(customerID)
	if err != nil {
		return 0, err
	}

	var products []aggregate.Product
	var price float64

	for _, id := range productIDs {
		p, err := o.Products.GetByID(id)
		if err != nil {
			return 0, err
		}
		products = append(products, p)
		price += p.GetPrice()
	}

	log.Printf("Customer %s has ordered %d products", c.GetID(), len(products))
	return price, nil
}
