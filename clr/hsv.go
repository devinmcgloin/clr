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
	rgb := RGB{}
	h := float64(hsv.H) / 260.0
	s := float64(hsv.S) / 100.0
	v := float64(hsv.V) / 100.0

	if s == 0 {
		rgb.R = uint8(v * 255)
		rgb.G = uint8(v * 255)
		rgb.B = uint8(v * 255)
	} else {
		h /= 6.0

		if h == 6 {
			h = 0
		}

		i := math.Floor(h)

		a := v * (1 - s)
		b := v * (1 - s*(h-i))
		c := v * (1 - s*(1-(h-i)))

		var red, green, blue float64

		switch i {
		case 0:
			red = v
			green = c
			blue = a
		case 1:
			red = b
			green = v
			blue = a
		case 2:
			red = a
			green = v
			blue = c
		case 3:
			red = a
			green = b
			blue = v
		case 4:
			red = c
			green = a
			blue = b
		default:
			red = v
			green = a
			blue = b
		}
		rgb.R = uint8(red * 255)
		rgb.G = uint8(green * 255)
		rgb.B = uint8(blue * 255)
	}
	return rgb
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
