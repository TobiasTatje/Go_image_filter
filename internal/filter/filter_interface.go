package filter

type Filterer interface {
	Convert(filterValues *FilterValues)
}
