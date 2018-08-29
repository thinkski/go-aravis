package aravis

import (
	"image"
	"image/color"
)

type BayerRG struct {
	// Pix holds the image's pixels, in bayer order. The pixel at
	// (x, y) starts at Pix[(y-Rect.Min.Y)*Stride + (x-Rect.Min.X)].
	Pix []uint8
	// Stride is the Pix stride (in bytes) between vertically adjacent pixels.
	Stride int
	// Rect is the image's bounds.
	Rect image.Rectangle
}

func NewBayerRG(r image.Rectangle) *BayerRG {
	w, h := r.Dx(), r.Dy()
	pix := make([]uint8, w*h)
	return &BayerRG{pix, w, r}
}

func (p *BayerRG) ColorModel() color.Model { return color.RGBAModel }

func (p *BayerRG) Bounds() image.Rectangle { return p.Rect }

// At returns an RGBA pixel with simple nearest-neighbor debayering
func (p *BayerRG) At(x, y int) color.Color {
	if x&1 == 0 && y&1 == 0 {
		// top-left: red
		return color.RGBA{
			p.Pix[(y-p.Rect.Min.Y)*p.Stride+(x-p.Rect.Min.X)],
			p.Pix[(y-p.Rect.Min.Y)*p.Stride+(x+1-p.Rect.Min.X)],
			p.Pix[(y+1-p.Rect.Min.Y)*p.Stride+(x+1-p.Rect.Min.X)],
			0,
		}
	} else if x&1 == 1 && y&1 == 0 {
		// top-right: green
		return color.RGBA{
			p.Pix[(y-p.Rect.Min.Y)*p.Stride+(x-1-p.Rect.Min.X)],
			p.Pix[(y-p.Rect.Min.Y)*p.Stride+(x-p.Rect.Min.X)],
			p.Pix[(y+1-p.Rect.Min.Y)*p.Stride+(x-p.Rect.Min.X)],
			0,
		}
	} else if x&1 == 0 && y&1 == 1 {
		// bottom-left: green
		return color.RGBA{
			p.Pix[(y-1-p.Rect.Min.Y)*p.Stride+(x-p.Rect.Min.X)],
			p.Pix[(y-p.Rect.Min.Y)*p.Stride+(x-p.Rect.Min.X)],
			p.Pix[(y-p.Rect.Min.Y)*p.Stride+(x+1-p.Rect.Min.X)],
			0,
		}
	} else {
		// bottom-right: blue
		return color.RGBA{
			p.Pix[(y-1-p.Rect.Min.Y)*p.Stride+(x-1-p.Rect.Min.X)],
			p.Pix[(y-1-p.Rect.Min.Y)*p.Stride+(x-p.Rect.Min.X)],
			p.Pix[(y-p.Rect.Min.Y)*p.Stride+(x-p.Rect.Min.X)],
			0,
		}
	}
}
