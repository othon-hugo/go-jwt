package jwt

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidTokenSupplied = errors.New("jwt: invalid token supplied")
	ErrSignatureMismatch    = errors.New("jwt: signature mismatch during verification")
)

type UnsupportedAlgorithmError struct {
	alg string
}

func (e UnsupportedAlgorithmError) Error() string {
	return fmt.Sprintf("jwt: unsupported algorithm: %s", e.alg)
}

type UnsupportedTypeError struct {
	typ string
}

func (e UnsupportedTypeError) Error() string {
	return fmt.Sprintf("jwt: unsupported type: %s", e.typ)
}
