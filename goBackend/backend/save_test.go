package main

import (
	"fmt"
	"testing"
)

func TestSaveImage(t *testing.T) {
	i, err := loadJpeg4("temporary.jpeg")

	saveImage4("nameTest.jpeg", i)

	if sameThing4("nameTest.jpeg", i) != 0 {
		t.Error("does not load the right image")
	}
	fmt.Println(err)
}
