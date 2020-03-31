package main

import (
	"fmt"
   	"testing"
)


func TestLoadImage(t *testing.T) {
	i, err := loadJpeg("image.png")

	if comparator(arrayToGray(rgbaToGray(i))) != 0 {
		t.Error("not correct gray conversion")
	}

	fmt.Println(err)
}