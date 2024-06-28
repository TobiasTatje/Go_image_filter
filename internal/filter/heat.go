package filter

import (
	"image/color"

	"bib.de/img_proc/internal/utils"
)

type HeatFilter struct{}

func (filter *HeatFilter) Convert(values [5]color.RGBA, isSet [5]bool, filterValues FilterValues) (new color.RGBA) {
	i := utils.Intensity(values[0])
	new.A = 255
	if i <= 42 {
		new.R, new.G, new.B = 0, 0, 0
		return
	}
	if i <= 84 {
		new.R, new.G, new.B = 0, 0, 255
		return
	}
	if i <= 126 {
		new.R, new.G, new.B = 0, 255, 255
		return
	}
	if i <= 168 {
		new.R, new.G, new.B = 0, 255, 0
		return
	}
	if i <= 210 {
		new.R, new.G, new.B = 255, 255, 0
		return
	}
	new.R, new.G, new.B = 255, 0, 0
	return
}
