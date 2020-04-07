package main

import (
	"fmt"
	"testing"
)

func TestSaveImage(t *testing.T) {
	i, err := loadJpeg4("temporary.jpeg")

	saveImage4("nameTest.jpeg", i)

	if sameThing4("nameTest.jpeg", i) > 1 {
		fmt.Println(sameThing4("nameTest.jpeg", i))
		t.Error("image quality difference exceeds 1 percent")
	}
	fmt.Println(err)
}
