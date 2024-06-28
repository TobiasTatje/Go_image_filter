package filter

import (
	"image/color"

	"bib.de/img_proc/internal/utils"
)

type SpotFilter struct{}

func (filter *SpotFilter) Convert(values [5]color.RGBA, isSet [5]bool, filterValues FilterValues) (new color.RGBA) {
	new = utils.SpotColor(utils.DistanceFromOriginWithOffset(
		int(filterValues.W)/2,
		int(filterValues.H)/2,
		int(filterValues.X),
		int(filterValues.Y),
		int(filterValues.X_Offset),
		int(filterValues.Y_Offset)),
		int(filterValues.Rad),
		values[0])
	return
}
