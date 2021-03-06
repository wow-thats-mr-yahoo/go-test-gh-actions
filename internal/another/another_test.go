package another_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/wow-thats-mr-yahoo/go-test-gh-actions/internal/another"
)

func TestSomeVeryComplicatedFunc(t *testing.T) {
	cfg := another.SomeVeryComplicatedFuncConfig{
		A: 12,
		B: 43,
		C: 42,
		D: 32,
	}

	res := another.SomeVeryComplicatedFunc(cfg)
	expected := 12 + 43 + 42 - 32
	require.Equal(t, expected, res)
}
