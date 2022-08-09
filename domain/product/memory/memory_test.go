package prodmemory

import (
	"dummy/aggregate"
	"dummy/domain/product"
	"testing"

	"github.com/google/uuid"
)

func TestMemoryProductRepositoryAdd(t *testing.T) {
	repo := New()

	product, err := aggregate.NewProduct("Beer", "Good for you're health", 1.99)
	if err != nil {
		t.Error(err)
	}

	repo.Add(product)
	if len(repo.products) != 1 {
		t.Errorf("Expected 1 product, got %d", len(repo.products))
	}
}

func TestMemoryProductRepositoryGet(t *testing.T) {
	repo := New()
	existingProd, err := aggregate.NewProduct("Beer", "Good for you're health", 1.99)
	if err != nil {
		t.Error(err)
	}

	repo.Add(existingProd)
	if len(repo.products) != 1 {
		t.Errorf("Expected 1 product, got %d", len(repo.products))
	}

	type testCase struct {
		name        string
		id          uuid.UUID
		expectedErr error
	}

	testCases := []testCase{
		{
			name:        "Get product by id",
			id:          existingProd.GetID(),
			expectedErr: nil,
		},
		{
			name:        "Get non-existing product by id",
			id:          uuid.New(),
			expectedErr: product.ErrProductNotFound,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := repo.GetByID(tc.id)
			if err != tc.expectedErr {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}
		})
	}
}

func TestMemoryProductRepositoryDelete(t *testing.T) {
	repo := New()

	existingProd, err := aggregate.NewProduct("beer", "good for you're health", 1.99)

	if err != nil {
		t.Error(err)
	}

	repo.Add(existingProd)

	if len(repo.products) != 1 {
		t.Errorf("Expected 1 product, got %v", len(repo.products))
	}

	err = repo.Delete(existingProd.GetID())
	if err != nil {
		t.Error(err)
	}

	if len(repo.products) != 0 {
		t.Errorf("Expected 0 products, got %v", len(repo.products))
	}
}
