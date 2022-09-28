package phone

import (
	"errors"
)

const (
	MaxLength = 11
	MinLength = 11
)

var (
	ErrNotValidPhone = errors.New("provided phone is not valid")
)

type Phone string

func (n Phone) String() string {
	return string(n)
}

func NewPhone(phone string) (*Phone, error) {
	phone = getNumbers(phone)
	if (len([]rune(phone))) < MinLength || (len([]rune(phone))) > MaxLength {
		return nil, ErrNotValidPhone
	}

	n := Phone(phone)

	return &n, nil
}

func getNumbers(input string) string {
	var number string

	for _, t := range input {
		if t >= 48 && t <= 57 { // 48 - 57 in ASCII this numbers   0-9
			number += string(t)
		}
	}

	return number
}
