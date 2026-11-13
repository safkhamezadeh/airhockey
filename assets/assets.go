package assets

import (
	"bytes"
	_ "embed"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Asset int

const (
	Puck Asset = iota
	Mallet
)

//go:embed puck.png
var puckPNG []byte

//go:embed paddle.png
var malletPNG []byte

var assets = map[Asset][]byte{
	Puck:   puckPNG,
	Mallet: malletPNG,
}

func LoadAsset(asset Asset) *ebiten.Image {
	img, _, err := ebitenutil.NewImageFromReader(bytes.NewReader(assets[asset]))
	if err != nil {
		log.Fatal(err)
	}
	return img
}
