package pass

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"unicode"
)

const (
	MinLength = 8
)

var (
	ErrPassLength         = fmt.Errorf("is too short, min length: %d", MinLength)
	ErrNoUpperCaseLetter  = errors.New("must contain at least one upper case letter")
	ErrNoLowerCaseLetter  = errors.New("must contain at least one lower case letter")
	ErrNoNumber           = errors.New("must contain at least one number")
	ErrNoSpecialCharacter = errors.New("must contain at least one special character")
)

type Pass string

func (n Pass) String() string {
	return string(n)
}

func NewPass(pass, salt string) (*Pass, error) {
	if err := vaildatePassword(pass); err != nil {
		return nil, err
	}

	n := Pass(generateHashPassword(pass, salt))

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
		err = wrapError(err, ErrPassLength)
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
		err = wrapError(err, ErrNoSpecialCharacter)
	}
	if !number {
		err = wrapError(err, ErrNoNumber)
	}
	if !lower {
		err = wrapError(err, ErrNoLowerCaseLetter)
	}
	if !upper {
		err = wrapError(err, ErrNoUpperCaseLetter)
	}

	return
}

func wrapError(err, appendError error) error {
	if err != nil {
		return fmt.Errorf("%s:%s", appendError.Error(), err.Error())
	}
	return appendError
}

func generateHashPassword(password string, salt string) string {
	hash := sha256.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
