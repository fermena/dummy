package mongo_test

import (
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
