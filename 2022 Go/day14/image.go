package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
)

func MapIMG(bounds image.Rectangle, lines [][]image.Point) *image.RGBA {
	img := image.NewRGBA(bounds)
	// Set image to black
	for i := bounds.Min.Y; i < bounds.Max.Y; i++ {
		for j := bounds.Min.X; j < bounds.Max.X; j++ {
			img.Set(j, i, color.Black)
		}
	}

	// Set lines to white
	for _, l := range lines {
		for t := 1; t < len(l); t++ {
			i := l[t-1].Y
			for ; i != l[t].Y; i += sig(l[t].Y - l[t-1].Y) {
				img.Set(l[t].X, i, color.White)
			}
			img.Set(l[t].X, i, color.White)
			j := l[t-1].X
			for ; j != l[t].X; j += sig(l[t].X - l[t-1].X) {
				img.Set(j, l[t].Y, color.White)
			}
			img.Set(j, l[t].Y, color.White)
		}
	}

	// Set start position's color
	img.Set(500, 0, color.White)
	return img
}

func PrintIMG(img *image.RGBA, name string) {
	f, err := os.Create(name + ".png")
	if err != nil {
		log.Fatal(err)
	}
	err = png.Encode(f, img)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println("img written")
}

func sig(val int) int {
	if val < 0 {
		return -1
	} else {
		return 1
	}
}
