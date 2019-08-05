package main

import (
	"fmt"
	"math/rand"
	"os"

	"github.com/ajstarks/deck/generate"
)

const (
	labelcolor = "grey"
	freecolor  = "maroon"
	linecolor  = "white"
	slavecolor = "black"
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

func hline(deck *generate.Deck, x, y, width, height float64, color string) {
	deck.Line(x, y, x+width, y, height, color)
}

func slavefree(deck *generate.Deck, pvalue, cvalue, dmax, x1, x2, y, incr float64, color string) {
	xp := make([]float64, 4)
	yp := make([]float64, 4)

	pct1 := 1.0 - (pvalue / dmax)
	pct2 := 1.0 - (cvalue / dmax)
	xp[0] = x1 + ((x2 - x1) * pct1)
	xp[1], xp[2] = x2, x2
	xp[3] = x1 + ((x2 - x1) * pct2)

	yp[0], yp[1] = y, y
	yp[2], yp[3] = y-incr+0.1, y-incr+0.1
	if cvalue == 3 {
		yp[3] = y - (incr / 2) + 0.1
	}
	deck.Polygon(xp, yp, color)
	deck.Line(xp[0], yp[0], xp[3], yp[3], 0.3, linecolor)
	deck.Line(x1, yp[2], xp[2], yp[2], 0.2, linecolor)

	if cvalue == 3 {
		xp[0] = x1
		xp[1] = x2
		xp[2] = x2
		xp[3] = x1
		yp[0] = y - (incr / 2) + 0.1
		yp[1] = y - (incr / 2) + 0.1
		yp[2] = y - incr + 0.1
		yp[3] = y - incr + 0.1
		deck.Polygon(xp, yp, color)
	}
}

func main() {
	deck := generate.NewSlides(os.Stdout, 0, 0)
	deck.StartDeck()
	deck.StartSlide()

	deck.TextMid(50, 95, "SLAVE AND FREE NEGROES.", "sans", 3, "black")
	deck.Rect(50, 45, 40, 80, slavecolor)

	top := 85.0
	y := top - 0.5
	data := []float64{1.3, 1.7, 1.7, 1.2, 0.8, 0.9, 0.7, 0.8, 3}

	deck.TextMid(75, 91, "PERCENT\nOF\nFREE NEGROES", "sans", 1.5, labelcolor)
	year := 1790
	for i := 0; i < len(data); i++ {
		deck.TextMid(25, y, fmt.Sprintf("%4d", year), "sans", 2, labelcolor)
		if i < 8 {
			label := fmt.Sprintf("%.1f", data[i])
			if i == 0 {
				label += "%"
			}
			deck.TextMid(75, y, label, "sans", 2, labelcolor)
		} else {
			deck.TextMid(75, y, "100%", "sans", 2, labelcolor)
		}
		y -= 10.0
		year += 10
	}

	y = top
	yincr := 10.0
	for i := 1; i < len(data); i++ {
		slavefree(deck, data[i-1], data[i], 3, 30, 70, y, yincr, "red")
		y -= yincr
	}
	vdeckle(deck, 30, 5, -1.5, 80, 250, slavecolor)
	deck.EndSlide()
	deck.EndDeck()
}
