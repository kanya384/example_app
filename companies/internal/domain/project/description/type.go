package description

import (
	"fmt"
)

const (
	MaxLength = 400
	MinLength = 10
)

var (
	ErrWrongLength = fmt.Errorf("number of characters in the description must be between %d and %d", MinLength, MaxLength)
)

type Description string

func (n Description) String() string {
	return string(n)
}

func NewDescription(description string) (*Description, error) {
	if (len([]rune(description))) < MinLength || (len([]rune(description))) > MaxLength {
		return nil, ErrWrongLength
	}

	n := Description(description)

	return &n, nil
}
