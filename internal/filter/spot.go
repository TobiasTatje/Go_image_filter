package filter

import (
	"bib.de/img_proc/internal/utils"
)

type SpotFilter struct{}

func (filter *SpotFilter) Convert(filterValues *FilterValues) {
	filterValues.NewRGBAValue = utils.SpotColor(
		utils.DistanceFromOriginWithOffset(
			int(filterValues.W)/2,
			int(filterValues.H)/2,
			int(filterValues.X),
			int(filterValues.Y),
			int(filterValues.X_Offset),
			int(filterValues.Y_Offset)),
		int(filterValues.Rad),
		filterValues.RGBAValues[0])

}
