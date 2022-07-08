package lib_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/wow-thats-mr-yahoo/go-test-gh-actions/internal/lib"
)

func TestUseAnother(t *testing.T) {
	res, err := lib.UseAnother()
	require.NoError(t, err)
	require.Equal(t, 65, res)
}
