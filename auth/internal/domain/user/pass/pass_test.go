package pass

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

var (
	salt = "WcrMaOZhpYrpk79vFk"
)

func TestString(t *testing.T) {
	req := require.New(t)

	testPass, _ := NewPass("Password1!", salt)
	tests := map[string]struct {
		input *Pass
		want  string
	}{
		"success": {input: testPass, want: generateHashPassword("Password1!", salt)},
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
	testPass, _ := NewPass("Password12!", salt)
	tests := map[string]struct {
		input string
		want  *Pass
		err   error
	}{
		"success":                     {input: "Password12!", want: testPass, err: nil},
		"length err":                  {input: "Ps12!", want: nil, err: ErrPassLength},
		"no special characters err":   {input: "Password12", want: nil, err: ErrNoSpecialCharacter},
		"no numbers err":              {input: "Password!", want: nil, err: ErrNoNumber},
		"no upper case err":           {input: "password1!", want: nil, err: ErrNoUpperCaseLetter},
		"no lower case err":           {input: "PASSWORD1!", want: nil, err: ErrNoLowerCaseLetter},
		"no number and uppercase err": {input: "password!", want: nil, err: fmt.Errorf("%s:%s", ErrNoUpperCaseLetter, ErrNoNumber)},
	}

	for testPass, testCase := range tests {
		t.Run(testPass, func(t *testing.T) {
			res, err := NewPass(testCase.input, salt)
			req.Equal(testCase.want, res)
			req.Equal(testCase.err, err)
		})
	}
}
