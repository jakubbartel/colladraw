package main

import (
	"github.com/lucasb-eyer/go-colorful"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func generateColor() string {
	c := colorful.Hsv(rand.Float64()*360.0, 0.8, 0.8)
	return c.Hex()
}
