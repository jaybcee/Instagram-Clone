package main

import (
	"fmt"
	"testing"
)

func TestLoadImage(t *testing.T) {

	i, err := loadJpeg2("temporary.png")
	fmt.Println(err)
	if sameThing2("temporary.png", i) != 0 {
		t.Error("Expected 2 + 2 to equal 4")
	}
}
