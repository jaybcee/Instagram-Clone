package main

import (
	"fmt"
	"testing"
)

func TestLoadImage(t *testing.T) {

	i, err := loadJpeg("temporary.png")
	fmt.Println(err)
	if sameThing("temporary.png", i) != 0 {
		t.Error("Expected 2 + 2 to equal 4")
	}
}