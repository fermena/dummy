package aggregate_test

import (
	"dummy/aggregate"
	"fmt"
	"testing"

	"github.com/google/uuid"
)

func TestCustomerNewCustomer(t *testing.T) {
	type testCase struct {
		test        string
		name        string
		expectedErr error
	}

	testCases := []testCase{
		{
			test:        "Empty name validation",
			name:        "",
			expectedErr: aggregate.ErrInvalidPerson,
		},
		{
			test:        "Valid name",
			name:        "Filete",
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			_, err := aggregate.NewCustomer(tc.name)
			if err != tc.expectedErr {
				t.Errorf("Expected error %v got %v", tc.expectedErr, err)
			}
		})
	}
}

func TestCustomerGetID(t *testing.T) {
	c, err := aggregate.NewCustomer("Filete")
	if err != nil {
		t.Error(err)
	}

	got := len(c.GetID())
	expect := 16

	if got != expect {
		t.Errorf("Got %d expect %d", got, expect)
	}
	fmt.Printf("c.GetID(): %v\n", c.GetID())
}

func TestCustomerSetID(t *testing.T) {
	c, err := aggregate.NewCustomer("Filete")
	if err != nil {
		t.Error(err)
	}

	id := uuid.New()

	c.SetID(id)
	got := c.GetID()

	if got != id {
		t.Errorf("Got %v expect %v", got, id)
	}
}

func TestCustomerSetName(t *testing.T) {
	c, err := aggregate.NewCustomer("Filete")
	if err != nil {
		t.Error(err)
	}

	c.SetName("Filetito")
	got := c.GetName()
	expect := "Filetito"

	if got != expect {
		t.Errorf("Got %s expect %s", got, expect)
	}

}

func TestCustomerGetName(t *testing.T) {
	c, err := aggregate.NewCustomer("Filete")
	if err != nil {
		t.Error(err)
	}

	got := c.GetName()
	expect := "Filete"

	if got != expect {
		t.Errorf("Got %s expect %s", got, expect)
	}
}
