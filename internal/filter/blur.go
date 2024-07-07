package filter

import "image/color"

type BlurFilter struct{}

func (filter *BlurFilter) Convert(fv *FilterValues) {
	var div, R, G, B, A int64
	//Upper and lower bound

	if fv.X > int64(fv.Bounds.Min.X) {
		div++
		p := fv.RefOldImg.At(int(fv.X-1), int(fv.Y)).(color.RGBA)
		R += int64(p.R)
		G += int64(p.G)
		B += int64(p.B)
		A += int64(p.A)
	}
	if fv.X < int64(fv.Bounds.Max.X) {
		div++
		p := fv.RefOldImg.At(int(fv.X+1), int(fv.Y)).(color.RGBA)
		R += int64(p.R)
		G += int64(p.G)
		B += int64(p.B)
		A += int64(p.A)
	}
	if fv.Y > int64(fv.Bounds.Min.Y) {
		div++
		p := fv.RefOldImg.At(int(fv.X), int(fv.Y-1)).(color.RGBA)
		R += int64(p.R)
		G += int64(p.G)
		B += int64(p.B)
		A += int64(p.A)
	}
	if fv.Y < int64(fv.Bounds.Max.Y) {
		div++
		p := fv.RefOldImg.At(int(fv.X), int(fv.Y-1)).(color.RGBA)
		R += int64(p.R)
		G += int64(p.G)
		B += int64(p.B)
		A += int64(p.A)
	}
	div++
	p := fv.RefOldImg.At(int(fv.X), int(fv.Y)).(color.RGBA)
	R += int64(p.R)
	G += int64(p.G)
	B += int64(p.B)
	A += int64(p.A)

	var np color.RGBA
	np.R = uint8(R / div)
	np.G = uint8(G / div)
	np.B = uint8(B / div)
	np.A = uint8(A / div)
	fv.RefNewImg.SetRGBA(int(fv.X), int(fv.Y), np)
}
