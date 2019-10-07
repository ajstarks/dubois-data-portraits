package main

import (
	"flag"
	"math/rand"
	"os"

	"github.com/ajstarks/deck/generate"
)

// hdeckle makes a horizontal deckled edge running from (x1, y1) to (x2, y2)
func hdeckle(deck *generate.Deck, x, y, width, height float64, n int, color string) {
	xp := make([]float64, n)
	yp := make([]float64, n)

	// left point
	xp[0] = x
	yp[0] = y

	// right point
	xp[n-2] = x + width
	yp[n-2] = y

	// back to the left
	xp[n-1] = x
	yp[n-1] = y

	xincr := width / float64(n)
	for i := 1; i <= n-3; i++ {
		xp[i] = xp[i-1] + xincr
		yp[i] = y + rand.Float64()*height
	}
	deck.Polygon(xp, yp, color)
}

func vdeckle(deck *generate.Deck, x, y, width, height float64, n int, color string) {
	xp := make([]float64, n)
	yp := make([]float64, n)

	// bottom point
	xp[0] = x
	yp[0] = y

	// top point
	xp[n-2] = x
	yp[n-2] = y + height

	// back to the bottom
	xp[n-1] = x
	yp[n-1] = y

	yincr := height / float64(n)
	for i := 1; i <= n-3; i++ {
		yp[i] = yp[i-1] + yincr
		xp[i] = x + rand.Float64()*width
	}
	deck.Polygon(xp, yp, color)
}

func main() {
	deck := generate.NewSlides(os.Stdout, 0, 0)
	var x, y, width, height float64
	var n int
	var color, dtype string

	flag.Float64Var(&x, "x", 50, "x")
	flag.Float64Var(&y, "y", 50, "y")
	flag.Float64Var(&width, "w", 20, "width")
	flag.Float64Var(&height, "h", 20, "height")
	flag.IntVar(&n, "n", 250, "number of bumps")
	flag.StringVar(&color, "color", "gray", "color")
	flag.StringVar(&dtype, "type", "h", "v for vertical, h for horizontal")
	flag.Parse()

	switch dtype {
	case "v":
		vdeckle(deck, x, y, width, height, n, color)
	case "h":
		vdeckle(deck, x, y, width, height, n, color)
	}
}
