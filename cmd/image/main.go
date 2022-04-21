package main

import (
	"flag"
	"image"
	"image/color"
	"image/png"
	"os"
)

func main() {
	var width, height int
	flag.IntVar(&width, "w", 200, "-w=200")
	flag.IntVar(&height, "h", 200, "-w=200")
	flag.Parse()
	img := image.NewRGBA(
		image.Rectangle{
			image.Point{0, 0},
			image.Point{width, height},
		},
	)
	cyan := color.RGBA{100, 200, 200, 0xff}

	for w := 0; w < width; w++ {
		for h := 0; h < height; h++ {
			x := float64(w)/float64(width)*2. - 1.
			y := float64(h)/float64(height)*2. - 1.
			if float64(x*x+y*y) < 0.5 {
				img.Set(w, h, cyan)
			}
		}
	}

	// Encode as PNG.
	f, _ := os.Create("/tmp/image.png")
	png.Encode(f, img)
}
