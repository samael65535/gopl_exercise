// Mandelbrot emits a PNG image of the Mandelbrot fractal.
package main

/*
练习 8.5：
使用一个已有的CPU绑定的顺序程序，比如在3.3节中我们写的Mandelbrot程序或者3.2节中的3-D surface计算程序，
并将他们的主循环改为并发形式，使用channel来进行通信
*/

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
	"sync"
)

type Cell struct {
	py int
	y  float64
}

const (
	xmin, ymin, xmax, ymax = -2, -2, +2, +2
	width, height          = 1024, 1024
)

var limit chan Cell
var wg sync.WaitGroup

func main() {
	limit = make(chan Cell, height)
	wg = sync.WaitGroup{}
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	for py := 0; py < height; py++ {
		limit <- Cell{py, float64(py)/height*(ymax-ymin) + ymin}
		wg.Add(1)
		go handleImg(img)
	}

	wg.Wait()
	png.Encode(os.Stdout, img)
	_ = os.Stdout
	_ = png.Encode
	_ = fmt.Print
}

func handleImg(img *image.RGBA) {
	cell := <-limit
	py := cell.py
	y := cell.y
	for px := 0; px < width; px++ {
		x := float64(px)/width*(xmax-xmin) + xmin
		z := complex(x, y)
		img.Set(px, py, mandelbrot(z))
	}
	wg.Done()
}

func mandelbrot(z complex128) color.Color {
	const iterations = 255
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}
