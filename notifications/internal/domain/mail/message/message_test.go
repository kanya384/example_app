package message

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestString(t *testing.T) {
	req := require.New(t)
	testMessage := Message("Lorem Ipsum is simply dummy message of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy message ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged.")
	tests := map[string]struct {
		input *Message
		want  string
	}{
		"success": {input: &testMessage, want: "Lorem Ipsum is simply dummy message of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy message ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged."},
	}

	for message, testCase := range tests {
		t.Run(message, func(t *testing.T) {
			res := testCase.input.String()
			req.Equal(testCase.want, res)
		})
	}

}

func TestNewMessage(t *testing.T) {
	req := require.New(t)
	testName := Message("Lorem Ipsum is simply dummy message of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy message ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged.")
	tests := map[string]struct {
		input string
		want  *Message
		err   error
	}{
		"success":      {input: "Lorem Ipsum is simply dummy message of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy message ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged.", want: &testName, err: nil},
		"error length": {input: "г. ", want: nil, err: ErrWrongLength},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			res, err := NewMessage(testCase.input)
			req.Equal(testCase.want, res)
			req.Equal(testCase.err, err)
		})
	}
}
