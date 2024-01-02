package csvutils

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLoadCSVForBeans(t *testing.T) {
	beans, err := LoadCSVForBeans()
	require.NoError(t, err)
	require.NotEmpty(t, beans)
	assert.Len(t, beans, 5)
}
