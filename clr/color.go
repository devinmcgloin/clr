package clr

// Color is the primary interface this library defines. Hex, HSL and RGB Types all satisfy this interface.
type Color interface {
	Valid() bool
	RGB() (int, int, int)
	HSL() (int, int, int)
	HSV() (int, int, int)
	CMYK() (int, int, int, int)
	XYZ() (float64, float64, float64)
	CIELAB() (l, a, b float64)
	Hex() string
	ColorName(colors ColorTable) ColorSpace
	Distance(c Color) float64
	RGBA() (r, g, b, a uint32)
}

// ColorTable allows colors to be mapped to names, this must be fulfilled to the library as there is no built in ColorTable.
type ColorTable interface {
	Iterate() []Color
	Lookup(hexCode string) ColorSpace
}

// ColorSpace is a alias for the name of a region of color
type ColorSpace string
