package refreshToken

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestString(t *testing.T) {
	req := require.New(t)
	testRefreshToken, _ := NewFromString("123e4567-e89b-12d3-a456-426614174000")
	tests := map[string]struct {
		input *RefreshToken
		want  string
	}{
		"success": {input: testRefreshToken, want: "123e4567-e89b-12d3-a456-426614174000"},
	}

	for refreshToken, testCase := range tests {
		t.Run(refreshToken, func(t *testing.T) {
			res := testCase.input.String()
			req.Equal(testCase.want, res)
		})
	}

}

func TestNewFromString(t *testing.T) {
	req := require.New(t)
	testRefreshToken, _ := NewFromString("123e4567-e89b-12d3-a456-426614174000")
	tests := map[string]struct {
		input string
		want  *RefreshToken
		err   error
	}{
		"success":      {input: "123e4567-e89b-12d3-a456-426614174000", want: testRefreshToken, err: nil},
		"error length": {input: "123", want: nil, err: ErrWrongLength},
	}

	for testRefreshToken, testCase := range tests {
		t.Run(testRefreshToken, func(t *testing.T) {
			res, err := NewFromString(testCase.input)
			req.Equal(testCase.want, res)
			req.Equal(testCase.err, err)
		})
	}
}

func TestNew(t *testing.T) {
	req := require.New(t)

	t.Run("create token success", func(t *testing.T) {
		res, err := New()
		req.Empty(err)
		req.NotEmpty(res)
	})
}
