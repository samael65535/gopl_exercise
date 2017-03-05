package main

/*
练习9.6:
测试一下计算密集型的并发程序(练习8.5那样的)会被GOMAXPROCS怎样影响到。在你的电脑上最佳的值是多少？你的电脑CPU有多少个核心？
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
