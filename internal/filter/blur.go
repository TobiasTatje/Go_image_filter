package filter

type BlurFilter struct{}

func (filter *BlurFilter) Convert(filterValues *FilterValues) {
	var div int64
	for i := 0; i < len(filterValues.IsValueSet); i++ {
		if filterValues.IsValueSet[i] {
			div++
		}
	}

	var R, G, B, A int64

	for i := 0; i < len(filterValues.RGBAValues); i++ {
		R += int64(filterValues.RGBAValues[i].R)
		G += int64(filterValues.RGBAValues[i].G)
		B += int64(filterValues.RGBAValues[i].B)
	}

	A = 255

	filterValues.NewRGBAValue.R = uint8(R / div)
	filterValues.NewRGBAValue.G = uint8(G / div)
	filterValues.NewRGBAValue.B = uint8(B / div)
	filterValues.NewRGBAValue.A = uint8(A)
}
