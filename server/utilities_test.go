package main

import (
	"testing"
)

func TestGenerateColor(t *testing.T) {
	color := generateColor()
	if len(color) != 7 {
		t.Error("Color code is not 6 characters long")
	}
	if color[0:1] != "#" {
		t.Error("Color first character has to be be \"#\"")
	}
}
