package failure

import "errors"

var (
	ErrInvalidId    = errors.New("invalid id")
	ErrInvalidPrice = errors.New("invalid price")
	ErrInvalidTax   = errors.New("invalid tax")
)
