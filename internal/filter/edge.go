package filter

import (
	"image/color"
	"math"

	"bib.de/img_proc/internal/utils"
)

type EdgeFilter struct{}

func (filter *EdgeFilter) Convert(values [5]color.RGBA, isSet [5]bool, filterValues FilterValues) (new color.RGBA) {

	i := uint8(
		math.Min(255,
			math.Abs(utils.Intensity(values[1])-utils.Intensity(values[3]))+
				math.Abs(utils.Intensity(values[2])-utils.Intensity(values[4]))))
	new.A = 255
	new.R, new.G, new.B = i, i, i
	return
}
