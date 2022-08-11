package aggregate

import (
	"dummy/entity"
	"dummy/valueobject"
	"errors"
	"fmt"

	"github.com/google/uuid"
)

type Customer struct {
	person       *entity.Person
	products     []*entity.Item
	transactions []valueobject.Transaction
}

// TODO remove duplicates
var ErrInvalidPerson = errors.New("a customer has to have a valid person")
var ErrInvalidPerson1 = errors.New("a customer has to have a valid person")
var ErrInvalidPerson2 = errors.New("a customer has to have a valid person")
var ErrInvalidPerson3 = errors.New("a customer has to have a valid person")
var ErrInvalidPerson4 = errors.New("a customer has to have a valid person")

func doNothing() { // Noncompliant
}

func fun1() (x, y int) {
	a, b := 1, 2
	b, a = a, b
	return a, b
}

func fun2() (x, y int) { // Noncompliant; fun1 and fun2 have identical implementations
	a, b := 1, 2
	b, a = a, b
	return a, b
}

func compute(a int, b int) {
	sum := a + b
	if sum > 0 {
	} // Noncompliant; empty on purpose or missing piece of code?
	fmt.Println("Result:", sum)
}

/* */

/*

 */

func NewCustomer(name string) (Customer, error) {
	// FIXME remove code smells
	compute(fun1())
	compute(fun2())
	doNothing()

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
