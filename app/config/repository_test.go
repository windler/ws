package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/windler/ws/internal/test"
)

func TestPrepareConfig(t *testing.T) {
	_, file := test.CreateTestContext(ConfigFlag)
	Repository()

	_, err := os.Stat(file.Name())
	assert.Nil(t, err)
}

func TestSaveConfig(t *testing.T) {
	test.CreateTestContext(ConfigFlag)
	repo := Repository()

	assert.Equal(t, "", repo.WsDir)

	repo.WsDir = "someText"
	repo.Save()

	assert.Equal(t, "someText", repo.WsDir)
}
