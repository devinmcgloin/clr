package clr

import "math"

// HSV represents the HSV Colorspace.
type HSV struct {
	H int `json:"h"`
	S int `json:"s"`
	V int `json:"v"`
}

// Valid confirms that the HSV Type is a valid point
// in the colorspace.
func (hsv HSV) Valid() bool {
	return hsv.H <= 360 && hsv.S <= 100 && hsv.V <= 100
}

// toRGB allows HSV to have access to all RGB Methods.
func (hsv HSV) toRGB() RGB {
	var r, g, b float64
	h := float64(hsv.H) / 360
	s := float64(hsv.S) / 100
	v := float64(hsv.V) / 100

	i := math.Floor(h * 6)
	f := h*6 - i
	p := v * (1 - s)
	q := v * (1 - f*s)
	t := v * (1 - (1-f)*s)

	switch int(i) % 6 {
	case 0:
		r = v
		g = t
		b = p
	case 1:
		r = q
		g = v
		b = p

	case 2:
		r = p
		g = v
		b = t
	case 3:
		r = p
		g = q
		b = v
	case 4:
		r = t
		g = p
		b = v
	case 5:
		r = v
		g = p
		b = q
	}

	return RGB{R: int(r * 255), G: int(g * 255), B: int(b * 255)}
}

// RGB Converts HSV to RGB using toRGB
func (hsv HSV) RGB() (int, int, int) {
	rgb := hsv.toRGB()
	return rgb.R, rgb.G, rgb.B
}

// RGBA is an alias of RGB.RGBA
func (hsv HSV) RGBA() (r, g, b, a uint32) {
	return hsv.toRGB().RGBA()
}

// HSL is an alias of RGB.HSL
func (hsv HSV) HSL() (int, int, int) {
	return hsv.toRGB().HSL()
}

// HSV is an alias of RGB.HSV
func (hsv HSV) HSV() (int, int, int) {
	return hsv.H, hsv.S, hsv.V
}

// CMYK is an alias of RGB.CMYK
func (hsv HSV) CMYK() (int, int, int, int) {
	return hsv.toRGB().CMYK()
}

// XYZ is an alias of RGB.XYZ
func (hsv HSV) XYZ() (float64, float64, float64) {
	return hsv.toRGB().XYZ()
}

// CIELAB is an alias of RGB.CIELAB
func (hsv HSV) CIELAB() (float64, float64, float64) {
	return hsv.toRGB().CIELAB()
}

// Hex is an alias of RGB.HEX
func (hsv HSV) Hex() string {
	return hsv.toRGB().Hex()
}

// ColorName is an alias of RGB.ColorName
func (hsv HSV) ColorName(colors ColorTable) ColorSpace {
	return hsv.toRGB().ColorName(colors)
}

// Distance calculates the distance between colors in CIELAB using
// simple Euclidean distance.
func (hsv HSV) Distance(c Color) float64 {
	l1, a1, b1 := hsv.CIELAB()
	l2, a2, b2 := c.CIELAB()
	return math.Sqrt(
		math.Pow(l1-l2, 2) +
			math.Pow(a1-a2, 2) +
			math.Pow(b1-b2, 2))
}
