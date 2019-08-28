// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 58.
//!+

// Package surface2 computes an SVG rendering of a 3-D surface function.
package surface2

import (
	"fmt"
	"io"
	"math"
	"os"
)

const (
	angle = float64(math.Pi / 6) // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

type Polygon struct {
	ax, ay, bx, by, cx, cy, dx, dy float64
	color                          string
}

func (p Polygon) String() string {
	return fmt.Sprintf(
		"<polygon fill='%s' points='%g,%g %g,%g %g,%g %g,%g'/>",
		p.color, p.ax, p.ay, p.bx, p.by, p.cx, p.cy, p.dx, p.dy,
	)
}

func (p Polygon) midPoint() (float64, float64) {
	return (p.ax + p.bx + p.cx + p.dx) / 4.0, (p.ay + p.by + p.cy + p.dy) / 4.0
}

// SurfacePlot writes an svg surface plot to w
func SurfacePlot(w io.Writer, width, height, cells int) {

	xyrange := 30.0                         // axis ranges (-xyrange..+xyrange)
	xyscale := float64(width) / 2 / xyrange // pixels per x or y unit
	zscale := float64(height) * 0.4         // pixels per z unit

	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	// polygons := make([]Polygon, 0)

	var zmax, zmin float64
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			_, _, z1 := getXYZ(i+1, j, cells, xyrange)
			_, _, z2 := getXYZ(i, j, cells, xyrange)
			_, _, z3 := getXYZ(i, j+1, cells, xyrange)
			_, _, z4 := getXYZ(i+1, j+1, cells, xyrange)
			for _, z := range [4]float64{z1, z2, z3, z4} {
				if z > zmax {
					zmax = z
				}
				if z < zmin {
					zmin = z
				}
			}
		}
	}

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			var poly Polygon
			poly.ax, poly.ay = corner(width, height, i+1, j, cells, xyrange, xyscale, zscale)
			poly.bx, poly.by = corner(width, height, i, j, cells, xyrange, xyscale, zscale)
			poly.cx, poly.cy = corner(width, height, i, j+1, cells, xyrange, xyscale, zscale)
			poly.dx, poly.dy = corner(width, height, i+1, j+1, cells, xyrange, xyscale, zscale)

			x, y := poly.midPoint()
			z := f(x, y)
			// mid1, mid2 := midPoint(ax, ay, bx, by)
			// mid3, mid4 := midPoint(cx, cy, dx, dy)
			// midX, midY := midPoint(mid1, mid2, mid3, mid4)
			// avgHeight := (midX + midY) / 2
			poly.color = getZColor(z, zmin, zmax)
			fmt.Fprintf(w, "%s\n", poly.String())
			// polygons = append(polygons, poly)
		}
	}

	fmt.Fprint(w, "</svg>")
	fmt.Fprintf(os.Stdout, "zmin: %v zmax: %v", zmin, zmax)
}

func getZColor(z, min, max float64) string {
	scale := max - min
	var green, blue uint8
	if z < 0 {
		green = toHex(z, scale)
	} else {
		blue = toHex(z, scale)
	}
	fmt.Fprintf(os.Stdout, "z: %v scale: %v, green: %02x, blue: %02x\n", z, scale, green, blue)
	return fmt.Sprintf("#%02x00%02x", green, blue)
}

func toHex(v, scale float64) uint8 {
	return uint8(math.Abs(v) * 255.0 / scale)
}

func midPoint(x1, y1, x2, y2 float64) (float64, float64) {
	return (x1 + x2) / 2.0, (y1 + y2) / 2.0
}

func getXYZ(i, j, cells int, rng float64) (x, y, z float64) {

	x = rng * (float64(i)/float64(cells) - 0.5)
	y = rng * (float64(j)/float64(cells) - 0.5)

	// Compute surface height z.
	z = f(x, y)
	if math.IsInf(z, 0) {
		z = 0
	}
	return
}

func corner(w, h, i, j, cells int, rng, xyscale, zscale float64) (float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x, y, z := getXYZ(i, j, cells, rng)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := float64(w)/2 + (x-y)*cos30*xyscale
	sy := float64(w)/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}

//!-
