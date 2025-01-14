package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"io"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/gomono"
	"golang.org/x/image/math/fixed"
)

type box struct {
	top, right, bottom, left int
}

type drawOptions struct {
	padding   box // absolute padding of the image
	width     int // width of the image
	ttype     truetype.Options
	tpad      int // padding for the line height
	fontColor color.NRGBA
}

func drawText(w io.Writer, lines []string, opt *drawOptions) error {
	if opt == nil {
		opt = &drawOptions{
			padding:   box{top: 50, bottom: 50, left: 20},
			width:     *flagImgWidth,
			ttype:     truetype.Options{Size: 12, DPI: 300},
			tpad:      10,
			fontColor: color.NRGBA{0, 0, 0, 255},
		}
	}
	lineHeight := calcLineHeight(opt.ttype) + opt.tpad
	imgHeight := lineHeight*len(lines) + opt.padding.top + opt.padding.bottom
	rect := image.Rect(0, 0, opt.width, imgHeight)
	img := image.NewNRGBA(rect)
	draw.Draw(img, rect, image.White, image.ZP, draw.Src)
	py := lineHeight + opt.padding.top
	for _, l := range lines {
		addText(img, opt.padding.left, py, l, *opt)
		py += lineHeight
	}
	if err := png.Encode(w, img); err != nil {
		return fmt.Errorf("drawText/Encode: %w", err)
	}
	return nil
}

func calcLineHeight(opt truetype.Options) int {
	return int(opt.DPI / 72 * opt.Size)
}

func addText(img *image.NRGBA, x, y int, text string, opt drawOptions) {
	point := fixed.Point26_6{fixed.Int26_6(x * 64), fixed.Int26_6(y * 64)}

	d := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(opt.fontColor),
		Face: goFontFace(&opt.ttype),
		Dot:  point,
	}
	d.DrawString(text)
}

func goFontFace(opt *truetype.Options) font.Face {
	f, err := truetype.Parse(gomono.TTF)
	if err != nil {
		panic(fmt.Sprint("cannot parse font:", err))
	}
	nf := truetype.NewFace(f, opt)
	return nf
}
