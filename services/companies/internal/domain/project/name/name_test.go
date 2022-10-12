package name

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestString(t *testing.T) {
	req := require.New(t)
	testAddr := Name("ООО \"Лидген\"")
	tests := map[string]struct {
		input *Name
		want  string
	}{
		"success": {input: &testAddr, want: "ООО \"Лидген\""},
	}

	for name, testCase := range tests {
		t.Run(name, func(t *testing.T) {
			res := testCase.input.String()
			req.Equal(testCase.want, res)
		})
	}

}

func TestNewAdress(t *testing.T) {
	req := require.New(t)
	testName := Name("ООО \"Лидген\"")
	tests := map[string]struct {
		input string
		want  *Name
		err   error
	}{
		"success":      {input: "ООО \"Лидген\"", want: &testName, err: nil},
		"error length": {input: "г. ", want: nil, err: ErrWrongLength},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			res, err := NewName(testCase.input)
			req.Equal(testCase.want, res)
			req.Equal(testCase.err, err)
		})
	}
}
