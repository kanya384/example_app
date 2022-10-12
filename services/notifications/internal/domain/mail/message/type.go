package message

import (
	"fmt"
)

const (
	MinLength = 5
)

var (
	ErrWrongLength = fmt.Errorf("number of characters in the message must be at least %d", MinLength)
)

type Message string

func (n Message) String() string {
	return string(n)
}

func NewMessage(message string) (*Message, error) {
	if (len([]rune(message))) < MinLength {
		return nil, ErrWrongLength
	}

	n := Message(message)

	return &n, nil
}
