// Package grid divides the game world into a series of contiguous grid cells
// using regular tessellation.
package grid

import (
	"github.com/mewmew/pgg/tileset"
)

type Map [][]Cel

type Cell tileset.TileID
