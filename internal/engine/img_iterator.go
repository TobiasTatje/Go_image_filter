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

	for y := s * sec_h; y < int(math.Min(float64((s+1)*sec_h), float64(b.Max.Y))); y++ {
		if filterDefs.Values.UsingEntireRow {
			var rgbaValues []color.RGBA
			for x := b.Min.X; x < b.Max.X; x++ {
				rgbaValues = append(rgbaValues, o_img.At(x, y).(color.RGBA))
			}
			filterDefs.Values.Row = rgbaValues
			filterDefs.Filter.Convert(&filterDefs.Values)
			for x := b.Min.X; x < b.Max.X; x++ {
				n_img.SetRGBA(x, y, filterDefs.Values.Row[x])
			}
		} else {
			for x := b.Min.X; x < b.Max.X; x++ {
				filterDefs.Values.IsValueSet[0] = true
				filterDefs.Values.RGBAValues[0] = o_img.At(x, y).(color.RGBA)

				filterDefs.Values.X = int64(x)
				filterDefs.Values.Y = int64(y)

				if filterDefs.Values.UsingNeighbors {
					if x > b.Min.X {
						filterDefs.Values.IsValueSet[1] = true
						filterDefs.Values.RGBAValues[1] = o_img.At(x-1, y).(color.RGBA)
					}
					if y < b.Max.Y {
						filterDefs.Values.IsValueSet[2] = true
						filterDefs.Values.RGBAValues[2] = o_img.At(x, y+1).(color.RGBA)
					}
					if x < b.Max.X {
						filterDefs.Values.IsValueSet[3] = true
						filterDefs.Values.RGBAValues[3] = o_img.At(x+1, y).(color.RGBA)
					}
					if y > b.Min.Y {
						filterDefs.Values.IsValueSet[4] = true
						filterDefs.Values.RGBAValues[4] = o_img.At(x, y-1).(color.RGBA)
					}
				}
				filterDefs.Filter.Convert(&filterDefs.Values)
				n_img.SetRGBA(x, y, filterDefs.Values.NewRGBAValue)
			}
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

		wg.Add(threadCount)

		for j := range threadCount {
			go sectorIterate(o_img, n_img, j, sec_h, &wg, filter)
		}

		wg.Wait()

		o_img = n_img
	}
	return o_img
}
