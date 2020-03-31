package main

import (
	"fmt"
   	"testing"
)


func TestLoadImage(t *testing.T) {
	i, err := loadJpeg("temporary.jpeg")
	
	saveImage("nameTest.jpeg", i)

	if sameThing("nameTest.jpeg", i) != 0 {
		t.Error("does not load the right image")
	}
	fmt.Println(err)
}