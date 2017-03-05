package main

/*
练习 10.1：
 扩展jpeg程序，以支持任意图像格式之间的相互转换，使用image.Decode检测支持的格式类型，然后通过flag命令行标志参数选择输出的格式。
*/
import (
	"flag"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"

	"io"
	"os"
)

func main() {
	var format string
	flag.StringVar(&format, "fmt", "png", "convert to image format")
	flag.Parse()

	if err := convertFormat(os.Stdin, os.Stdout, format); err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}
}

func convertFormat(in io.Reader, out io.Writer, format string) error {
	img, kind, err := image.Decode(in)
	if err != nil {
		return err
	}
	fmt.Println("current format ", kind)
	switch format {
	case "png":
		return png.Encode(out, img)
	case "jpeg":
		return jpeg.Encode(out, img, &jpeg.Options{Quality: 95})
	default:
		return fmt.Errorf("Format is not vaild")
	}
}
