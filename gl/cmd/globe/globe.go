// globe is a tool which initializes and renders a simple game world using
// OpenGL.
package main

import (
	"image"
	"log"
	"runtime"
	"time"

	"github.com/mewmew/pgg/gl/tileset"
	"github.com/mewmew/pgg/grid"
	"github.com/mewmew/pgg/view"
	"github.com/mewmew/we"
	"github.com/mewmew/win"
)

func init() {
	// Specify the width and height of grid cells.
	grid.CellWidth = 48
	grid.CellHeight = 48
}

func main() {
	err := globe()
	if err != nil {
		log.Fatalln(err)
	}
}

// Number of columns and rows of the entire map.
const (
	MapCols = 9
	MapRows = 11
)

// Tile identifiers.
const (
	Grass  tileset.TileID = 1
	Sand   tileset.TileID = 2
	Water  tileset.TileID = 3
	Gravel tileset.TileID = 4
)

// fps corresponds to the number of frames per second that should be drawn.
const fps = 60

// globe initializes and renders the game world.
func globe() (err error) {
	// OpenGL requires a dedicated OS thread.
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	// Initialize map.
	m := grid.NewMap(MapCols, MapRows)
	initLevel(m)

	// Initialize view.
	viewCols := 6
	viewRows := 6
	width := viewCols * grid.CellWidth
	height := viewRows * grid.CellHeight
	mapWidth := MapCols * grid.CellWidth
	mapHeight := MapRows * grid.CellHeight
	end := image.Pt(mapWidth, mapHeight)
	v := view.NewView(width, height, end)

	// Initialize window.
	err = win.Open(width, height)
	if err != nil {
		return err
	}
	defer win.Close()

	// Register that we are interested in receiving the following events.
	win.EnableCloseChan()
	win.EnableKeyPressChan()
	win.EnableKeyRepeatChan()

	// Initialize tileset.
	tileWidth := grid.CellWidth
	tileHeight := grid.CellHeight
	ts, err := tileset.Open("tileset 2.png", tileWidth, tileHeight)
	if err != nil {
		return err
	}

	// drawTile draws the tile at the specified column and row.
	drawTile := func(col, row int) {
		viewCol := col + v.Col()
		viewRow := row + v.Row()
		id := tileset.TileID(m[viewCol][viewRow])
		x := col*grid.CellWidth - v.X()
		y := row*grid.CellHeight - v.Y()
		dp := image.Pt(x, y)
		ts.DrawTile(id, dp)
	}

	c := time.Tick(time.Second / fps)
	for {
		// Draw loop.
		for col := 0; col < v.Cols(); col++ {
			for row := 0; row < v.Rows(); row++ {
				drawTile(col, row)
			}
		}

		// Swap buffers to display all drawings since last screen update.
		win.SwapBuffers()

		select {
		case <-win.CloseChan:
			// handle close events.
			return nil
		case e := <-win.KeyPressChan:
			handleKeyPress(e, v)
		case e := <-win.KeyRepeatChan:
			handleKeyPress(we.KeyPress(e), v)
		case <-c:
			// very simple implementation to update 60 times per second.
			continue
		}
		// Keep the frame rate constant.
		<-c
	}
}

// handleKeyPress handles key press events.
func handleKeyPress(e we.KeyPress, v *view.View) {
	// handle key press events.
	switch e.Key {
	case we.KeyUp:
		v.Move(image.Pt(0, -2))
	case we.KeyDown:
		v.Move(image.Pt(0, 2))
	case we.KeyRight:
		v.Move(image.Pt(2, 0))
	case we.KeyLeft:
		v.Move(image.Pt(-2, 0))
	}
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
	m[3][2] = grid.Cell(Gravel)
	m[3][3] = grid.Cell(Gravel)
	m[3][4] = grid.Cell(Gravel)
	m[3][5] = grid.Cell(Gravel)
	m[3][6] = grid.Cell(Sand)
	m[3][7] = grid.Cell(Sand)
	m[3][8] = grid.Cell(Sand)
	m[3][9] = grid.Cell(Water)
	m[3][10] = grid.Cell(Water)

	// Col 4.
	m[4][0] = grid.Cell(Sand)
	m[4][1] = grid.Cell(Gravel)
	m[4][2] = grid.Cell(Gravel)
	m[4][3] = grid.Cell(Gravel)
	m[4][4] = grid.Cell(Gravel)
	m[4][5] = grid.Cell(Gravel)
	m[4][6] = grid.Cell(Gravel)
	m[4][7] = grid.Cell(Gravel)
	m[4][8] = grid.Cell(Sand)
	m[4][9] = grid.Cell(Sand)
	m[4][10] = grid.Cell(Water)

	// Col 5.
	m[5][0] = grid.Cell(Gravel)
	m[5][1] = grid.Cell(Gravel)
	m[5][2] = grid.Cell(Gravel)
	m[5][3] = grid.Cell(Grass)
	m[5][4] = grid.Cell(Grass)
	m[5][5] = grid.Cell(Grass)
	m[5][6] = grid.Cell(Gravel)
	m[5][7] = grid.Cell(Gravel)
	m[5][8] = grid.Cell(Gravel)
	m[5][9] = grid.Cell(Sand)
	m[5][10] = grid.Cell(Water)

	// Col 6.
	m[6][0] = grid.Cell(Gravel)
	m[6][1] = grid.Cell(Gravel)
	m[6][2] = grid.Cell(Gravel)
	m[6][3] = grid.Cell(Grass)
	m[6][4] = grid.Cell(Grass)
	m[6][5] = grid.Cell(Grass)
	m[6][6] = grid.Cell(Gravel)
	m[6][7] = grid.Cell(Gravel)
	m[6][8] = grid.Cell(Gravel)
	m[6][9] = grid.Cell(Sand)
	m[6][10] = grid.Cell(Water)

	// Col 7.
	m[7][0] = grid.Cell(Gravel)
	m[7][1] = grid.Cell(Gravel)
	m[7][2] = grid.Cell(Gravel)
	m[7][3] = grid.Cell(Grass)
	m[7][4] = grid.Cell(Grass)
	m[7][5] = grid.Cell(Grass)
	m[7][6] = grid.Cell(Gravel)
	m[7][7] = grid.Cell(Gravel)
	m[7][8] = grid.Cell(Gravel)
	m[7][9] = grid.Cell(Sand)
	m[7][10] = grid.Cell(Water)

	// Col 8.
	m[8][0] = grid.Cell(Water)
	m[8][1] = grid.Cell(Sand)
	m[8][2] = grid.Cell(Water)
	m[8][3] = grid.Cell(Grass)
	m[8][4] = grid.Cell(Water)
	m[8][5] = grid.Cell(Sand)
	m[8][6] = grid.Cell(Water)
	m[8][7] = grid.Cell(Grass)
	m[8][8] = grid.Cell(Water)
	m[8][9] = grid.Cell(Sand)
	m[8][10] = grid.Cell(Water)
}
