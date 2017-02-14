// Mandelbrot emits a PNG image of the Mandelbrot fractal.
//  另一个生成分形图像的方式是使用牛顿法来求解一个复数方程，例如z^4-1=0, 每个起点到四个根的迭代次数对应阴影的灰度。方程根对应的点用颜色表示。
package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
//	"io/ioutil"
	//	"fmt"
	"os"
)


func main() {
	const (
		xmin, ymin, xmax, ymax = -1, -1, +1, +1
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

func mandelbrot(v complex128) color.Color {
	const iterations = 200
	const contrast = 15

	for n := uint8(0); n < iterations; n++ {
		v = v - f(v) / fd(v)
		if cmplx.Abs(v) == 1 {
			return color.Gray{255 - contrast * n}
		}

	}
	return color.Black
}

func f(z complex128) complex128 {
	return cmplx.Pow(z, 4) - 1
}

func fd(z complex128) complex128 {
	return 4 * cmplx.Pow(z, 3)
}
