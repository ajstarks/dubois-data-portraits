package main

import (
	"flag"
	"fmt"
	"io"
	"math"
	"os"
)

func text(w io.Writer, x, y, fs float64, s, font, color string) {
	fmt.Fprintf(w, "<text xp=\"%.2f\" yp=\"%.2f\" sp=\"%.3f\"  font=%q color=%q>%s</text>\n", x, y, fs, font, color, s)
}

func rtext(w io.Writer, x, y, angle, fs float64, s, font, color string) {
	fmt.Fprintf(w, "<text xp=\"%.2f\" yp=\"%.2f\" sp=\"%.3f\" rotation=\"%.2f\" font=%q color=%q>%s</text>\n", x, y, fs, angle, font, color, s)
}

func drtext(w io.Writer, x, y, angle, fs float64, s, font, color string) {
	fmt.Fprintf(w, "rtext %q %.2f %.2f %.2f %.2f %q %q\n", s, x, y, angle, fs, font, color)
}

func beginDeck(w io.Writer) {
	fmt.Fprintln(w, "<deck>")
}

func beginSlide(w io.Writer) {
	fmt.Fprintln(w, "<slide>")
}

func endSlide(w io.Writer) {
	fmt.Fprintln(w, "</slide>")
}

func endDeck(w io.Writer) {
	fmt.Fprintln(w, "</deck>")
}

func toc(w io.Writer, s, font string, cx, cy, r, begin, end, textsize, cw, ch float64) {
	angle := begin
	interval := (end - begin) / float64(len(s)-1)
	for _, t := range s {
		px, py := polar(cx, cy, r, angle, cw, ch)
		textangle := angle + 90
		drtext(w, px, py, textangle, textsize, string(t), font, "")
		angle += interval
	}
}

func rotext(s string, begin, end, interval float64) {
	for angle := begin; angle < end; angle += interval {
		rtext(os.Stdout, 50, 50, angle, 8, s, "sans", "black")
	}
}

// polar to Cartesian coordinates, corrected for aspect ratio
func polar(cx, cy, r, theta, cw, ch float64) (float64, float64) {
	ry := r * (cw / ch)
	t := theta * (math.Pi / 180)
	return cx + (r * math.Cos(t)), cy + (ry * math.Sin(t))
}

func main() {
	var cx, cy, beginAngle, endAngle, radius, textsize, cw, ch float64
	var font string
	flag.Float64Var(&cx, "cx", 50, "begin angle")
	flag.Float64Var(&cy, "cy", 50, "begin angle")
	flag.Float64Var(&beginAngle, "beg", 180, "begin angle")
	flag.Float64Var(&endAngle, "end", 360, "angle interval")
	flag.Float64Var(&radius, "radius", 30, "radius")
	flag.Float64Var(&textsize, "ts", 2, "text size")
	flag.Float64Var(&cw, "cw", 792, "canvas width")
	flag.Float64Var(&ch, "ch", 612, "canvas height")
	flag.StringVar(&font, "font", "mono", "font")
	flag.Parse()

	a := flag.Args()
	if len(a) > 0 {
		//beginDeck(os.Stdout)
		//beginSlide(os.Stdout)
		toc(os.Stdout, a[0], font, cx, cy, radius, beginAngle, endAngle, textsize, cw, ch)
		//endSlide(os.Stdout)
		//endDeck(os.Stdout)
	}
}
