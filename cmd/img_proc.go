package main

import (
	"errors"
	"flag"
	"fmt"
	"image"
	"os"
	"strconv"
	"strings"
	"time"

	"bib.de/img_proc/internal/engine"
	"bib.de/img_proc/internal/filter"
	"bib.de/img_proc/internal/io"
)

var helpFlag = flag.Bool("h", false, "display flag help")
var pictureFlag = flag.String("p", "[relative path][name].png", "Relative path to the picture Ex: '.\\res\\temp.png'")
var filterFlag = flag.String("f", "", "to be applied Filter(s). Can be chained with \",\".\nPossible Values are: \n"+filterHelpMsg())
var iterationFlag = flag.Int("i", -1, "Number of iterations of Applying. If not specified filter default values will be used")
var threadFlag = flag.Int("t", -1, "Number of threads to be used for converting. If not specified, uses max available Cores of the System.")
var radiusFlag = flag.Int("r", -1, "Radius in pixel from the middle of image")
var radiusPercentFlag = flag.Int("rp,", -1, "Radius in percent of the smallest side of the image")
var coordsOffsetFlag = flag.String("o", "0,0", "Offset in Pixel from the middle of picture. Usage: x,y")

func applyFilterValuesFromFlags(fv *filter.FilterValues) {
	if *iterationFlag > 0 {
		fv.I = uint8(*iterationFlag)
	}
	if *radiusFlag > 0 {
		fv.UsingRadius = true
		fv.Rad = int64(*radiusFlag)
	}
	if *radiusPercentFlag > 0 {
		fv.UsingRadius = true
		fv.RadInPercent = true
		fv.RPercent = uint8(*radiusPercentFlag)
	}

	coords := strings.Split(*coordsOffsetFlag, ",")

	if len(coords) == 2 {
		if X, err := strconv.Atoi(coords[0]); errors.Is(err, os.ErrNotExist) {
			fv.X_Offset = 0
		} else {
			fv.X_Offset = int64(X)
		}

		if Y, err := strconv.Atoi(coords[1]); errors.Is(err, os.ErrNotExist) {
			fv.Y_Offset = 0
		} else {
			fv.Y_Offset = int64(Y)
		}
	}
}

func filterHelpMsg() (output string) {
	for _, f := range filter.FilterDefs {
		output += fmt.Sprintf("%s : %s\n", f.Flag, f.Name)
	}
	return
}

func checkFilterFlag(flags []string) bool {
	c := 0
	for _, flag := range flags {
		for _, f := range filter.FilterDefs {
			if f.Flag == flag {
				c++
				continue
			}
		}
	}
	return c == len(flags)
}

func getFilterFromFlag(flagsString []string) (filterArr []filter.FilterDef) {
	for _, flag := range flagsString {
		for _, filter := range filter.FilterDefs {
			if filter.Flag == flag {
				filterArr = append(filterArr, filter)
			}
		}
	}
	return
}

func applyFilter(img image.Image, filter filter.FilterDef) (img_n image.Image) {
	applyFilterValuesFromFlags(&filter.Values)
	fmt.Printf("Start converting using %s\n", filter.Name)
	t := time.Now()
	img_n = engine.Iterate(img, filter, *threadFlag)
	fmt.Printf("Converted img using %s in %d ms\n\n", filter.Name, time.Since(t).Milliseconds())
	return
}

func newFilename(s string, appliedFilter []filter.FilterDef) string {
	filterString := ""
	for _, f := range appliedFilter {
		filterString += "_" + f.Name
	}

	return strings.Replace(s, ".png", filterString+".png", 1)
}

func main() {

	flag.Parse()
	if *helpFlag {
		flag.PrintDefaults()
		return
	}

	if _, err := os.Stat(*pictureFlag); errors.Is(err, os.ErrNotExist) {
		println("Provided Picture (" + *pictureFlag + ") does not exists")
		return
	}

	filterFlags := strings.Split(*filterFlag, ",")

	if !checkFilterFlag(filterFlags) {
		println("Specified flag is not valid\nPossible Values:\n")
		println(filterHelpMsg())
		return
	}

	println("Image filter by Tobias Tatje\nStart processing")

	t := time.Now()
	img, err := io.ReadPng(*pictureFlag)
	if err != nil {
		panic("Error while reading: " + err.Error())
	}
	fmt.Printf("Read img in %d ms\n\n", time.Since(t).Milliseconds())

	selectedFilter := getFilterFromFlag(filterFlags)

	for _, filter := range selectedFilter {
		img_n := applyFilter(img, filter)
		img = img_n
	}

	t = time.Now()
	newName := newFilename(*pictureFlag, selectedFilter)
	if err = io.WritePng(newName, img); err != nil {
		panic("Error while Writing: " + err.Error())
	}
	fmt.Printf("Saved img %s in %d ms\n\n", newName, time.Since(t).Milliseconds())
}
