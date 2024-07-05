package filter

type InvertFilter struct{}

func (filter *InvertFilter) Convert(filterValues *FilterValues) {
	filterValues.NewRGBAValue.A = 255
	filterValues.NewRGBAValue.R, filterValues.NewRGBAValue.G, filterValues.NewRGBAValue.B = 255-filterValues.RGBAValues[0].R, 255-filterValues.RGBAValues[0].G, 255-filterValues.RGBAValues[0].B
}
