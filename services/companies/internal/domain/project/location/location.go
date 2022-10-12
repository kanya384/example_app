package location

import (
	"fmt"
)

const (
	MinLatitude  = -90
	MaxLatitude  = 90
	MinLongitude = -180
	MaxLongitude = 180
)

var (
	ErrLatiude   = fmt.Errorf("latiude value must be between or equal %d and %d", MinLatitude, MaxLatitude)
	ErrLongitude = fmt.Errorf("longitude value must be between or equal %d and %d", MinLongitude, MaxLongitude)
)

type Location struct {
	longitude float32
	latitude  float32
}

func NewLocation(longitude, latitude float32) (*Location, error) {
	if latitude < MinLatitude || latitude > MaxLatitude {
		return nil, ErrLatiude
	}

	if longitude < MinLongitude || longitude > MaxLongitude {
		return nil, ErrLongitude
	}

	n := Location{
		longitude: longitude,
		latitude:  latitude,
	}

	return &n, nil
}

func (l Location) Longitude() float32 {
	return l.longitude
}

func (l Location) Latitude() float32 {
	return l.latitude
}
