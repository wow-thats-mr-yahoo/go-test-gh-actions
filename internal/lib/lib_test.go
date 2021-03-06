package lib_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/wow-thats-mr-yahoo/go-test-gh-actions/internal/lib"
)

func TestDoStuff(t *testing.T) {
	require.Equal(t, 3422, lib.DoStuff())
}

func TestDoStuff_AnotherWay(t *testing.T) {
	require.NotEqual(t, 3421, lib.DoStuff())
}
