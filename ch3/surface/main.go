// Surface computes an SVG rendering of a 3-D surface function
package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
)

const (
	width, height = 800, 520
	cells         = 150
	xyrange       = 30.0
	xyscale       = width / 2 / xyrange
	zscale        = height * 0.2
	angle         = math.Pi / 6
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/svg+xml")
		fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
			"style='stroke: grey; fill: white; stroke-width: 0.7' "+
			"width='%d' height='%d'>", width, height)
		for i := 0; i < cells; i++ {
			for j := 0; j < cells; j++ {
				ax, ay, aErr, aColor := corner(i+1, j)
				bx, by, bErr, bColor := corner(i, j)
				cx, cy, cErr, cColor := corner(i, j+1)
				dx, dy, dErr, dColor := corner(i+1, j+1)

				if aErr || bErr || cErr || dErr {
					continue
				}

				colorVal := uint32((math.Abs(aColor+bColor+cColor+dColor) / 4) * 0xffffff)

				fmt.Fprintf(w, "<polygon style='fill: #%06x' points='%g,%g %g,%g %g,%g %g,%g'/>\n", colorVal, ax, ay, bx, by, cx, cy, dx, dy)
			}
		}
		fmt.Fprintf(w, "</svg>")
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func corner(i, j int) (float64, float64, bool, float64) {
	// Find point (x,y) at corner of cell (i,j)
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// compute surface height z
	z := f(x, y)

	colorVal := z

	if math.IsNaN(z) {
		return 0.0, 0.0, true, colorVal
	}

	// project (x,y,z) isometrically onto 2-D SVG canvas (sx, sy)
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, false, colorVal
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Cos(r)
}
