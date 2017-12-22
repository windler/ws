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

	c := Repository()
	assert.Equal(t, "", c.Get(WsDir))

	c.Set(WsDir, "someText")
	c.Save()

	assert.Equal(t, "someText", c.Get(WsDir))
}
