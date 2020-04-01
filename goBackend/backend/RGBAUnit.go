package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
)

func comparator3(img image.Image) float64 {

	bounds := img.Bounds()

	gray := image.NewGray(bounds)
	cyan := color.RGBA{100, 200, 200, 0xff}
	gray.Set(1, 1, cyan)
	cian := gray.At(1, 1)

	for x := 0; x < bounds.Max.X; x++ {
		for y := 0; y < bounds.Max.Y; y++ {
			switch {
			case x < bounds.Max.X/2 && y < bounds.Max.Y/2: // upper left quadrant

				if img.At(x, y) != cian {

					picks := img.At(x, y)
					fmt.Println(picks)
					fmt.Println(cian)
					return 1
					os.Exit(1)
				}

			default:
				// Use zero value.
			}
		}
	}
	return 0
}

func rgbaToGray3(img image.Image) image.Image {
	var (
		bounds = img.Bounds()
		gray   = image.NewGray(bounds)
		rect   = make([][]int, bounds.Max.X)
	)
	for i := 0; i < len(rect); i++ {
		rect[i] = make([]int, bounds.Max.Y)
	}
	for x := 0; x < bounds.Max.X; x++ {
		for y := 0; y < bounds.Max.Y; y++ {
			var rgba = img.At(x, y)
			gray.Set(x, y, rgba)
			rect[x][y] = int(gray.GrayAt(x, y).Y)
		}
	}
	f, _ := os.Create("grayImage.png")
	png.Encode(f, gray)
	return gray

}

func loadJpeg3(filepath string) (image.Image, error) {
	infile, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer infile.Close()
	img, _, err := image.Decode(infile)
	if err != nil {
		return nil, err
	}
	return img, nil
}
