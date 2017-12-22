package config

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrepareConfig(t *testing.T) {
	file, _ := ioutil.TempFile("", "projherotest")
	Prepare(file.Name())

	_, err := os.Stat(file.Name())
	assert.Nil(t, err)
}

func TestSaveConfig(t *testing.T) {
	file, _ := ioutil.TempFile("", "projherotest")
	Prepare(file.Name())

	c := Get()
	assert.Equal(t, "", c.WsDir)

	c.WsDir = "someText"
	c.Save()

	cNew := Get()
	assert.Equal(t, "someText", cNew.WsDir)
}
