// Package tileset handles collections of one or more tile images using OpenGL.
package tileset

import (
	"image"

	"github.com/mewmew/glfw/win"
)

// A TileSet is a collection of one or more tile images, all of which have the
// same width and height.
type TileSet struct {
	// Tile set sprite sheet.
	img *win.Image
	// Tile width.
	TileWidth int
	// Tile height.
	TileHeight int
}

// Open opens the sprite sheet specified by imgPath and returns a tile set based
// upon it.
func Open(imgPath string, tileWidth, tileHeight int) (ts *TileSet, err error) {
	ts = &TileSet{
		TileWidth:  tileWidth,
		TileHeight: tileHeight,
	}
	ts.img, err = win.OpenImage(imgPath)
	if err != nil {
		return nil, err
	}
	return ts, nil
}

// A TileID uniquely identifies a tile image in a specific tile set. The zero
// value represents no tile image.
type TileID int

// IsValid returns true if the tile identifier is valid and false if it's the
// zero value.
func (id TileID) IsValid() bool {
	return id != 0
}

// tilePoint returns the top left point of the tile image in the tile set.
func (ts *TileSet) tilePoint(id TileID) image.Point {
	tsCols := ts.img.Width / ts.TileWidth
	i := int(id - 1)
	col := i % tsCols
	row := i / tsCols
	x := col * ts.TileWidth
	y := row * ts.TileHeight
	return image.Pt(x, y)
}

// DrawTile draws the tile image specified by id at the provided destination
// point dp.
func (ts *TileSet) DrawTile(id TileID, dp image.Point) {
	dr := image.Rect(dp.X, dp.Y, dp.X+ts.TileWidth, dp.Y+ts.TileHeight)
	sp := ts.tilePoint(id)
	ts.img.DrawRect(dr, sp)
}

// LastID returns the last tile identifier contained within the tile set. An
// empty tile set always returns the zero value.
func (ts *TileSet) LastID() (id TileID) {
	// TODO(u): ignore trailing empty tiles?
	tsCols := ts.img.Width / ts.TileWidth
	tsRows := ts.img.Height / ts.TileHeight
	id = TileID(tsCols * tsRows)
	return id
}
