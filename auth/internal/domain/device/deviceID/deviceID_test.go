package deviceID

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestString(t *testing.T) {
	req := require.New(t)
	testDeviceID, _ := NewDeviceID("00000000-000000000000000")
	tests := map[string]struct {
		input *DeviceID
		want  string
	}{
		"success": {input: testDeviceID, want: "00000000-000000000000000"},
	}

	for deviceID, testCase := range tests {
		t.Run(deviceID, func(t *testing.T) {
			res := testCase.input.String()
			req.Equal(testCase.want, res)
		})
	}

}

func TestNewDeviceID(t *testing.T) {
	req := require.New(t)
	testDeviceID := DeviceID("00000000-000000000000000")
	tests := map[string]struct {
		input string
		want  *DeviceID
		err   error
	}{
		"success":      {input: "00000000-000000000000000", want: &testDeviceID, err: nil},
		"error length": {input: "00", want: nil, err: ErrWrongLength},
	}

	for testDeviceID, testCase := range tests {
		t.Run(testDeviceID, func(t *testing.T) {
			res, err := NewDeviceID(testCase.input)
			req.Equal(testCase.want, res)
			req.Equal(testCase.err, err)
		})
	}
}
