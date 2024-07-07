package engine

import (
	"fmt"
	"image"
	"math"
	"runtime"
	"sync"

	filter "bib.de/img_proc/internal/filter"
)

func sectorIterate(s, sec_h int64, wg *sync.WaitGroup, filterDefs filter.FilterDef) {
	defer wg.Done()

	for y := s * sec_h; y < int64(math.Min(float64((s+1)*sec_h), float64(filterDefs.Values.Bounds.Max.Y))); y++ {
		filterDefs.Values.Y = y
		if filterDefs.Values.UsingEntireRow {
			filterDefs.Values.X = -1
			filterDefs.Filter.Convert(&filterDefs.Values)
		} else {
			for x := filterDefs.Values.Bounds.Min.X; x < filterDefs.Values.Bounds.Max.X; x++ {
				filterDefs.Values.X = int64(x)
				filterDefs.Filter.Convert(&filterDefs.Values)
			}
		}
	}
}

func Iterate(o_img image.Image, filter filter.FilterDef, threadCount int) image.Image {

	var wg sync.WaitGroup
	filter.Values.RefOldImg = o_img
	filter.Values.Bounds = o_img.Bounds()

	if filter.Values.UsingRadius {
		if filter.Values.RadInPercent {
			w, h := filter.Values.Bounds.Max.X-filter.Values.Bounds.Min.X, filter.Values.Bounds.Max.Y-filter.Values.Bounds.Min.Y
			filter.Values.Rad = int64(math.Min(float64(w), float64(h)) * (float64(filter.Values.RPercent) / 100.0))
		}
	}
	filter.Values.RefOldImg = o_img

	if threadCount <= 0 {
		threadCount = runtime.GOMAXPROCS(runtime.NumCPU())
	}

	fmt.Printf("Using %d Threads\n", threadCount)

	sec_h := int64(math.Ceil(float64(filter.Values.Bounds.Max.Y-filter.Values.Bounds.Min.Y) / float64(threadCount)))

	for i := 0; i < int(filter.Values.I); i++ {

		n_img := image.NewRGBA(image.Rect(
			filter.Values.Bounds.Min.X,
			filter.Values.Bounds.Min.Y,
			filter.Values.Bounds.Max.X,
			filter.Values.Bounds.Max.Y))
		filter.Values.RefNewImg = n_img
		wg.Add(threadCount)

		for j := range threadCount {
			go sectorIterate(int64(j), sec_h, &wg, filter)
		}

		wg.Wait()

		filter.Values.RefOldImg = filter.Values.RefNewImg
	}
	return filter.Values.RefOldImg
}
