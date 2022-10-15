package auth

import (
	"errors"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestNewManager(t *testing.T) {
	req := require.New(t)
	tests := map[string]struct {
		signingkey string
		err        error
	}{
		"int":    {signingkey: "asddsfjkjsfii", err: nil},
		"string": {signingkey: "", err: errors.New(ErrEmptySigningKey)},
	}

	for name, testCase := range tests {
		t.Run(name, func(t *testing.T) {
			_, err := NewManager(testCase.signingkey)
			req.Equal(testCase.err, err)
		})
	}
}

func TestNewJWT(t *testing.T) {
	req := require.New(t)
	manager, err := NewManager("asddsfjkjsfii")
	if err != nil {
		return
	}

	tests := map[string]struct {
		input    JwtClaims
		duration time.Duration
		err      error
	}{
		"succsess": {input: JwtClaims{UserID: uuid.New().String(), UserName: "testName", UserRole: "administrator"}, duration: time.Duration(time.Second * 1), err: nil},
	}

	for name, testCase := range tests {
		t.Run(name, func(t *testing.T) {
			_, err := manager.NewJWT(testCase.input, testCase.duration)
			req.Equal(testCase.err, err)
		})
	}
}

func TestParse(t *testing.T) {
	req := require.New(t)
	manager, err := NewManager("asddsfjkjsfii")
	if err != nil {
		return
	}

	testClaims := JwtClaims{UserID: uuid.New().String(), UserName: "testName", UserRole: "administrator"}

	testToken, _ := manager.NewJWT(testClaims, time.Hour)

	t.Run("succsess", func(t *testing.T) {
		claims, err := manager.Parse(testToken)
		req.Equal(nil, err)
		req.Equal(testClaims, claims)
	})

	testTokenExpired, _ := manager.NewJWT(testClaims, -time.Duration(time.Second*10))

	t.Run("is expired", func(t *testing.T) {
		_, err := manager.Parse(testTokenExpired)
		req.NotNil(err)
	})

	t.Run("not valid token", func(t *testing.T) {
		_, err := manager.Parse("asdadas")
		req.NotNil(err)
	})
}
