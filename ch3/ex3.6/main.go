// Mandelbrot emits a PNG image of the Mandelbrot fractal.
// 升采样技术可以降低每个像素对计算颜色值和平均值的影响。简单的方法是将每个像素分成四个子像素，实现它。
package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)


func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py+1, mandelbrot(z))
			img.Set(px+1, py, mandelbrot(z))
			img.Set(px+1, py+1, mandelbrot(z))
			img.Set(px, py, mandelbrot(z))

		}
	}
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.RGBA{(255 - contrast*n / 2), 255 - contrast*n, 0, 255}
		}
	}
	return color.RGBA{0, 0, 100, 255}
}