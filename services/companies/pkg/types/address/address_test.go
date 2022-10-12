package address

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestString(t *testing.T) {
	req := require.New(t)
	testAddr := Address("г. Майкоп, ул. Краснооктябрьская 10А")
	tests := map[string]struct {
		input *Address
		want  string
	}{
		"success": {input: &testAddr, want: "г. Майкоп, ул. Краснооктябрьская 10А"},
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
	testAddr := Address("г. Майкоп, ул. Краснооктябрьская 10А")
	tests := map[string]struct {
		input string
		want  *Address
		err   error
	}{
		"success":      {input: "г. Майкоп, ул. Краснооктябрьская 10А", want: &testAddr, err: nil},
		"error length": {input: "г. ", want: nil, err: ErrWrongLength},
	}

	for name, testCase := range tests {
		t.Run(name, func(t *testing.T) {
			res, err := NewAdress(testCase.input)
			req.Equal(testCase.want, res)
			req.Equal(testCase.err, err)
		})
	}
}
