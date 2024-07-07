package filter

import (
	"image/color"

	"bib.de/img_proc/internal/utils"
)

type SpotFilter struct{}

func (filter *SpotFilter) Convert(fv *FilterValues) {
	np := utils.SpotColor(
		utils.DistanceFromOriginWithOffset(
			int(fv.Bounds.Max.X-fv.Bounds.Min.X)/2,
			int(fv.Bounds.Max.Y-fv.Bounds.Min.Y)/2,
			int(fv.X),
			int(fv.Y),
			int(fv.X_Offset),
			int(fv.Y_Offset)),
		int(fv.Rad),
		fv.RefOldImg.At(int(fv.X-1), int(fv.Y)).(color.RGBA))

	fv.RefNewImg.SetRGBA(int(fv.X), int(fv.Y), np)
}
