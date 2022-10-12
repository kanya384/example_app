package phone

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestString(t *testing.T) {
	req := require.New(t)
	testPhone, _ := NewPhone("+7 (962) 777-00-00")
	tests := map[string]struct {
		input *Phone
		want  string
	}{
		"success": {input: testPhone, want: "79627770000"},
	}

	for phone, testCase := range tests {
		t.Run(phone, func(t *testing.T) {
			res := testCase.input.String()
			req.Equal(testCase.want, res)
		})
	}

}

func TestNewPhone(t *testing.T) {
	req := require.New(t)
	testPhone, _ := NewPhone("+7 (962) 777-00-00")
	tests := map[string]struct {
		input string
		want  *Phone
		err   error
	}{
		"success":      {input: "+7 (962) 777-00-00", want: testPhone, err: nil},
		"error length": {input: "tt ", want: nil, err: ErrNotValidPhone},
	}

	for testPhone, testCase := range tests {
		t.Run(testPhone, func(t *testing.T) {
			res, err := NewPhone(testCase.input)
			req.Equal(testCase.want, res)
			req.Equal(testCase.err, err)
		})
	}
}
