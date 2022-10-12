package subject

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestString(t *testing.T) {
	req := require.New(t)
	testSubject := Subject("Lorem Ipsum is simply dummy")
	tests := map[string]struct {
		input *Subject
		want  string
	}{
		"success": {input: &testSubject, want: "Lorem Ipsum is simply dummy"},
	}

	for message, testCase := range tests {
		t.Run(message, func(t *testing.T) {
			res := testCase.input.String()
			req.Equal(testCase.want, res)
		})
	}

}

func TestNewSubject(t *testing.T) {
	req := require.New(t)
	testSubject := Subject("Lorem Ipsum is simply dummy")
	tests := map[string]struct {
		input string
		want  *Subject
		err   error
	}{
		"success":      {input: "Lorem Ipsum is simply dummy", want: &testSubject, err: nil},
		"error length": {input: "г. ", want: nil, err: ErrWrongLength},
	}

	for testSubject, testCase := range tests {
		t.Run(testSubject, func(t *testing.T) {
			res, err := NewSubject(testCase.input)
			req.Equal(testCase.want, res)
			req.Equal(testCase.err, err)
		})
	}
}
