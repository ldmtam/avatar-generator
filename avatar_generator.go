package avatargenerator

import (
	"errors"
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"math/rand"
	"os"
	"path"

	"github.com/spaolacci/murmur3"
)

const imgSize = 256

var (
	AvatarType_CAT       = "cat"
	AvatarType_BIRD      = "bird"
	AvatarType_MOBILIZON = "mobilizon"
)

type part struct {
	name string
	id   int
}

func GetCat(seed, outPath string) error {
	id := murmur3.Sum64([]byte(seed))
	rand.Seed(int64(id))

	parts := []*part{
		{name: "body", id: rand.Intn(15) + 1},
		{name: "fur", id: rand.Intn(10) + 1},
		{name: "eyes", id: rand.Intn(15) + 1},
		{name: "mouth", id: rand.Intn(10) + 1},
		{name: "accessorie", id: rand.Intn(20) + 1},
	}

	return createAvatar(id, AvatarType_CAT, outPath, parts)
}

func GetBird(seed, outPath string) error {
	id := murmur3.Sum64([]byte(seed))
	rand.Seed(int64(id))

	parts := []*part{
		{name: "tail", id: rand.Intn(9) + 1},
		{name: "hoop", id: rand.Intn(10) + 1},
		{name: "body", id: rand.Intn(9) + 1},
		{name: "wing", id: rand.Intn(9) + 1},
		{name: "eyes", id: rand.Intn(9) + 1},
		{name: "bec", id: rand.Intn(9) + 1},
		{name: "accessorie", id: rand.Intn(20) + 1},
	}

	return createAvatar(id, AvatarType_BIRD, outPath, parts)
}

func GetMobilizon(seed, outPath string) error {
	id := murmur3.Sum64([]byte(seed))
	rand.Seed(int64(id))

	parts := []*part{
		{name: "body", id: rand.Intn(25) + 1},
		{name: "nose", id: rand.Intn(10) + 1},
		{name: "tail", id: rand.Intn(5) + 1},
		{name: "eyes", id: rand.Intn(10) + 1},
		{name: "mouth", id: rand.Intn(10) + 1},
		{name: "accessories", id: rand.Intn(20) + 1},
		{name: "misc", id: rand.Intn(20) + 1},
		{name: "hat", id: rand.Intn(20) + 1},
	}

	return createAvatar(id, AvatarType_MOBILIZON, outPath, parts)
}

func getAvatarPath(typ string) (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return path.Join(dir, fmt.Sprintf("/avatars/%v", typ)), nil
}

func createAvatar(id uint64, typ, outFolder string, parts []*part) error {
	avatarImg := image.NewRGBA(image.Rect(0, 0, imgSize, imgSize))

	var imgFolder string
	var err error
	if typ == AvatarType_CAT || typ == AvatarType_BIRD || typ == AvatarType_MOBILIZON {
		imgFolder, err = getAvatarPath(typ)
		if err != nil {
			return err
		}
	} else {
		return errors.New("type is not valid")
	}

	for _, p := range parts {
		partFileName := fmt.Sprintf("%v_%v.png", p.name, p.id)
		partFile, err := os.Open(path.Join(imgFolder, partFileName))
		if err != nil {
			return err
		}
		defer partFile.Close()

		partImg, err := png.Decode(partFile)
		if err != nil {
			return err
		}

		draw.Draw(avatarImg, avatarImg.Bounds(), partImg, image.Point{}, draw.Over)
	}

	avatarFileName := fmt.Sprintf("%v_%v.png", typ, id)
	avatarFile, err := os.Create(path.Join(outFolder, avatarFileName))
	if err != nil {
		return err
	}
	defer avatarFile.Close()

	return png.Encode(avatarFile, avatarImg)
}
