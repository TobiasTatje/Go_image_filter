package filter

import (
	"bib.de/img_proc/internal/utils"
)

type HeatFilter struct{}

func (filter *HeatFilter) Convert(filterValues *FilterValues) {
	i := utils.Intensity(filterValues.RGBAValues[0])
	filterValues.NewRGBAValue.A = 255
	if i <= 42 {
		filterValues.NewRGBAValue.R, filterValues.NewRGBAValue.G, filterValues.NewRGBAValue.B = 0, 0, 0
		return
	}
	if i <= 84 {
		filterValues.NewRGBAValue.R, filterValues.NewRGBAValue.G, filterValues.NewRGBAValue.B = 0, 0, 255
		return
	}
	if i <= 126 {
		filterValues.NewRGBAValue.R, filterValues.NewRGBAValue.G, filterValues.NewRGBAValue.B = 0, 255, 255
		return
	}
	if i <= 168 {
		filterValues.NewRGBAValue.R, filterValues.NewRGBAValue.G, filterValues.NewRGBAValue.B = 0, 255, 0
		return
	}
	if i <= 210 {
		filterValues.NewRGBAValue.R, filterValues.NewRGBAValue.G, filterValues.NewRGBAValue.B = 255, 255, 0
		return
	}
	filterValues.NewRGBAValue.R, filterValues.NewRGBAValue.G, filterValues.NewRGBAValue.B = 255, 0, 0
}
