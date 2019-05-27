package clr

import (
	"fmt"
	"math"
)

// RGB Represents a point in the RGB Colorspace.
type RGB struct {
	R int `json:"r"`
	G int `json:"g"`
	B int `json:"b"`
}

// Valid checks if the RGB instance is a valid point in the RGB ColorSpace.
func (rgb RGB) Valid() bool {
	return rgb.R <= 255 && rgb.G <= 255 && rgb.B <= 255
}

// RGB returns the RGB components for the given point.
func (rgb RGB) RGB() (int, int, int) {
	return rgb.R, rgb.G, rgb.B
}

// RGBA is similar to RGB but adds an alpha channel. It conforms with the
// color.Color interface, so values need to be upscaled accordingly.
func (rgb RGB) RGBA() (r, g, b, a uint32) {
	r = uint32(rgb.R)
	r |= r << 8
	g = uint32(rgb.G)
	g |= g << 8
	b = uint32(rgb.B)
	b |= b << 8
	a = 255
	a |= a << 8
	return
}

// HSL converts RGB values into HSL ones in which
// H = 0 - 360, S = 0 - 100 and V = 0 - 100
func (rgb RGB) HSL() (int, int, int) {
	var h, s, l float64
	R := float64(rgb.R) / 255
	G := float64(rgb.G) / 255
	B := float64(rgb.B) / 255

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
	return int(h * 360), int(s * 100), int(l * 100)
}

// HSV converts RGB values into HSV ones in which
// H = 0 - 360, S = 0 - 100 and V = 0 - 100
func (rgb RGB) HSV() (int, int, int) {
	var h, s, v float64
	R := float64(rgb.R) / 255
	G := float64(rgb.G) / 255
	B := float64(rgb.B) / 255

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
	return int(h * 360), int(s * 100), int(v * 100)
}

// CMYK converts RGB colorspace into CMYK
func (rgb RGB) CMYK() (int, int, int, int) {
	r := float64(rgb.R) / 255.0
	g := float64(rgb.G) / 255.0
	b := float64(rgb.B) / 255.0

	var dc, dy, dm, dk float64

	dk = (1.0 - max(r, g, b))
	dc = (1 - r - dk) / (1 - dk)
	dm = (1 - g - dk) / (1 - dk)
	dy = (1 - b - dk) / (1 - dk)
	return int(dc * 100), int(dm * 100), int(dy * 100), int(dk * 100)
}

// Hex formats rgb in Hex
func (rgb RGB) Hex() string {
	return fmt.Sprintf("%02X%02X%02X", rgb.R, rgb.G, rgb.B)
}

// ColorName matches this RGB color to a color name
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

// Distance calculates the distance between colors in CIELAB using
// simple Euclidean distance.
func (rgb RGB) Distance(c Color) float64 {
	l1, a1, b1 := rgb.CIELAB()
	l2, a2, b2 := c.CIELAB()
	return math.Sqrt(
		math.Pow(l1-l2, 2) +
			math.Pow(a1-a2, 2) +
			math.Pow(b1-b2, 2))
}

// XYZ Converts to the XYZ Colorspace.
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

// CIELAB converts RGB to CIELAB, which is useful for comparing
// between colors how people actually view them.
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
