package clr

import "math"

// HSV represents the HSV Colorspace.
type HSV struct {
	H uint16 `json:"h"`
	S uint8  `json:"s"`
	V uint8  `json:"v"`
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

	return RGB{R: uint8(r * 255), G: uint8(g * 255), B: uint8(b * 255)}
}

// RGB Converts HSV to RGB using toRGB
func (hsv HSV) RGB() (uint8, uint8, uint8) {
	rgb := hsv.toRGB()
	return rgb.R, rgb.G, rgb.B
}

// RGBA is an alias of RGB.RGBA
func (hsv HSV) RGBA() (r, g, b, a uint8) {
	return hsv.toRGB().RGBA()
}

// HSL is an alias of RGB.HSL
func (hsv HSV) HSL() (uint16, uint8, uint8) {
	return hsv.toRGB().HSL()
}

// HSV is an alias of RGB.HSV
func (hsv HSV) HSV() (uint16, uint8, uint8) {
	return hsv.H, hsv.S, hsv.V
}

// CMYK is an alias of RGB.CMYK
func (hsv HSV) CMYK() (uint8, uint8, uint8, uint8) {
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
