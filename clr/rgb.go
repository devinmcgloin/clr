package clr

import (
	"fmt"
	"math"
)

type RGB struct {
	R, G, B uint8
}

func (c RGB) RGB() (r, g, b uint8) {
	return c.R, c.G, c.B
}

// HSL converts RGB values into HSL ones in which
// H = 0 - 360, S = 0 - 100 and V = 0 - 100
func (c RGB) HSL() (uint16, uint8, uint8) {
	var h, s, l float64
	R := float64(c.R) / 255
	G := float64(c.G) / 255
	B := float64(c.B) / 255

	minVal := min(R, G, B)
	maxVal := max(R, G, B)
	delta := maxVal - minVal

	l = (maxVal + minVal) / 2.0

	if delta == 0 {
		h = 0
		s = 0
	} else {
		d := maxVal - minVal
		if l > 0.5 {
			s = d / (2 - maxVal - minVal)
		} else {
			s = d / (maxVal + minVal)
		}

		switch maxVal {
		case R:
			if G < B {
				h = (G-B)/d + 6
			} else {
				h = (G-B)/d + 0
			}
		case G:
			h = (B-R)/d + 2
		case B:
			h = (R-G)/d + 4
		}
		h /= 6
	}
	return uint16(h * 360), uint8(s * 100), uint8(l * 100)
}

// HSV converts RGB values into HSV ones in which
// H = 0 - 360, S = 0 - 100 and V = 0 - 100
func (c RGB) HSV() (uint16, uint8, uint8) {
	var h, s, v float64
	R := float64(c.R) / 255
	G := float64(c.G) / 255
	B := float64(c.B) / 255

	minVal := min(R, G, B)
	maxVal := max(R, G, B)
	delta := maxVal - minVal

	v = maxVal

	if delta == 0 {
		h = 0
		s = 0
	} else {
		d := maxVal - minVal
		s = delta / maxVal
		switch maxVal {
		case R:
			if G < B {
				h = (G-B)/d + 6
			} else {
				h = (G-B)/d + 0
			}
		case G:
			h = (B-R)/d + 2
		case B:
			h = (R-G)/d + 4
		}
		h /= 6
	}
	return uint16(h * 360), uint8(s * 100), uint8(v * 100)
}

func (rgb RGB) CMYK() (uint8, uint8, uint8, uint8) {
	r := float64(rgb.R) / 255.0
	g := float64(rgb.G) / 255.0
	b := float64(rgb.B) / 255.0

	var dc, dy, dm, dk float64

	dk = (1.0 - max(r, g, b))
	dc = (1 - r - dk) / (1 - dk)
	dm = (1 - g - dk) / (1 - dk)
	dy = (1 - b - dk) / (1 - dk)
	return uint8(dc * 100), uint8(dm * 100), uint8(dy * 100), uint8(dk * 100)
}

func (rgb RGB) Hex() string {
	return fmt.Sprintf("%02X%02X%02X", rgb.R, rgb.G, rgb.B)
}

func (rgb RGB) ColorName(colors ColorTable) ColorSpace {
	var minHex string
	minDist := math.MaxFloat64
	var hex = rgb.Hex()

	for _, c := range colors.Iterate() {
		if c.Hex() == hex {
			return colors.Lookup(c.Hex())
		}
		dist := rgb.Distance(c)
		if dist < minDist {
			minHex = c.Hex()
			minDist = dist
		}
	}
	return colors.Lookup(minHex)
}

func (rgb RGB) Distance(c Color) float64 {
	l1, a1, b1 := rgb.CIELAB()
	l2, a2, b2 := c.CIELAB()
	return math.Sqrt(
		math.Pow(l1-l2, 2) +
			math.Pow(a1-a2, 2) +
			math.Pow(b1-b2, 2))
}

func (rgb RGB) XYZ() (float64, float64, float64) {
	r := float64(rgb.R) / 255.0
	g := float64(rgb.G) / 255.0
	b := float64(rgb.B) / 255.0

	if r > 0.04045 {
		r = math.Pow((r+0.055)/1.055, 2.4)
	} else {
		r = r / 12.92
	}

	if g > 0.04045 {
		g = math.Pow((g+0.055)/1.055, 2.4)
	} else {
		g = g / 12.92
	}

	if b > 0.04045 {
		b = math.Pow((b+0.055)/1.055, 2.4)
	} else {
		b = b / 12.92
	}

	r *= 100
	g *= 100
	b *= 100

	x := r*0.4124 + g*0.3576 + b*0.1805
	y := r*0.2126 + g*0.7152 + b*0.0722
	z := r*0.0193 + g*0.1192 + b*0.9505
	return x, y, z
}

func (rgb RGB) CIELAB() (l, a, b float64) {
	x, y, z := rgb.XYZ()
	x /= 95.682
	y /= 100
	z /= 92.149

	if x > 0.008856 {
		x = math.Pow(x, 1.0/3.0)
	} else {
		x = (7.787 * x) + (16.0 / 116)
	}
	if y > 0.008856 {
		y = math.Pow(y, 1.0/3.0)
	} else {
		y = (7.787 * y) + (16.0 / 116)
	}

	if z > 0.008856 {
		z = math.Pow(z, 1.0/3.0)
	} else {
		z = (7.787 * z) + (16.0 / 116)
	}

	l = (116 * x) - 16
	a = 500 * (x - y)
	b = 200 * (y - z)
	return l, a, b
}
