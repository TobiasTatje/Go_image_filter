package filter

import (
	"image/color"

	"bib.de/img_proc/internal/utils"
)

type ComicFilter struct{}

func (filter *ComicFilter) Convert(fv *FilterValues) {
	i := utils.Intensity(fv.RefOldImg.At(int(fv.X), int(fv.Y)).(color.RGBA))
	var np color.RGBA
	np.A = 255
	if i <= 85 {
		np.R = 42
		np.G, np.B = np.R, np.R
		return
	}
	if i <= 170 {
		np.R = 128
		np.G, np.B = np.R, np.R
		return
	}
	np.R = 212
	np.G, np.B = np.R, np.R
	fv.RefNewImg.SetRGBA(int(fv.X), int(fv.Y), np)
}
