package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"log"
	"os"
)

func saveImage4(filename string, img image.Image) {
	out, err := os.Create(filename)
	err = jpeg.Encode(out, img, nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func loadJpeg4(filepath string) (image.Image, error) {
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

func diff4(a, b uint32) int64 {
	if a > b {
		return int64(a - b)
	}
	return int64(b - a)
}

func sameThing4(filepath string, i1 image.Image) float64 {

	i2, err := loadJpeg4(filepath)
	if err != nil {
		log.Fatal(err)
	}

	if i1.ColorModel() != i2.ColorModel() {
		log.Fatal("different color models")
	}

	b := i1.Bounds()
	if !b.Eq(i2.Bounds()) {
		log.Fatal("different image sizes")
	}

	var sum int64
	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			r1, g1, b1, _ := i1.At(x, y).RGBA()
			r2, g2, b2, _ := i2.At(x, y).RGBA()
			sum += diff4(r1, r2)
			sum += diff4(g1, g2)
			sum += diff4(b1, b2)
		}
	}

	nPixels := (b.Max.X - b.Min.X) * (b.Max.Y - b.Min.Y)

	return float64(sum*100) / (float64(nPixels) * 0xffff * 3)

}
