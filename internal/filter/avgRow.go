package filter

import (
	"bib.de/img_proc/internal/utils"
)

type AvgRowFilter struct{}

func (filter *AvgRowFilter) Convert(filterValues *FilterValues) {
	c := utils.AvgRgb(filterValues.Row)
	for i := range filterValues.Row {
		filterValues.Row[i] = c
	}
}
