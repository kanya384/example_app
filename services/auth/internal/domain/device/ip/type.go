package ip

import (
	"errors"
	"fmt"
	"net"
)

const (
	MaxLength = 15
	MinLength = 7
)

var (
	ErrWrongLength = fmt.Errorf("number of characters in the ip must be between %d and %d", MinLength, MaxLength)
	ErrNotValidIP  = errors.New("invalid ip address")
)

type Ip string

func (n Ip) String() string {
	return string(n)
}

func NewIp(ip string) (*Ip, error) {
	if (len([]rune(ip))) < MinLength || (len([]rune(ip))) > MaxLength {
		return nil, ErrWrongLength
	}

	if !checkIPAddress(ip) {
		return nil, ErrNotValidIP
	}

	n := Ip(ip)

	return &n, nil
}

func checkIPAddress(ip string) bool {
	switch net.ParseIP(ip) {
	case nil:
		return false
	default:
		return true
	}
}
