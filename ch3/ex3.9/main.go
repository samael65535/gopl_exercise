package main
// 编写一个web服务器，用于给客户端生成分形的图像。运行客户端用过HTTP参数参数指定x,y和zoom参数。
import (
	"net/http"
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"strconv"
	"math"
	"fmt"
)

const (
	width, height          = 1024, 1024
)

func main() {
	http.HandleFunc("/", handle)
	http.ListenAndServe("localhost:8001", nil)
}

func handle(w http.ResponseWriter, r *http.Request) {
	params := map[string] float64{
		"x": 2,
		"y": 2,
		"zoom": 2,
	}
	for name := range params {
		v := r.FormValue(name)
		if v == "" {
			continue
		}
		f, err := strconv.ParseFloat(v,64)
		if err != nil {
			http.Error(w, fmt.Sprintf("%s be must float64", name), http.StatusBadRequest);
			return
		}
		params[name] = math.Abs(f)
	}
	x := params["x"] * params["zoom"] / 2
	y := params["y"] * params["zoom"] / 2

	xmin, ymin, xmax, ymax := -x, -y, x, y

	fmt.Println(xmin, ymin, xmax, ymax)
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
	png.Encode(w, img) // NOTE: ignoring errors
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
