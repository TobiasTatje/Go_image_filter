package engine

import (
	"fmt"
	"image"
	"image/color"
	"math"
	"runtime"
	"sync"

	filter "bib.de/img_proc/internal/filter"
)

func sectorIterate(o_img image.Image, n_img *image.RGBA, s, sec_h int, wg *sync.WaitGroup, filterDefs filter.FilterDef) {
	defer wg.Done()

	b := o_img.Bounds()

	var values [5]color.RGBA
	var isSet [5]bool

	for x := b.Min.X; x < b.Max.X; x++ {
		for y := s * sec_h; y < int(math.Min(float64((s+1)*sec_h), float64(b.Max.Y))); y++ {
			isSet[0] = true
			values[0] = o_img.At(x, y).(color.RGBA)

			filterDefs.Values.X = int64(x)
			filterDefs.Values.Y = int64(y)

			if filterDefs.Values.UsingNeighbors {
				if x > b.Min.X {
					isSet[1] = true
					values[1] = o_img.At(x-1, y).(color.RGBA)
				}
				if y < b.Max.Y {
					isSet[2] = true
					values[2] = o_img.At(x, y+1).(color.RGBA)
				}
				if x < b.Max.X {
					isSet[3] = true
					values[3] = o_img.At(x+1, y).(color.RGBA)
				}
				if y > b.Min.Y {
					isSet[4] = true
					values[4] = o_img.At(x, y-1).(color.RGBA)
				}
			}
			n_img.SetRGBA(x, y, filterDefs.Filter.Convert(values, isSet, filterDefs.Values))
		}
	}
}

func Iterate(o_img image.Image, filter filter.FilterDef, threadCount int) image.Image {

	var wg sync.WaitGroup

	b := o_img.Bounds()

	w, h := b.Max.X-b.Min.X, b.Max.Y-b.Min.Y

	if filter.Values.NeedImgBounds {
		filter.Values.W = int64(w)
		filter.Values.H = int64(h)
	}

	if filter.Values.UsingRadius {
		if filter.Values.RadInPercent {
			filter.Values.Rad = int64(math.Min(float64(w), float64(h)) * (float64(filter.Values.RPercent) / 100.0))
		}
	}

	if threadCount <= 0 {
		threadCount = runtime.GOMAXPROCS(runtime.NumCPU())
	}

	fmt.Printf("Using %d Threads\n", threadCount)

	sec_h := int(math.Ceil(float64(h) / float64(threadCount)))

	for i := 0; i < int(filter.Values.I); i++ {

		n_img := image.NewRGBA(image.Rect(b.Min.X, b.Min.Y, b.Max.X, b.Max.Y))

		for j := range threadCount {
			wg.Add(1)
			go sectorIterate(o_img, n_img, j, sec_h, &wg, filter)
		}

		wg.Wait()

		o_img = n_img
	}
	return o_img
}
