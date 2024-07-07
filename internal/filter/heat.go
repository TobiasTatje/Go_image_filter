package filter

import (
	"image/color"

	"bib.de/img_proc/internal/utils"
)

type HeatFilter struct{}

func (filter *HeatFilter) Convert(fv *FilterValues) {
	var np color.RGBA
	i := utils.Intensity(fv.RefOldImg.At(int(fv.X), int(fv.Y)).(color.RGBA))
	np.A = 255
	if i <= 42 {
		np.R, np.G, np.B = 0, 0, 0
		return
	}
	if i <= 84 {
		np.R, np.G, np.B = 0, 0, 255
		return
	}
	if i <= 126 {
		np.R, np.G, np.B = 0, 255, 255
		return
	}
	if i <= 168 {
		np.R, np.G, np.B = 0, 255, 0
		return
	}
	if i <= 210 {
		np.R, np.G, np.B = 255, 255, 0
		return
	}
	np.R, np.G, np.B = 255, 0, 0
	fv.RefNewImg.SetRGBA(int(fv.X), int(fv.Y), np)
}
