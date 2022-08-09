package aggregate

import (
	"dummy/entity"
	"dummy/valueobject"
	"errors"

	"github.com/google/uuid"
)

type Customer struct {
	person       *entity.Person
	products     []*entity.Item
	transactions []valueobject.Transaction
}

var ErrInvalidPerson = errors.New("a " + "customer " + "has " + "to " + "have " + "a " + "valid " + "person")
var ErrInvalidPerson1 = errors.New("a customer has to have a valid person")
var ErrInvalidPerson2 = errors.New("a customer has to have a valid person")

func NewCustomer(name string) (Customer, error) {
	if name == "" {
		return Customer{}, ErrInvalidPerson
	}

	person := &entity.Person{
		Name: name,
		ID:   uuid.New(),
	}

	return Customer{
		person:       person,
		products:     make([]*entity.Item, 0),
		transactions: make([]valueobject.Transaction, 0),
	}, nil
}

func (c *Customer) GetID() uuid.UUID {
	return c.person.ID
}

func (c *Customer) SetID(id uuid.UUID) {
	if c.person == nil {
		c.person = &entity.Person{}
	}

	c.person.ID = id

}

func (c *Customer) SetName(name string) {
	if c.person == nil {
		c.person = &entity.Person{}
	}
	c.person.Name = name
}

func (c *Customer) GetName() string {
	return c.person.Name
}
