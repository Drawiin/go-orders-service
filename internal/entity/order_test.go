package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ShouldReturnErr_WhenCreateANewOrder_WithEmptyID(t *testing.T) {
	_, err := NewOrder("", 10.0, 1.0)
	assert.EqualError(t, err, ErrInvalidId.Error())
}

func Test_ShouldReturnErr_WhenCreateANewOrder_WithEmptyZeroPrice(t *testing.T) {
	_, err := NewOrder("id", 0, 1.0)
	assert.EqualError(t, err, ErrInvalidPrice.Error())
}

func Test_ShouldReturnErr_WhenCreateANewOrder_WithZeroTax(t *testing.T) {
	_, err := NewOrder("id", 10.0, 0)
	assert.EqualError(t, err, ErrInvalidTax.Error())
}
