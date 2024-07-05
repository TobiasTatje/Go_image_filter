package filter

import (
	"bib.de/img_proc/internal/utils"
)

type ComicFilter struct{}

func (filter *ComicFilter) Convert(filterValues *FilterValues) {
	i := utils.Intensity(filterValues.RGBAValues[0])
	filterValues.NewRGBAValue.A = 255
	if i <= 85 {
		filterValues.NewRGBAValue.R = 42
		filterValues.NewRGBAValue.G, filterValues.NewRGBAValue.B = filterValues.NewRGBAValue.R, filterValues.NewRGBAValue.R
		return
	}
	if i <= 170 {
		filterValues.NewRGBAValue.R = 128
		filterValues.NewRGBAValue.G, filterValues.NewRGBAValue.B = filterValues.NewRGBAValue.R, filterValues.NewRGBAValue.R
		return
	}
	filterValues.NewRGBAValue.R = 212
	filterValues.NewRGBAValue.G, filterValues.NewRGBAValue.B = filterValues.NewRGBAValue.R, filterValues.NewRGBAValue.R
}
