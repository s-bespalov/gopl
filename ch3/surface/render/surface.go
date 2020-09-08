// Surface computes an SVG rendering of a 3-D surface function
package render

import (
	"fmt"
	"math"
	"strings"
)

const (
	cells   = 100         // number of grid cells
	xyrange = 30.0        // axis ranges (-xyrange..+xyrange)
	angle   = math.Pi / 6 // angle of x, y axes (=30°)
	//blue          = "#0000ff"   // blue color code
	//red           = "#ff0000"   // red color code
)

var width, height int = 600, 320 // canvas size in pixels
var xyscale float64              // pixels per x or y unit
var zscale float64               // pixels per z unit

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

// Surf returns a rendered 3-D surface in SVG format
func Surf(iheight int, iwidth int, pcolor string, vcolor string) string {
	width, height = iwidth, iheight
	xyscale = float64(width) / 2 / xyrange
	zscale = float64(height) * 0.4
	var b strings.Builder
	b.WriteString(fmt.Sprintf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height))
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, z1, ok0 := corner(i+1, j)
			bx, by, z2, ok1 := corner(i, j)
			cx, cy, z3, ok2 := corner(i, j+1)
			dx, dy, z4, ok3 := corner(i+1, j+1)
			color := vcolor
			if (z1+z2+z3+z4)/4 > 0. {
				color = pcolor
			}
			if ok0 && ok1 && ok2 && ok3 {
				b.WriteString(fmt.Sprintf("<polygon points='%g,%g,%g,%g,%g,%g,%g,%g' stroke='%s'/>\n",
					ax, ay, bx, by, cx, cy, dx, dy, color))
			}
		}
	}
	b.WriteString("</svg>")
	return b.String()
}

func corner(i, j int) (float64, float64, float64, bool) {
	// Find point (x,y) at corner of cell (i,j)
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)
	if math.IsNaN(z) || math.IsInf(z, 0) {
		return 0, 0, 0, false
	}

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx, sy)
	sx := float64(width)/2 + (x-y)*cos30*xyscale
	sy := float64(height)/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, z, true
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	sinplus := math.Sin(x + y)
	sinminus := math.Sin(x - y)
	return (sinplus * sinminus) / r
}
