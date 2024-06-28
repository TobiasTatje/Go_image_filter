package filter

import "image/color"

type InvertFilter struct{}

func (filter *InvertFilter) Convert(values [5]color.RGBA, isSet [5]bool, filterValues FilterValues) (new color.RGBA) {
	new.A = 255
	new.R, new.G, new.B = 255-values[0].R, 255-values[0].G, 255-values[0].B
	return
}
