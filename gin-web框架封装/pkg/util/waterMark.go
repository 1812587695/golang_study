package util

import (
	"image"
	"os"
	"image/jpeg"
	"golang.org/x/image/font"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/math/fixed"
	"bufio"
	"image/png"
	"github.com/golang/freetype"
	"io/ioutil"
	"image/draw"
)

var (
	dpi = float64(102)
	fontFile = "pkg/fonts/msyhbd.ttc"
	size = float64(14)
)


func WaterMark(source string, out string, text string) error {
	fontSourceBytes, _ := ioutil.ReadFile(fontFile)
	f, err := freetype.ParseFont(fontSourceBytes)
	if err != nil{
		return err
	}
	//字体背景
	fg := image.White

	//原始图片
	imgb, _ := os.Open(source)
	img, err := jpeg.Decode(imgb)
	if err != nil {
		return err
	}

	imgWidth, imgHeight := img.Bounds().Dx(), img.Bounds().Dy()

	rgba := image.NewNRGBA(image.Rect(0, 0, imgWidth, imgHeight))

	defer imgb.Close()

	draw.Draw(rgba, rgba.Bounds(), img, image.ZP, draw.Src)

	h := font.HintingNone
	d := &font.Drawer{
		Dst: rgba,
		Src:fg,
		Face:truetype.NewFace(f, &truetype.Options{
			Size: size,
			DPI:dpi,
			Hinting:h,
		}),
	}
	y := imgHeight - 10
	d.Dot = fixed.Point26_6{
		X:fixed.I(imgWidth - 10) - d.MeasureString(text),
		Y:fixed.I(y),
	}

	d.DrawString(text)

	outFile, err := os.Create(out)
	if err != nil {
		return err
	}
	defer outFile.Close()
	b := bufio.NewWriter(outFile)
	err = png.Encode(b, rgba)
	if err != nil {
		return err
	}
	err = b.Flush()
	if err != nil {
		return err
	}
	return nil
}