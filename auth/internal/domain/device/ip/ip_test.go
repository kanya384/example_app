package ip

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestString(t *testing.T) {
	req := require.New(t)
	testIp, _ := NewIp("8.8.8.8")
	tests := map[string]struct {
		input *Ip
		want  string
	}{
		"success": {input: testIp, want: "8.8.8.8"},
	}

	for ip, testCase := range tests {
		t.Run(ip, func(t *testing.T) {
			res := testCase.input.String()
			req.Equal(testCase.want, res)
		})
	}

}

func TestNewIp(t *testing.T) {
	req := require.New(t)
	testIp := Ip("8.8.8.8")
	tests := map[string]struct {
		input string
		want  *Ip
		err   error
	}{
		"success":      {input: "8.8.8.8", want: &testIp, err: nil},
		"error length": {input: "00", want: nil, err: ErrWrongLength},
	}

	for testIp, testCase := range tests {
		t.Run(testIp, func(t *testing.T) {
			res, err := NewIp(testCase.input)
			req.Equal(testCase.want, res)
			req.Equal(testCase.err, err)
		})
	}
}
