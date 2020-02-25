// This file has automatically been generated on Wed Feb 26 02:10:00 +05 2020.
// DO NOT EDIT.
package color

import (
	"image/color"
	_ "unsafe"
)

//go:linkname alpha16rgba color.sub_alpha16rgba
func alpha16rgba(c color.Alpha16) uint32 {
	return c.RGBA()
}

//go:linkname Alpha16RGBA color.sub_alpha16rgba
//go:noescape
func Alpha16RGBA(c color.Alpha16) uint32

//go:linkname nycbcrargba color.sub_nycbcrargba
func nycbcrargba(c color.NYCbCrA) (uint32, uint32, uint32, uint32) {
	return c.RGBA()
}

//go:linkname NYCbCrARGBA color.sub_nycbcrargba
//go:noescape
func NYCbCrARGBA(c color.NYCbCrA) (uint32, uint32, uint32, uint32)

//go:linkname rgbargba color.sub_rgbargba
func rgbargba(c color.RGBA) uint32 {
	return c.RGBA()
}

//go:linkname RGBARGBA color.sub_rgbargba
//go:noescape
func RGBARGBA(c color.RGBA) uint32

//go:linkname CMYKToRGB image/color.CMYKToRGB
//go:noescape
func CMYKToRGB(c, m, y, k uint8) (uint8, uint8, uint8)

//go:linkname cmykrgba color.sub_cmykrgba
func cmykrgba(c color.CMYK) (uint32, uint32, uint32, uint32) {
	return c.RGBA()
}

//go:linkname CMYKRGBA color.sub_cmykrgba
//go:noescape
func CMYKRGBA(c color.CMYK) (uint32, uint32, uint32, uint32)

//go:linkname nrgbargba color.sub_nrgbargba
func nrgbargba(c color.NRGBA) uint32 {
	return c.RGBA()
}

//go:linkname NRGBARGBA color.sub_nrgbargba
//go:noescape
func NRGBARGBA(c color.NRGBA) uint32

//go:linkname nrgba64rgba color.sub_nrgba64rgba
func nrgba64rgba(c color.NRGBA64) uint32 {
	return c.RGBA()
}

//go:linkname NRGBA64RGBA color.sub_nrgba64rgba
//go:noescape
func NRGBA64RGBA(c color.NRGBA64) uint32

//go:linkname paletteconvert color.sub_paletteconvert
func paletteconvert(p color.Palette, c color.Color) color.Color {
	return p.Convert(c)
}

//go:linkname PaletteConvert color.sub_paletteconvert
//go:noescape
func PaletteConvert(p color.Palette, c color.Color) color.Color

//go:linkname paletteindex color.sub_paletteindex
func paletteindex(p color.Palette, c color.Color) int {
	return p.Index(c)
}

//go:linkname PaletteIndex color.sub_paletteindex
//go:noescape
func PaletteIndex(p color.Palette, c color.Color) int

//go:linkname rgba64rgba color.sub_rgba64rgba
func rgba64rgba(c color.RGBA64) uint32 {
	return c.RGBA()
}

//go:linkname RGBA64RGBA color.sub_rgba64rgba
//go:noescape
func RGBA64RGBA(c color.RGBA64) uint32

//go:linkname alphargba color.sub_alphargba
func alphargba(c color.Alpha) uint32 {
	return c.RGBA()
}

//go:linkname AlphaRGBA color.sub_alphargba
//go:noescape
func AlphaRGBA(c color.Alpha) uint32

//go:linkname ModelFunc image/color.ModelFunc
//go:noescape
func ModelFunc(f func(color.Color) color.Color) color.Model

//go:linkname RGBToCMYK image/color.RGBToCMYK
//go:noescape
func RGBToCMYK(r, g, b uint8) (uint8, uint8, uint8, uint8)

//go:linkname RGBToYCbCr image/color.RGBToYCbCr
//go:noescape
func RGBToYCbCr(r, g, b uint8) (uint8, uint8, uint8)

//go:linkname YCbCrToRGB image/color.YCbCrToRGB
//go:noescape
func YCbCrToRGB(y, cb, cr uint8) (uint8, uint8, uint8)

//go:linkname grayrgba color.sub_grayrgba
func grayrgba(c color.Gray) uint32 {
	return c.RGBA()
}

//go:linkname GrayRGBA color.sub_grayrgba
//go:noescape
func GrayRGBA(c color.Gray) uint32

//go:linkname gray16rgba color.sub_gray16rgba
func gray16rgba(c color.Gray16) uint32 {
	return c.RGBA()
}

//go:linkname Gray16RGBA color.sub_gray16rgba
//go:noescape
func Gray16RGBA(c color.Gray16) uint32

//go:linkname ycbcrrgba color.sub_ycbcrrgba
func ycbcrrgba(c color.YCbCr) (uint32, uint32, uint32, uint32) {
	return c.RGBA()
}

//go:linkname YCbCrRGBA color.sub_ycbcrrgba
//go:noescape
func YCbCrRGBA(c color.YCbCr) (uint32, uint32, uint32, uint32)
