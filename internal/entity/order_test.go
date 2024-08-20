package entity

import (
	"testing"

	"github.com/drawiin/go-orders-service/internal/failure"
	"github.com/stretchr/testify/assert"
)

func Test_ShouldReturnErr_WhenCreateANewOrder_WithEmptyID(t *testing.T) {
	_, err := NewOrder("", 10.0, 1.0)
	assert.EqualError(t, err, failure.ErrInvalidId.Error())
}

func Test_ShouldReturnErr_WhenCreateANewOrder_WithEmptyZeroPrice(t *testing.T) {
	_, err := NewOrder("id", 0, 1.0)
	assert.EqualError(t, err, failure.ErrInvalidPrice.Error())
}

func Test_ShouldReturnErr_WhenCreateANewOrder_WithZeroTax(t *testing.T) {
	_, err := NewOrder("id", 10.0, 0)
	assert.EqualError(t, err, failure.ErrInvalidTax.Error())
}
