package user

import (
	"errors"
	"fmt"
)

var ErrFistNameRequeried = errors.New("first name is required")
var ErrLastNameRequeried = errors.New("last name is required")
var ErrEmailRequeried = errors.New("email is required")
var ErrThereArentFields = errors.New("there aren´t fields")

type ErrNotFound struct {
	ID uint64
}

func (e ErrNotFound) Error() string {
	return fmt.Sprintf("user id '%d' not found", e.ID)
}
