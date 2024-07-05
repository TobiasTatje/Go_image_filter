package utils

import (
	"image/color"
	"math"
)

func Intensity(c color.RGBA) float64 {
	return 0.2126*float64(c.R) + 0.7152*float64(c.G) + 0.0722*float64(c.B)
}

func DistanceFromOriginWithOffset(s_w, s_h, x, y, x_offset, y_offset int) int {
	return int(math.Sqrt(math.Pow(math.Abs(float64(x-s_w-x_offset)), 2) + math.Pow(math.Abs(float64(y-s_h-y_offset)), 2)))
}

func SpotColor(d, r int, c color.RGBA) (n color.RGBA) {
	n.A = c.A
	n.R = spotColor(r, d, c.R)
	n.G = spotColor(r, d, c.G)
	n.B = spotColor(r, d, c.B)
	return
}

func spotColor(r, d int, c uint8) uint8 {
	if d == 0 {
		return c
	}
	if d > r {
		return 0
	}

	i := float32(d) / float32(r)

	return uint8(int(c) - int(float32(c)*i))
}

func RadixRgba(c []color.RGBA) []color.RGBA {
	var n0, n1 []color.RGBA
	comparator := 0b1
	for range 8 * 3 {
		for j := range len(c) {
			val := uint(c[j].R) | uint(c[j].G)<<8 | uint(c[j].B)<<16
			if uint(comparator)&(val) == 0 {
				n0 = append(n0, c[j])
			} else {
				n1 = append(n1, c[j])
			}
		}
		c = append(n0, n1...)
		n0, n1 = n0[:0], n1[:0]
		comparator <<= 1
	}
	return c
}
