package filter

import "image/color"

type InvertFilter struct{}

func (filter *InvertFilter) Convert(fv *FilterValues) {
	var np color.RGBA
	np.A = 255
	p := fv.RefOldImg.At(int(fv.X), int(fv.Y)).(color.RGBA)
	np.R, np.G, np.B = 255-p.R, 255-p.G, 255-p.B
	fv.RefNewImg.SetRGBA(int(fv.X), int(fv.Y), np)
}
