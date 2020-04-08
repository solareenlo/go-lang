package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/color"
	"image/png"
)

type Image struct {
	w int
	h int
}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.w, i.h)
}

func (Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (Image) At(x, y int) color.Color {
	v := x + y
	return color.RGBA{uint8(v), uint8(v), 255, 255}
}

func main() {
	m := Image{200, 200}
	ShowImage(m)
}

func Show(f func(int, int) [][]uint8) {
	const (
		dx = 256
		dy = 256
	)
	data := f(dx, dy)
	m := image.NewNRGBA(image.Rect(0, 0, dx, dy))
	for y := 0; y < dy; y++ {
		for x := 0; x < dx; x++ {
			v := data[y][x]
			i := y*m.Stride + x*4
			m.Pix[i] = v
			m.Pix[i+1] = v
			m.Pix[i+2] = 255
			m.Pix[i+3] = 255
		}
	}
	ShowImage(m)
}

func ShowImage(m image.Image) {
	var buf bytes.Buffer
	err := png.Encode(&buf, m)
	if err != nil {
		panic(err)
	}
	enc := base64.StdEncoding.EncodeToString(buf.Bytes())
	fmt.Println("IMAGE:" + enc)
}
