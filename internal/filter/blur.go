package filter

import (
	"image/color"
)

type BlurFilter struct{}

func (filter *BlurFilter) Convert(values [5]color.RGBA, isSet [5]bool, filterValues FilterValues) (new color.RGBA) {
	var div int64
	for i := 0; i < len(isSet); i++ {
		if isSet[i] {
			div++
		}
	}

	var R, G, B, A int64

	for i := 0; i < len(values); i++ {
		R += int64(values[i].R)
		G += int64(values[i].G)
		B += int64(values[i].B)
	}

	A = 255

	new.R = uint8(R / div)
	new.G = uint8(G / div)
	new.B = uint8(B / div)
	new.A = uint8(A)
	return
}
