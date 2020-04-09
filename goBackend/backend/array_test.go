package main

import (
	"fmt"
	"testing"
)

func TestArrayImage(t *testing.T) {
	i, err := loadJpeg1("image.png")

	if comparator1(arrayToGray1(rgbaToGray1(i))) != 0 {
		t.Error("not correct gray conversion")
	}

	fmt.Println(err)
}
