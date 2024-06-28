package filter

type FilterDef struct {
	Name   string
	Flag   string
	Filter Filterer
	Values FilterValues
}

type FilterValues struct {
	NeedImgBounds      bool
	W, H               int64 //Width and Height of img. NEEDS TO BE SET AT RUNTIME!
	UsingPosition      bool
	X_Offset, Y_Offset int64 // Origin transformation of Filter position from middle, and radius of filter
	X, Y               int64 //Current position in image
	UsingRadius        bool
	RadInPercent       bool  // If radius shall be measured in percent of smallest picture
	Rad                int64 // Radius of filter
	RPercent           uint8 // Radius in percent.
	I                  uint8 //IterationCount
	UsingNeighbors     bool  //considering neighbors
}

var FilterDefs = []FilterDef{
	{"blur", "b", &BlurFilter{}, defBlurVal()},
	{"comic", "c", &ComicFilter{}, defComicVal()},
	{"heat", "h", &HeatFilter{}, defHeatVal()},
	{"edge", "e", &EdgeFilter{}, defEdgeVal()},
	{"invert", "i", &InvertFilter{}, defInvertVal()},
	{"spot", "s", &SpotFilter{}, defSpotVal()},
}

//Functions for default Values, change as needed

func defBlurVal() (fv FilterValues) {
	fv.I = 10
	setNoNeighbors(&fv)
	return
}

func defComicVal() (fv FilterValues) {
	defaultIterations(&fv)
	return
}

func defHeatVal() (fv FilterValues) {
	defaultIterations(&fv)
	return
}

func defEdgeVal() (fv FilterValues) {
	defaultIterations(&fv)
	setNoNeighbors(&fv)
	return
}

func defInvertVal() (fv FilterValues) {
	defaultIterations(&fv)
	return
}

func defSpotVal() (fv FilterValues) {
	defaultIterations(&fv)
	fv.NeedImgBounds = true
	fv.UsingRadius = true
	fv.RadInPercent = true
	fv.RPercent = 60
	return
}

func setNoNeighbors(fv *FilterValues) {
	fv.UsingNeighbors = true
}

func defaultIterations(fv *FilterValues) {
	fv.I = 1
}
