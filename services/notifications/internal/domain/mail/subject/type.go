package subject

import (
	"fmt"
)

const (
	MaxLength = 78
	MinLength = 5
)

var (
	ErrWrongLength = fmt.Errorf("number of characters in the subject must be between %d and %d", MinLength, MaxLength)
)

type Subject string

func (n Subject) String() string {
	return string(n)
}

func NewSubject(subject string) (*Subject, error) {
	if (len([]rune(subject))) < MinLength || (len([]rune(subject))) > MaxLength {
		return nil, ErrWrongLength
	}

	n := Subject(subject)

	return &n, nil
}
