// Package view supervises visible portions of the screen.
package view

import (
	"image"

	"github.com/mewmew/pgg/grid"
)

// A View is a visible portion of the screen.
type View struct {
	// The width and height of the view.
	Width, Height int
	// The width and height of the view in number of columns and rows
	// respectively.
	cols, rows int
	// The pixel offset between the top left point of the world and the view.
	off image.Point
	// The maximum valid pixel offset of the view.
	max image.Point
}

// NewView returns a new view of the specified dimensions. The top left point
// of the world is assumed to be located at (0, 0) and the bottom right point of
// the world is specified by end.
func NewView(width, height int, end image.Point) (v *View) {
	v = &View{
		Width:  width,
		Height: height,
		cols:   width / grid.CellWidth,
		rows:   height / grid.CellHeight,
		max:    end.Sub(image.Pt(width, height)),
	}
	return v
}

// Move moves the view based on the provided delta offset.
func (v *View) Move(delta image.Point) {
	off := v.off.Add(delta)
	// TODO(u): consider creating a geom.Clamp function to encapsulate this
	// behaviour. It would have the following function definition:
	//    func Clamp(p, min, max image.Point) image.Point
	if off.X < 0 {
		off.X = 0
	}
	if off.Y < 0 {
		off.Y = 0
	}
	if off.X > v.max.X {
		off.X = v.max.X
	}
	if off.Y > v.max.Y {
		off.Y = v.max.Y
	}
	v.off = off
}

// ColFirst returns the top left column visible through the view.
func (v *View) ColFirst() int {
	return v.off.X / grid.CellWidth
}

// RowFirst returns the top left row visible through the view.
func (v *View) RowFirst() int {
	return v.off.Y / grid.CellHeight
}

// ColLast returns the bottom right column visible through the view.
func (v *View) ColLast() int {
	cols := v.cols
	if v.off.X != 0 {
		// TODO(u): verify that views with a width of `n*grid.CellWidth + r`
		// don't cause an index overflow in the draw loop logic.
		cols++
	}
	return v.ColFirst() + cols
}

// RowLast returns the bottom right row visible through the view.
func (v *View) RowLast() int {
	rows := v.rows
	if v.off.Y != 0 {
		// TODO(u): verify that views with a height of `n*grid.CellHeight + r`
		// don't cause an index overflow in the draw loop logic.
		rows++
	}
	return v.RowFirst() + rows
}

// X returns the x offset to the grid columns visible through the view.
func (v *View) X() int {
	return v.off.X % grid.CellWidth
}

// Y returns the y offset to the grid rows visible through the view.
func (v *View) Y() int {
	return v.off.Y % grid.CellHeight
}
