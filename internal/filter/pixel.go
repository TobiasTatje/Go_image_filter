package filter

import "image/color"

type PixelFilter struct{}

func (filter PixelFilter) Convert(fv *FilterValues) {
	if fv.X%int64(fv.CurrentI+1) != 0 {
		return
	}
	if fv.Y%int64(fv.CurrentI+1) != 0 {
		return
	}
	var div, R, G, B int64
	for x := fv.X; x < int64(fv.CurrentI)+1+fv.X; x++ {
		for y := fv.Y; y < int64(fv.CurrentI)+1+fv.Y; y++ {
			if x < int64(fv.Bounds.Max.X) && y < int64(fv.Bounds.Max.Y) {
				div++
				p := fv.RefOldImg.At(int(x), int(y)).(color.RGBA)
				R += int64(p.R)
				G += int64(p.G)
				B += int64(p.B)
			}
		}
	}
	var np color.RGBA
	np.R, np.G, np.B, np.A = uint8(R/div), uint8(G/div), uint8(B/div), 255
	for x := fv.X; x < int64(fv.CurrentI)+1+fv.X; x++ {
		for y := fv.Y; y < int64(fv.CurrentI)+1+fv.Y; y++ {
			fv.RefNewImg.SetRGBA(int(x), int(y), np)
		}
	}
}
