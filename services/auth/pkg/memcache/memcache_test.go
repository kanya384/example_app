package memcache

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

var (
	memcache = New(time.Duration(time.Second*60*60*5), time.Duration(time.Minute*5))
)

func TestSet(t *testing.T) {
	req := require.New(t)
	tests := map[string]struct {
		key string
		val interface{}
	}{
		"int":    {key: "int", val: 1},
		"string": {key: "string", val: "string"},
	}

	for name, testCase := range tests {
		t.Run(name, func(t *testing.T) {
			memcache.Set(testCase.key, testCase.val, time.Duration(time.Second))

			memcache.Lock()
			defer memcache.Unlock()
			res := memcache.items[testCase.key]
			req.Equal(res.Value, testCase.val)
		})
	}
}

func TestGet(t *testing.T) {
	req := require.New(t)

	memcache.Set("test", "test", time.Duration(time.Second))
	tests := map[string]struct {
		key  string
		want error
	}{
		"ok":        {key: "test", want: nil},
		"not found": {key: "test2", want: ErrNotFound},
	}

	for name, testCase := range tests {
		t.Run(name, func(t *testing.T) {
			_, err := memcache.Get(testCase.key)
			req.Equal(err, testCase.want)
		})
	}
}

func TestDelete(t *testing.T) {
	req := require.New(t)
	memcache.Set("delete", "test", time.Duration(time.Second))

	tests := map[string]struct {
		key  string
		want error
	}{
		"ok":        {key: "delete", want: nil},
		"not found": {key: "test2", want: ErrNotFound},
	}

	for name, testCase := range tests {
		t.Run(name, func(t *testing.T) {
			err := memcache.Delete(testCase.key)
			req.Equal(err, testCase.want)
		})
	}
}
