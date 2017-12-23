package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/windler/workspacehero/internal/test"
)

func TestPrepareConfig(t *testing.T) {
	c, file := test.CreateTestContext(ConfigFlag)
	Repository(c)

	_, err := os.Stat(file.Name())
	assert.Nil(t, err)
}

func TestSaveConfig(t *testing.T) {
	c, _ := test.CreateTestContext(ConfigFlag)
	repo := Repository(c)

	assert.Equal(t, "", repo.WsDir)

	repo.WsDir = "someText"
	repo.Save()

	assert.Equal(t, "someText", repo.WsDir)
}
