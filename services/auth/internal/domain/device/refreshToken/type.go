package refreshToken

import (
	"fmt"

	"github.com/google/uuid"
)

const (
	MaxLength = 150
	MinLength = 5
)

var (
	ErrWrongLength = fmt.Errorf("number of characters in the refreshToken must be between %d and %d", MinLength, MaxLength)
)

type RefreshToken string

func (n RefreshToken) String() string {
	return string(n)
}

func New() (*RefreshToken, error) {
	n := RefreshToken(uuid.NewString())

	return &n, nil
}

func NewFromString(refreshToken string) (*RefreshToken, error) {
	if (len([]rune(refreshToken))) < MinLength || (len([]rune(refreshToken))) > MaxLength {
		return nil, ErrWrongLength
	}

	n := RefreshToken(refreshToken)

	return &n, nil
}
