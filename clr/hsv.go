package clr

import "math"

type HSV struct {
	H uint16
	S uint8
	V uint8
}

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

func (hsv HSV) RGB() (uint8, uint8, uint8) {
	rgb := hsv.toRGB()
	return rgb.R, rgb.G, rgb.B
}

func (hsv HSV) RGBA() (r, g, b, a uint8) {
	return hsv.toRGB().RGBA()
}

func (hsv HSV) HSL() (uint16, uint8, uint8) {
	return hsv.toRGB().HSL()
}

func (hsv HSV) HSV() (uint16, uint8, uint8) {
	return hsv.H, hsv.S, hsv.V
}

func (hsv HSV) CMYK() (uint8, uint8, uint8, uint8) {
	return hsv.toRGB().CMYK()
}

func (hsv HSV) XYZ() (float64, float64, float64) {
	return hsv.toRGB().XYZ()
}

func (hsv HSV) CIELAB() (float64, float64, float64) {
	return hsv.toRGB().CIELAB()
}

func (hsv HSV) Hex() string {
	return hsv.toRGB().Hex()
}
func (hsv HSV) ColorName(colors ColorTable) ColorSpace {
	return hsv.toRGB().ColorName(colors)
}
