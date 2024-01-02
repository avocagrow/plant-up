package randutils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandStr(t *testing.T) {
	str1 := RandStr(10)
	assert.Len(t, str1, 20)
	str2 := RandStr(16)
	assert.Len(t, str2, 32)
	assert.NotEqual(t, str2, str1)
	str3 := RandStr(5)
	assert.Len(t, str3, 16)
	assert.NotEqual(t, str3, str1)
}
