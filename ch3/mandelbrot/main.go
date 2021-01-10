// Mandelbrot emits a PNG image of the Mandelbrot fractal
package main

import (
	"image"
	"image/color"
	"image/color/palette"
	"image/png"
	"math/cmplx"
	"os"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -0.7, -0.7, -0.35, -0.35
		width, height          = 4096, 4096
	)

	var superPix [2][2]color.Color
	superPix[0][0] = color.White
	superPix[0][1] = color.White
	superPix[1][0] = color.White
	superPix[1][1] = color.White
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			superPix[px%2][py%2] = mandelbrot(z)
			// Image point (px,py) represents complex value z
			img.Set(px, py, sumSuperSample(superPix))
		}
		superPix[0][py%2] = color.Black
		superPix[1][py%2] = color.Black
	}
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

func mandelbrot(z complex128) color.Color {
	const iterations = 255
	const contrast = 8

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return palette.Plan9[255-contrast*n]
		}
	}
	return color.Black
}

func sumSuperSample(patch [2][2]color.Color) color.Color {
	var sumR, sumG, sumB, sumA uint32
	sumR, sumG, sumB, sumA = patch[0][0].RGBA()
	r, g, b, a := patch[0][1].RGBA()
	sumR += r
	sumG += g
	sumB += b
	sumA += a
	r, g, b, a = patch[1][0].RGBA()
	sumR += r
	sumG += g
	sumB += b
	sumA += a
	r, g, b, a = patch[1][1].RGBA()
	sumR += r
	sumG += g
	sumB += b
	sumA += a
	return color.RGBA{uint8(sumR / 4), uint8(sumG / 4), uint8(sumB / 4), uint8(sumA / 4)}
}
