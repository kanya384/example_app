package agent

import (
	"fmt"
)

const (
	MaxLength = 150
	MinLength = 5
)

var (
	ErrWrongLength = fmt.Errorf("number of characters in the agent must be between %d and %d", MinLength, MaxLength)
)

type Agent string

func (n Agent) String() string {
	return string(n)
}

func NewAgent(agent string) (*Agent, error) {
	if (len([]rune(agent))) < MinLength || (len([]rune(agent))) > MaxLength {
		return nil, ErrWrongLength
	}

	n := Agent(agent)

	return &n, nil
}
