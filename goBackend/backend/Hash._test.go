package main

import "testing"

func TestSum(t *testing.T) {
	h, _ := Bhash("password")
	success := CheckHash("password", h)
	if !success {
		t.Errorf("Hash Error")
	}
}
