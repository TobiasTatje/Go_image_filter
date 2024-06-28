package filter

import (
	"image/color"

	"bib.de/img_proc/internal/utils"
)

type ComicFilter struct{}

func (filter *ComicFilter) Convert(values [5]color.RGBA, isSet [5]bool, filterValues FilterValues) (new color.RGBA) {
	i := utils.Intensity(values[0])
	new.A = 255
	if i <= 85 {
		new.R = 42
		new.G, new.B = new.R, new.R
		return
	}
	if i <= 170 {
		new.R = 128
		new.G, new.B = new.R, new.R
		return
	}
	new.R = 212
	new.G, new.B = new.R, new.R
	return
}
