package clr

type Color interface {
	Valid() bool
	RGB() (uint8, uint8, uint8)
	HSL() (uint16, uint8, uint8)
	HSV() (uint16, uint8, uint8)
	CMYK() (uint8, uint8, uint8, uint8)
	XYZ() (float64, float64, float64)
	CIELAB() (l, a, b float64)
	Hex() string
	ColorName(colors ColorTable) ColorSpace
	Distance(c Color) float64
	RGBA() (r, g, b, a uint8)
}

type ColorTable interface {
	Iterate() []Color
	Lookup(hexCode string) ColorSpace
}

type ColorSpace string
