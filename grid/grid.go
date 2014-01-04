// Package grid divides the game world into a series of contiguous grid cells
// using regular tessellation.
package grid

import (
	"github.com/mewmew/pgg/tileset"
)

type Map [][]Cell

func NewMap(cols, rows int) (m Map) {
	m = make([][]Cell, cols)
	for col := range m {
		m[col] = make([]Cell, rows)
	}
	return m
}

type Cell tileset.TileID

type Location struct {
	Col int
	Row int
}

func Loc(col, row int) (loc Location) {
	loc = Location{
		Col: col,
		Row: row,
	}
	return loc
}
