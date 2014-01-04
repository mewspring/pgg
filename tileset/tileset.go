// Package tileset handles collections of one or more tile images.
package tileset

import (
	"image"

	"github.com/mewkiz/pkg/imgutil"
)

// A TileSet is a collection of one or more tile images.
type TileSet struct {
	// Tileset sprite sheet.
	imgutil.SubImager
	// Tileset width.
	width int
	// Tileset height.
	height int
	// Tile width.
	tileWidth int
	// Tile height.
	tileHeight int
}

// New returns a tile set based on the provided sprite sheet img.
func New(img image.Image, tileWidth, tileHeight int) (ts *TileSet) {
	ts = &TileSet{
		tileWidth:  tileWidth,
		tileHeight: tileHeight,
	}
	ts.SubImager = imgutil.SubFallback(img)
	bounds := ts.Bounds()
	ts.width = bounds.Dx()
	ts.height = bounds.Dy()
	return ts
}

// Open opens the sprite sheet specified by imgPath and returns a tile set based
// on it.
func Open(imgPath string, tileWidth, tileHeight int) (ts *TileSet, err error) {
	img, err := imgutil.ReadFile(imgPath)
	if err != nil {
		return nil, err
	}
	ts = New(img, tileWidth, tileHeight)
	return ts, nil
}

// A TileID uniquely identifies a tile image. The zero value represents no tile
// image.
type TileID int

// IsValid returns true if the tile identifier is valid and false if it's the
// zero value.
func (id TileID) IsValid() bool {
	return id != 0
}

// tileRect returns the bounding rectangle in the sprite sheet of the tile image
// specified by id.
func (ts *TileSet) tileRect(id TileID) image.Rectangle {
	tsCols := ts.width / ts.tileWidth
	i := int(id - 1)
	col := i % tsCols
	row := i / tsCols
	x := col * ts.tileWidth
	y := row * ts.tileHeight
	return image.Rect(x, y, x+ts.tileWidth, y+ts.tileHeight)
}

// Tile returns the tile image specified by id from the tile set.
func (ts *TileSet) Tile(id TileID) image.Image {
	rect := ts.tileRect(id)
	return ts.SubImage(rect)
}

// LastID returns the last tile ID contained within the tile set.
func (ts *TileSet) LastID() (id TileID) {
	tsCols := ts.width / ts.tileWidth
	tsRows := ts.height / ts.tileHeight
	id = TileID(tsCols * tsRows)
	return id
}
