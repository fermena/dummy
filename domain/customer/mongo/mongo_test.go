package mongo_test

import (
	"context"
	"dummy/aggregate"
	"dummy/domain/customer/mongo"
	"testing"
)

func TestMongoNewFromCustomer(t *testing.T) {
	c, err := aggregate.NewCustomer("Filete")
	if err != nil {
		t.Error(err)
	}
	mongoCustomer := mongo.NewFromCustomer(c)

	got := mongoCustomer.Name
	expect := c.GetName()

	if got != expect {
		t.Errorf("got %v expected %v", got, expect)
	}

}

func TestMongoToAggregate(t *testing.T) {
	mongoCustomer := mongo.MongoCustomer{}
	a := mongoCustomer.ToAggregate()

	got := a.GetID().String()
	expect := mongoCustomer.ID.String()

	if got != expect {
		t.Errorf("got %v expect %v", got, expect)
	}
}

func TestMongoNew(t *testing.T) {
	type testCase struct {
		name             string
		connectionString string
		expectedErr      error
		mrNil            bool
	}

	testCases := []testCase{
		{
			name:             "Test MongoRepository not nil",
			connectionString: "mongodb://mongo",
			expectedErr:      nil,
			mrNil:            false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()

			mr, err := mongo.New(ctx, tc.connectionString)
			if err != tc.expectedErr {
				t.Error(err)
			}

			if tc.mrNil && mr != nil {
				t.Errorf("Got %v expect nil", mr)
			}

			if !tc.mrNil && mr == nil {
				t.Errorf("Got %v expect MongoRepository", mr)

			}

		})
	}

}

func TestMongoAdd(t *testing.T) {
	customer, err := aggregate.NewCustomer("Filete")
	if err != nil {
		t.Error(err)
	}

	mr, err := mongo.New(context.Background(), "mongodb://mongo")
	if err != nil {
		t.Error(err)
	}

	err = mr.Add(customer)

	if err != nil {
		t.Error(err)
	}
}

func TestMongoGet(t *testing.T) {
	mr, err := mongo.New(context.Background(), "mongodb://mongo")
	if err != nil {
		t.Error(err)
	}
	customer, err := aggregate.NewCustomer("Filete")
	if err != nil {
		t.Error(err)
	}

	err = mr.Add(customer)
	if err != nil {
		t.Error(err)
	}

	outputCustomer, err := mr.Get(customer.GetID())
	if err != nil {
		t.Error(err)
	}

	got := outputCustomer.GetID()
	expect := customer.GetID()

	if got != expect {
		t.Errorf("Got %v expect %v", got, expect)
	}

}
