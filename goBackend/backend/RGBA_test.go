package main

import (
	"fmt"
	"testing"
)

func TestRGBAImage(t *testing.T) {
	i, err := loadJpeg3("image.png")

	if comparator3(rgbaToGray3(i)) != 0 {
		t.Error("not correct gray conversion")
	}

	fmt.Println(err)
}
