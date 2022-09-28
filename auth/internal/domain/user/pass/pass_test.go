package pass

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestString(t *testing.T) {
	req := require.New(t)
	testPass, _ := NewPass("Password1!")
	tests := map[string]struct {
		input *Pass
		want  string
	}{
		"success": {input: testPass, want: "Password1!"},
	}

	for pass, testCase := range tests {
		t.Run(pass, func(t *testing.T) {
			res := testCase.input.String()
			req.Equal(testCase.want, res)
		})
	}

}

func TestNewPass(t *testing.T) {
	req := require.New(t)
	testPass, _ := NewPass("Password12!")
	tests := map[string]struct {
		input string
		want  *Pass
		err   error
	}{
		"success":                   {input: "Password12!", want: testPass, err: nil},
		"length err":                {input: "Pass", want: nil, err: ErrPassLength},
		"no special characters err": {input: "Password12", want: nil, err: ErrNoSpecialCharacter},
	}

	for testPass, testCase := range tests {
		t.Run(testPass, func(t *testing.T) {
			res, err := NewPass(testCase.input)
			req.Equal(testCase.want, res)
			req.Equal(testCase.err, err)
		})
	}
}
