package surface

import (
	"fmt"
	"io"
	"math"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid in cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
	blue          = "#0000ff"           // blue color code
	red           = "#ff0000"           // red color code
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func Render(w io.Writer, f func(float64, float64) float64) {
	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>\n", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, _ := corner(i+1, j, f)
			bx, by, color := corner(i, j, f)
			cx, cy, _ := corner(i, j+1, f)
			dx, dy, _ := corner(i+1, j+1, f)
			if !checkFinite(ax, ay, bx, by, cx, cy, dx, dy) {
				continue
			}
			fmt.Fprintf(w, "<polygon points='%g,%g, %g,%g, %g,%g, %g,%g'", ax, ay, bx, by, cx, cy, dx, dy)
			fmt.Fprintf(w, " stroke='%s'/>\n", color)
		}
	}
	fmt.Fprintln(w, "</svg>")
}

func checkFinite(args ...float64) bool {
	for _, n := range args {
		if math.IsInf(n, 0) || math.IsNaN(n) {
			return false
		}
	}
	return true
}

func corner(i, j int, f func(float64, float64) float64) (sx, sy float64, color string) {
	// Find point (x, y) at corner of cell (i, j)
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Complete surface height z.
	z := f(x, y)

	if z < 0 {
		color = blue
	} else {
		color = red
	}

	// Project (x, y, z) isometrically onto 2-D svg canvas (sx, sy).
	sx = width/2 + (x-y)*cos30*xyscale
	sy = height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, color
}

func Sinr(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0, 0)
	return math.Sin(r) / r
}

func EggBox(x, y float64) float64 {
	a := .1
	b := 1.
	return a * (math.Sin(x/b) + math.Sin(y/b))
}

func Saddle(x, y float64) float64 {
	return .0015 * (x*x - y*y)
}

func Moguls(x, y float64) float64 {
	z := y*math.Sin(x) - x*math.Cos(y)
	return z / 100
}
