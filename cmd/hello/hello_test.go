package hello

import (
	"testing"
)

func TestHello(t *testing.T) {
	err := Hello()

	if err != nil {
		t.Error(err)
	}
}
