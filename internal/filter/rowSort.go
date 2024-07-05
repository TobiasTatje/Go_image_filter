package filter

import (
	"bib.de/img_proc/internal/utils"
)

type SortRowFilter struct{}

func (filter *SortRowFilter) Convert(filterValues *FilterValues) {
	filterValues.Row = utils.RadixRgba(filterValues.Row)
}
