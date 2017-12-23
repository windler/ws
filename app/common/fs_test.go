package common

import "testing"
import "github.com/stretchr/testify/assert"
import "strings"

func TestEsnureDir(t *testing.T) {
	dir := "asd"
	assert.Equal(t, "asd/", EnsureDirFormat(dir))

	dir = "asd/"
	assert.Equal(t, "asd/", EnsureDirFormat(dir))

	dir = "/asd/"
	assert.Equal(t, "/asd/", EnsureDirFormat(dir))

	dir = "~/asd/"
	assert.NotEqual(t, "~/asd/", EnsureDirFormat(dir))
	assert.False(t, strings.Contains(EnsureDirFormat(dir), "~"))
	assert.True(t, strings.Contains(EnsureDirFormat(dir), "/asd/"))
}

func TestEsnureFile(t *testing.T) {
	file := "asd"
	assert.Equal(t, "asd", EnsureFileFormat(file))

	file = "~/asd"
	assert.NotEqual(t, "~/asd", EnsureDirFormat(file))
	assert.False(t, strings.Contains(EnsureDirFormat(file), "~"))
	assert.True(t, strings.Contains(EnsureDirFormat(file), "/asd"))
}
