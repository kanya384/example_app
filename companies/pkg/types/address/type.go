package address

import (
	"fmt"
)

const (
	MaxLength = 300
	MinLength = 10
)

var (
	ErrWrongLength = fmt.Errorf("number of characters in the adress must be between %d and %d", MinLength, MaxLength)
)

type Address string

func (n Address) String() string {
	return string(n)
}

func NewAdress(address string) (*Address, error) {
	if (len([]rune(address))) < MinLength || (len([]rune(address))) > MaxLength {
		return nil, ErrWrongLength
	}

	n := Address(address)

	return &n, nil
}
