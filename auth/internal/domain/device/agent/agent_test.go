package agent

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestString(t *testing.T) {
	req := require.New(t)
	testAgent, _ := NewAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/52.0.2743.116 Safari/537.36 Edge/15.15063")
	tests := map[string]struct {
		input *Agent
		want  string
	}{
		"success": {input: testAgent, want: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/52.0.2743.116 Safari/537.36 Edge/15.15063"},
	}

	for agent, testCase := range tests {
		t.Run(agent, func(t *testing.T) {
			res := testCase.input.String()
			req.Equal(testCase.want, res)
		})
	}

}

func TestNewAgent(t *testing.T) {
	req := require.New(t)
	testAgent := Agent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/52.0.2743.116 Safari/537.36 Edge/15.15063")
	tests := map[string]struct {
		input string
		want  *Agent
		err   error
	}{
		"success":      {input: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/52.0.2743.116 Safari/537.36 Edge/15.15063", want: &testAgent, err: nil},
		"error length": {input: "Moz", want: nil, err: ErrWrongLength},
	}

	for testAgent, testCase := range tests {
		t.Run(testAgent, func(t *testing.T) {
			res, err := NewAgent(testCase.input)
			req.Equal(testCase.want, res)
			req.Equal(testCase.err, err)
		})
	}
}
