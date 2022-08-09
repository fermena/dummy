package memory

import (
	"dummy/aggregate"
	"dummy/domain/customer"
	"testing"

	"github.com/google/uuid"
)

func TestMemoryGetCustomer(t *testing.T) {
	type testCase struct {
		name        string
		id          uuid.UUID
		expectedErr error
	}

	cust, err := aggregate.NewCustomer("Filete")
	if err != nil {
		t.Fatal(err)
	}

	id := cust.GetID()

	repo := MemoryRepository{
		customers: map[uuid.UUID]aggregate.Customer{
			id: cust,
		},
	}

	testCases := []testCase{
		{
			name:        "No Customer By ID",
			id:          uuid.MustParse("f47ac10b-58cc-0372-8567-0e02b2c3d479"),
			expectedErr: customer.ErrCustomerNotFound,
		},
		{
			name:        "Customer By ID",
			id:          id,
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := repo.Get(tc.id)
			if err != tc.expectedErr {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}
		})
	}

}

func TestMemoryAddCustomer(t *testing.T) {
	type testCase struct {
		name        string
		cust        string
		expectedErr error
	}

	testCases := []testCase{
		{
			name:        "Add Customer",
			cust:        "Filete",
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			repo := MemoryRepository{
				customers: map[uuid.UUID]aggregate.Customer{},
			}

			cust, err := aggregate.NewCustomer(tc.cust)

			if err != nil {
				t.Fatal(err)
			}

			err = repo.Add(cust)
			if err != tc.expectedErr {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}

			found, err := repo.Get(cust.GetID())
			if err != nil {
				t.Fatal(err)
			}

			if found.GetID() != cust.GetID() {
				t.Errorf("Expected %v, got %v", cust.GetID(), found.GetID())
			}
		})
	}
}

func TestMemoryUpdate(t *testing.T) {
	type testCase struct {
		name         string
		originalName string
		updatedName  string
		customers    map[uuid.UUID]aggregate.Customer
		expectedErr  error
	}

	cust, err := aggregate.NewCustomer("Filete")

	testCases := []testCase{
		{
			name:         "Update customer, when customer exists",
			originalName: "Filete",
			updatedName:  "Filetito",
			customers: map[uuid.UUID]aggregate.Customer{
				cust.GetID(): cust,
			},

			expectedErr: nil,
		},
		{
			name:         "Update customer, when customer doesn't exist",
			originalName: "Filete",
			updatedName:  "Filetito",
			customers:    map[uuid.UUID]aggregate.Customer{},
			expectedErr:  customer.ErrUpdateCustomer,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			repo := MemoryRepository{
				customers: tc.customers,
			}
			err = repo.Update(cust)

			if err != tc.expectedErr {
				t.Errorf("Expected error: '%v' got: '%v'", tc.expectedErr, err)
			}

		})
	}
}

func TestMemoryNew(t *testing.T) {
	repo := New()

	got := len(repo.customers)
	expect := 0

	if got != expect {
		t.Errorf("Got %d expect %d", got, expect)
	}
}
