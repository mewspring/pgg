// world is a tool which initializes and renders a simple game world.
package main

import (
	"image"
	"image/color"
	"image/draw"
	"log"

	"github.com/mewkiz/pkg/imgutil"
	"github.com/mewmew/pgg/grid"
	"github.com/mewmew/pgg/tileset"
	"github.com/mewmew/pgg/view"
)

func main() {
	err := world()
	if err != nil {
		log.Fatalln(err)
	}
}

// Tile width and height in the tile set.
const (
	TileWidth  = 32
	TileHeight = 32
)

// Tile identifiers.
const (
	Water tileset.TileID = 1
	Dirt  tileset.TileID = 2
	Grass tileset.TileID = 3
	Sand  tileset.TileID = 4
)

// world initializes and renders the game world.
func world() (err error) {
	// Initialize map.
	m := grid.NewMap(15, 15)

	initLevel(m)

	// Initialize camera.
	cam := view.Camera{
		TopLeft: grid.Loc(1, 1),
		Cols:    7,
		Rows:    7,
		Offset:  image.Pt(16, 16),
	}

	// Initialize tileset.
	ts, err := tileset.Open("tileset 1.png", TileWidth, TileHeight)
	if err != nil {
		return err
	}

	// Initialize world image.
	width := TileWidth * cam.Cols
	height := TileHeight * cam.Rows
	world := image.NewRGBA(image.Rect(0, 0, width, height))

	// Make background color lime.
	lime := image.NewUniform(color.RGBA{0x00, 0xFF, 0x00, 0xFF})
	draw.Draw(world, world.Bounds(), lime, image.ZP, draw.Over)

	// Draw loop.
	for col := 0; col < cam.Cols; col++ {
		for row := 0; row < cam.Rows; row++ {
			camCol := col + cam.TopLeft.Col
			camRow := row + cam.TopLeft.Row
			id := tileset.TileID(m[camCol][camRow])
			tile := ts.Tile(id)
			xmin := col*TileWidth - cam.Offset.X
			ymin := row*TileHeight - cam.Offset.Y
			dr := image.Rect(xmin, ymin, xmin+TileWidth, ymin+TileHeight)
			sp := tile.Bounds().Min
			draw.Draw(world, dr, tile, sp, draw.Over)
		}
	}

	// Output world image.
	err = imgutil.WriteFile("world.png", world)
	if err != nil {
		return err
	}

	return nil
}

// initLevel initializes the provided map with the tiles of a simple level.
func initLevel(m grid.Map) {
	// Col 0.
	m[0][0] = grid.Cell(Water)
	m[0][1] = grid.Cell(Water)
	m[0][2] = grid.Cell(Water)
	m[0][3] = grid.Cell(Water)
	m[0][4] = grid.Cell(Water)
	m[0][5] = grid.Cell(Water)
	m[0][6] = grid.Cell(Water)
	m[0][7] = grid.Cell(Water)
	m[0][8] = grid.Cell(Water)
	m[0][9] = grid.Cell(Water)
	m[0][10] = grid.Cell(Water)

	// Col 1.
	m[1][0] = grid.Cell(Water)
	m[1][1] = grid.Cell(Water)
	m[1][2] = grid.Cell(Water)
	m[1][3] = grid.Cell(Water)
	m[1][4] = grid.Cell(Water)
	m[1][5] = grid.Cell(Water)
	m[1][6] = grid.Cell(Water)
	m[1][7] = grid.Cell(Water)
	m[1][8] = grid.Cell(Water)
	m[1][9] = grid.Cell(Water)
	m[1][10] = grid.Cell(Water)

	// Col 2.
	m[2][0] = grid.Cell(Water)
	m[2][1] = grid.Cell(Water)
	m[2][2] = grid.Cell(Sand)
	m[2][3] = grid.Cell(Sand)
	m[2][4] = grid.Cell(Sand)
	m[2][5] = grid.Cell(Sand)
	m[2][6] = grid.Cell(Sand)
	m[2][7] = grid.Cell(Water)
	m[2][8] = grid.Cell(Water)
	m[2][9] = grid.Cell(Water)
	m[2][10] = grid.Cell(Water)

	// Col 3.
	m[3][0] = grid.Cell(Sand)
	m[3][1] = grid.Cell(Sand)
	m[3][2] = grid.Cell(Dirt)
	m[3][3] = grid.Cell(Dirt)
	m[3][4] = grid.Cell(Dirt)
	m[3][5] = grid.Cell(Dirt)
	m[3][6] = grid.Cell(Sand)
	m[3][7] = grid.Cell(Sand)
	m[3][8] = grid.Cell(Sand)
	m[3][9] = grid.Cell(Water)
	m[3][10] = grid.Cell(Water)

	// Col 4.
	m[4][0] = grid.Cell(Sand)
	m[4][1] = grid.Cell(Dirt)
	m[4][2] = grid.Cell(Dirt)
	m[4][3] = grid.Cell(Dirt)
	m[4][4] = grid.Cell(Dirt)
	m[4][5] = grid.Cell(Dirt)
	m[4][6] = grid.Cell(Dirt)
	m[4][7] = grid.Cell(Dirt)
	m[4][8] = grid.Cell(Sand)
	m[4][9] = grid.Cell(Sand)
	m[4][10] = grid.Cell(Water)

	// Col 5.
	m[5][0] = grid.Cell(Dirt)
	m[5][1] = grid.Cell(Dirt)
	m[5][2] = grid.Cell(Dirt)
	m[5][3] = grid.Cell(Grass)
	m[5][4] = grid.Cell(Grass)
	m[5][5] = grid.Cell(Grass)
	m[5][6] = grid.Cell(Dirt)
	m[5][7] = grid.Cell(Dirt)
	m[5][8] = grid.Cell(Dirt)
	m[5][9] = grid.Cell(Sand)
	m[5][10] = grid.Cell(Water)

	// Col 6.
	m[6][0] = grid.Cell(Dirt)
	m[6][1] = grid.Cell(Dirt)
	m[6][2] = grid.Cell(Dirt)
	m[6][3] = grid.Cell(Grass)
	m[6][4] = grid.Cell(Grass)
	m[6][5] = grid.Cell(Grass)
	m[6][6] = grid.Cell(Dirt)
	m[6][7] = grid.Cell(Dirt)
	m[6][8] = grid.Cell(Dirt)
	m[6][9] = grid.Cell(Sand)
	m[6][10] = grid.Cell(Water)

	// Col 7.
	m[7][0] = grid.Cell(Dirt)
	m[7][1] = grid.Cell(Dirt)
	m[7][2] = grid.Cell(Dirt)
	m[7][3] = grid.Cell(Grass)
	m[7][4] = grid.Cell(Grass)
	m[7][5] = grid.Cell(Grass)
	m[7][6] = grid.Cell(Dirt)
	m[7][7] = grid.Cell(Dirt)
	m[7][8] = grid.Cell(Dirt)
	m[7][9] = grid.Cell(Sand)
	m[7][10] = grid.Cell(Water)
}
