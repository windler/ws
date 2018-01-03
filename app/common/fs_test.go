package common

import "testing"
import "github.com/stretchr/testify/assert"
import "strings"

func TestEsnureDir(t *testing.T) {
	dir := "ws"
	assert.Equal(t, "ws/", EnsureDirFormat(dir))

	dir = "ws/"
	assert.Equal(t, "ws/", EnsureDirFormat(dir))

	dir = "/ws/"
	assert.Equal(t, "/ws/", EnsureDirFormat(dir))

	dir = "~/ws/"
	assert.NotEqual(t, "~/ws/", EnsureDirFormat(dir))
	assert.False(t, strings.Contains(EnsureDirFormat(dir), "~"))
	assert.True(t, strings.Contains(EnsureDirFormat(dir), "/ws/"))
}

func TestEsnureFile(t *testing.T) {
	file := "ws"
	assert.Equal(t, "ws", EnsureFileFormat(file))

	file = "~/ws"
	assert.NotEqual(t, "~/ws", EnsureDirFormat(file))
	assert.False(t, strings.Contains(EnsureDirFormat(file), "~"))
	assert.True(t, strings.Contains(EnsureDirFormat(file), "/ws"))
}
