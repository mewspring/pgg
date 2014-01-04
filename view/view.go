// Package view supervises the visible portion of the screen.
package view

import (
	"image"

	"github.com/mewmew/pgg/grid"
)

type Camera struct {
	// The top left location visible by the camera.
	TopLeft grid.Location
	// Pixel offset from the top left grid cell.
	Offset image.Point
	// Number of columns and rows visible by the camera.
	Cols, Rows int
}
