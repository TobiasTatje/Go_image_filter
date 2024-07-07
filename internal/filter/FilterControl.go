package filter

import (
	"image"
	"image/color"
)

type FilterDef struct {
	Name   string
	Flag   string
	Filter Filterer
	Values FilterValues
}

type FilterValues struct {
	Bounds             image.Rectangle
	UsingPosition      bool
	X_Offset, Y_Offset int64 //Origin transformation of Filter position from middle, and radius of filter
	X, Y               int64 //Current position in image
	UsingRadius        bool
	RadInPercent       bool  //If radius shall be measured in percent of smallest picture
	Rad                int64 //Radius of filter
	RPercent           uint8 //Radius in percent
	UseIteration       bool  //if the filter shall use iterations
	I                  uint8 //Iterations
	CurrentI           uint8
	UsingEntireRow     bool //If filter uses the entire row instead of one pixels and neighbors
	Row                []color.RGBA
	RefNewImg          *image.RGBA
	RefOldImg          image.Image
}

var FilterDefs = []FilterDef{
	{"blur", "b", &BlurFilter{}, defBlurVal()},
	{"comic", "c", &ComicFilter{}, defComicVal()},
	{"heat", "h", &HeatFilter{}, defHeatVal()},
	{"edge", "e", &EdgeFilter{}, defEdgeVal()},
	{"invert", "i", &InvertFilter{}, defInvertVal()},
	{"spot", "s", &SpotFilter{}, defSpotVal()},
	{"sortRow", "sr", &SortRowFilter{}, defSortRowFilter()},
	{"averageRow", "ar", &AvgRowFilter{}, defAvgRowFilter()},
	{"pixel", "p", &PixelFilter{}, defPixelFilter()},
}

//Functions for default Values, change as needed

func defBlurVal() (fv FilterValues) {
	fv.UseIteration = true
	fv.I = 10
	return
}

func defComicVal() (fv FilterValues) {
	return
}

func defHeatVal() (fv FilterValues) {
	return
}

func defEdgeVal() (fv FilterValues) {
	return
}

func defInvertVal() (fv FilterValues) {
	return
}

func defSpotVal() (fv FilterValues) {
	fv.UsingRadius = true
	fv.RadInPercent = true
	fv.RPercent = 60
	return
}

func defSortRowFilter() (fv FilterValues) {
	fv.UsingEntireRow = true
	return
}

func defAvgRowFilter() (fv FilterValues) {
	fv.UsingEntireRow = true
	return
}

func defPixelFilter() (fv FilterValues) {
	fv.UseIteration = false
	fv.I = 10
	return
}
