package inn

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestString(t *testing.T) {
	req := require.New(t)
	var innVal uint32 = 1234567890
	testInn := Inn(innVal)
	tests := map[string]struct {
		input *Inn
		want  string
	}{
		"success": {input: &testInn, want: "1234567890"},
	}

	for name, testCase := range tests {
		t.Run(name, func(t *testing.T) {
			res := testCase.input.String()
			req.Equal(testCase.want, res)
		})
	}

}

func TestNewInn(t *testing.T) {
	req := require.New(t)
	var innVal uint32 = 1234567890
	testInn := Inn(innVal)
	tests := map[string]struct {
		input uint32
		want  *Inn
		err   error
	}{
		"success":      {input: 1234567890, want: &testInn, err: nil},
		"error length": {input: 123, want: nil, err: ErrWrongLength},
	}

	for name, testCase := range tests {
		t.Run(name, func(t *testing.T) {
			res, err := NewInn(testCase.input)
			req.Equal(testCase.want, res)
			req.Equal(testCase.err, err)
		})
	}
}
