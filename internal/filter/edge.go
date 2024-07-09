package filter

import (
	"image/color"
	"math"

	"bib.de/img_proc/internal/utils"
)

type EdgeFilter struct{}

func (filter *EdgeFilter) Convert(fv *FilterValues) {
	var np color.RGBA
	above := fv.RefOldImg.At(int(fv.X)-1, int(fv.Y)).(color.RGBA)
	under := fv.RefOldImg.At(int(fv.X)+1, int(fv.Y)).(color.RGBA)
	left := fv.RefOldImg.At(int(fv.X), int(fv.Y)-1).(color.RGBA)
	right := fv.RefOldImg.At(int(fv.X), int(fv.Y)+1).(color.RGBA)
	i := uint8(
		math.Min(255,
			math.Abs(utils.Intensity(above)-utils.Intensity(under))+
				math.Abs(utils.Intensity(left)-utils.Intensity(right))))
	np.A = 255
	np.R, np.G, np.B = i, i, i
	fv.RefNewImg.SetRGBA(int(fv.X), int(fv.Y), np)
}
