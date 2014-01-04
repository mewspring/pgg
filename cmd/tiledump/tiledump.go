/*
tiledump it a tool which extracts tile images contained within tile sets.

Usage:

	tiledump [OPTION]... IMG...

Flags:

	-w (default=32)
		Tile width.
	-h (default=32)
		Tile height.

Examples:

	tiledump -w 64 -h 64 tileset.png
*/
package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/mewkiz/pkg/imgutil"
	"github.com/mewkiz/pkg/pathutil"
	"github.com/mewmew/pgg/tileset"
)

// The tile width and height are specifiable from command line.
var tileWidth, tileHeight int

func init() {
	flag.IntVar(&tileWidth, "w", 32, "Tile width.")
	flag.IntVar(&tileHeight, "h", 32, "Tile height.")
	flag.Usage = usage
}

func usage() {
	fmt.Fprintln(os.Stderr, "Usage: tiledump [OPTION]... IMG...")
	fmt.Fprintln(os.Stderr)
	fmt.Fprintln(os.Stderr, "Flags:")
	flag.PrintDefaults()
	fmt.Fprintln(os.Stderr)
	fmt.Fprintln(os.Stderr, "Examples:")
	fmt.Fprintln(os.Stderr, "  tiledump -w 64 -h 64 tileset.png")
}

func main() {
	flag.Parse()
	for _, imgPath := range flag.Args() {
		err := tiledump(imgPath)
		if err != nil {
			log.Fatalln(err)
		}
	}
}

// tiledump extracts the tile images contained within the provided tile set.
func tiledump(imgPath string) (err error) {
	ts, err := tileset.Open(imgPath, tileWidth, tileHeight)
	if err != nil {
		return err
	}
	tileDir := pathutil.TrimExt(imgPath)
	err = os.Mkdir(tileDir, 0755)
	if err != nil {
		return err
	}
	for id := tileset.TileID(1); id <= ts.LastID(); id++ {
		tile := ts.Tile(id)
		tilePath := fmt.Sprintf("%s/tile_%04d.png", tileDir, id)
		err = imgutil.WriteFile(tilePath, tile)
		if err != nil {
			return err
		}
	}
	return nil
}
