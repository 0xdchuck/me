package main

import (
	"math/rand"
	"time"

	"github.com/fogleman/gg"
)

const (
	width  = 400
	height = 400
	cube   = 4
)

func main() {
	c := gg.NewContext(width, height)
	generatePattern(c)
	generateBody(c)
	generateFace(c)
	generateEyes(c)
	generateNose(c)
	generateMouth(c)
	generateHair(c)
	err := c.SavePNG("me.png")
	if err != nil {
		panic(err)
	}
}

func generatePattern(c *gg.Context) {
	source := rand.NewSource(time.Now().Unix())
	for w := 0; w <= width; w += cube {
		for h := 0; h <= height; h += cube {
			c.DrawRectangle(float64(w), float64(h), cube, cube)
			r := (w * h)
			r1 := float64(r+rand.New(source).Intn(30)) / 255
			r2 := float64(r+rand.New(source).Intn(20)) / 255
			r3 := float64(0+rand.New(source).Intn(50)) / 255
			c.SetRGB(r1, r2, r3)
			c.Fill()

		}
	}
	c.ClosePath()
}

func generateHair(c *gg.Context) {
	source := rand.NewSource(42)
	for x := 144; x <= (144 + 110); x += cube {
		for y := 120; y <= (120 + 30); y += cube {
			c.DrawRectangle(float64(x), float64(y), cube, cube)
			r := float64(rand.New(source).Intn(20))
			c.SetRGBA(
				float64((r+153.0)/255.0),
				float64((r+76.0)/255.0),
				0,
				1,
			)
			c.Fill()
			c.ClosePath()
		}
	}
}

func generateFace(c *gg.Context) {
	c.DrawRectangle(145, 130, 110, 140)
	c.SetRGB(1, 0.886, 0.7411)
	c.Fill()
}

func generateNose(c *gg.Context) {
	c.DrawRectangle(195, 180, 15, 40)
	c.SetRGB(1, 0.8, 0.6)
	c.Fill()

	c.DrawRectangle(182, 220, 40, 15)
	c.SetRGB(1, 0.8, 0.6)
	c.Fill()
}

func generateEyes(c *gg.Context) {
	c.DrawRectangle(220, 160, 20, 30)
	c.SetRGB(0, 0, 0)
	c.Fill()

	c.DrawRectangle(221, 162, 18, 27)
	c.SetRGB(0, 0, 0)
	c.Fill()

	c.DrawRectangle(224, 164, 4, 4)
	c.SetRGB(1, 1, 1)
	c.Fill()

	c.DrawRectangle(228, 168, 4, 4)
	c.SetRGB(1, 1, 1)
	c.Fill()

	c.DrawRectangle(232, 172, 4, 4)
	c.SetRGB(1, 1, 1)
	c.Fill()

	c.DrawRectangle(160, 160, 20, 30)
	c.SetRGB(0, 0, 0)
	c.Fill()

	c.DrawRectangle(161, 162, 18, 27)
	c.SetRGB(0, 0, 0)
	c.Fill()

	c.DrawRectangle(164, 164, 4, 4)
	c.SetRGB(1, 1, 1)
	c.Fill()

	c.DrawRectangle(168, 168, 4, 4)
	c.SetRGB(1, 1, 1)
	c.Fill()

	c.DrawRectangle(172, 172, 4, 4)
	c.SetRGB(1, 1, 1)
	c.Fill()

	c.DrawRectangle(180, 170, 40, 2)
	c.SetRGB(0, 0, 0)
	c.Fill()
}

func generateMouth(c *gg.Context) {
	c.DrawRectangle(190, 250, 20, 5)
	c.SetRGB(1, 0.79, 0.7)
	c.Fill()
}

func generateBody(c *gg.Context) {
	c.DrawRectangle(120, 250, 160, 220)
	c.SetRGB(0.1, 0.1, 0.1)
	c.Fill()
}
