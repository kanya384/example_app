package location

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewLocation(t *testing.T) {
	req := require.New(t)

	t.Run("success", func(t *testing.T) {
		var longitude float32 = 44.55
		var latitude float32 = 55.55
		location, err := NewLocation(longitude, latitude)
		req.Equal(longitude, location.Longitude())
		req.Equal(latitude, location.Latitude())
		req.Equal(err, nil)
	})

	t.Run("error longitude", func(t *testing.T) {
		var longitude float32 = -181
		var latitude float32 = 55.55
		_, err := NewLocation(longitude, latitude)
		req.Equal(err, ErrLongitude)
	})

	t.Run("error latitude", func(t *testing.T) {
		var longitude float32 = 89
		var latitude float32 = -91
		_, err := NewLocation(longitude, latitude)
		req.Equal(err, ErrLatiude)
	})
}
