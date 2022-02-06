package avatargenerator

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCat(t *testing.T) {
	err := os.MkdirAll("./out", 0777)
	assert.Nil(t, err)

	err = GetCat("cat1", "./out")
	assert.Nil(t, err)
}

func TestGetBird(t *testing.T) {
	err := os.MkdirAll("./out", 0777)
	assert.Nil(t, err)

	err = GetBird("bird2", "./out")
	assert.Nil(t, err)
}

func TestGetMobilizon(t *testing.T) {
	err := os.MkdirAll("./out", 0777)
	assert.Nil(t, err)

	err = GetMobilizon("mobilizon3", "./out")
	assert.Nil(t, err)
}
