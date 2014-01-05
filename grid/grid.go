// Package grid divides the game world into a series of contiguous grid cells
// using regular tessellation.
//
// Remember to specify the grid cell dimensions during init since they affect
// other logic.
package grid

import (
	"github.com/mewmew/pgg/tileset"
)

// A Map is a collection of cells that forms a complete map. It consists of a
// two-dimensional array with column and row indexes.
type Map [][]Cell

// NewMap returns a new map with the specified number of columns and rows.
func NewMap(cols, rows int) (m Map) {
	m = make([][]Cell, cols)
	for col := range m {
		m[col] = make([]Cell, rows)
	}
	return m
}

// Cols returns the number of columns in the map.
func (m Map) Cols() int {
	return len(m)
}

// Rows returns the number of rows in the map.
func (m Map) Rows() int {
	if m.Cols() < 1 {
		return 0
	}
	return len(m[0])
}

// A Cell corresponds to an individual grid cell which covers a portion of the
// grid. The area of a cell is specified by CellWidth and CellHeight.
type Cell tileset.TileID

// The width and height of grid cells. Remember to specify these dimensions
// during init since they affect other logic.
var (
	CellWidth  = 32
	CellHeight = 32
)

// TODO(u): define Col and Row types?

// A Location specifies a precise grid location corresponding to a specific cell
// of the map.
type Location struct {
	// The grid column.
	Col int
	// The grid row.
	Row int
}

// Loc is shorthand for Location{col, row}.
func Loc(col, row int) (loc Location) {
	loc = Location{
		Col: col,
		Row: row,
	}
	return loc
}
