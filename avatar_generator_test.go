package avatargenerator

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCat(t *testing.T) {
	err := os.MkdirAll("./out", 0777)
	assert.Nil(t, err)

	err = GetCat(AvatarType_CAT, "./out")
	assert.Nil(t, err)
}

func TestGetBird(t *testing.T) {
	err := os.MkdirAll("./out", 0777)
	assert.Nil(t, err)

	err = GetBird(AvatarType_BIRD, "./out")
	assert.Nil(t, err)
}

func TestGetMobilizon(t *testing.T) {
	err := os.MkdirAll("./out", 0777)
	assert.Nil(t, err)

	err = GetMobilizon(AvatarType_MOBILIZON, "./out")
	assert.Nil(t, err)
}
