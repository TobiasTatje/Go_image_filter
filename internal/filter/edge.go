package filter

import (
	"math"

	"bib.de/img_proc/internal/utils"
)

type EdgeFilter struct{}

func (filter *EdgeFilter) Convert(filterValues *FilterValues) {

	i := uint8(
		math.Min(255,
			math.Abs(utils.Intensity(filterValues.RGBAValues[1])-utils.Intensity(filterValues.RGBAValues[3]))+
				math.Abs(utils.Intensity(filterValues.RGBAValues[2])-utils.Intensity(filterValues.RGBAValues[4]))))
	filterValues.NewRGBAValue.A = 255
	filterValues.NewRGBAValue.R, filterValues.NewRGBAValue.G, filterValues.NewRGBAValue.B = i, i, i
}
