package pass

import (
	"errors"
	"fmt"
	"unicode"
)

const (
	MinLength = 8
)

var (
	ErrPassLength         = fmt.Errorf("provided pass is too short, min length: %d", MinLength)
	ErrNoUpperCaseLetter  = errors.New("pass must contain at least one upper case letter")
	ErrNoLowerCaseLetter  = errors.New("pass must contain at least one lower case letter")
	ErrNoNumberInPass     = errors.New("pass must contain at least one number")
	ErrNoSpecialCharacter = errors.New("pass must contain at least one special character")
)

type Pass string

func (n Pass) String() string {
	return string(n)
}

func NewPass(pass string) (*Pass, error) {
	if err := vaildatePassword(pass); err != nil {
		return nil, err
	}

	n := Pass(pass)

	return &n, nil
}

func vaildatePassword(field string) (err error) {
	var (
		upper   bool
		lower   bool
		number  bool
		special bool
	)

	if len(field) < MinLength {
		wrapError(err, ErrPassLength)
	}

	for _, c := range field {
		switch {
		case unicode.IsUpper(c):
			upper = true
		case unicode.IsLower(c):
			lower = true
		case unicode.IsNumber(c):
			number = true
		case unicode.IsPunct(c) || unicode.IsSymbol(c):
			special = true
		}
	}

	if !special {
		wrapError(err, ErrNoSpecialCharacter)
	}
	if !number {
		wrapError(err, ErrNoNumberInPass)
	}
	if !lower {
		wrapError(err, ErrNoLowerCaseLetter)
	}
	if !upper {
		wrapError(err, ErrNoUpperCaseLetter)
	}

	return
}

func wrapError(err, appendError error) error {
	if err != nil {
		return fmt.Errorf("%s:%w", appendError.Error(), err)
	}
	return appendError
}
