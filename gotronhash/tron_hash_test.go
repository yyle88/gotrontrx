package gotronhash

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	m.Run()
	os.Exit(0)
}

func TestExample(t *testing.T) {
	var err error
	require.NoError(t, err)

	var b = true
	if !assert.True(t, b) {
		return
	}
}
