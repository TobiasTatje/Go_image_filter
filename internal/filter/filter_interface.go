package filter

import (
	"image/color"
)

type Filterer interface {
	Convert(values [5]color.RGBA, isSet [5]bool, filterValues FilterValues) color.RGBA
}
