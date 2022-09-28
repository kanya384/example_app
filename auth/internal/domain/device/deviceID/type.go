package deviceID

import (
	"fmt"
)

const (
	MaxLength = 100
	MinLength = 10
)

var (
	ErrWrongLength = fmt.Errorf("number of characters in the deviceID must be between %d and %d", MinLength, MaxLength)
)

type DeviceID string

func (n DeviceID) String() string {
	return string(n)
}

func NewDeviceID(deviceID string) (*DeviceID, error) {
	if (len([]rune(deviceID))) < MinLength || (len([]rune(deviceID))) > MaxLength {
		return nil, ErrWrongLength
	}

	n := DeviceID(deviceID)

	return &n, nil
}
