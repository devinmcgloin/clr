# clr

clr is a go library to manage different color spaces, convert between them and
compare colors.

clr currently supports sRGB HSV, and Hex types, with conversions to sRGB, Hex, HSV,
HSL, CMYK, XYZ and CIE-Lab for each.

Comparisons are done in CIE-Lab using Delta E*. I may add more comparison
options in the future, but for now Delta E* performs well enough.

Given a mapping from Hex to color names you can find the nearest color visually
in the mapping from any given color. Data for these mappings is up to you to
provide by implementing `ColorTable`.

There is also a CLI tool to explore the package called `clr`.

```
$ clr --hex "#0c6624"
&{Code:0c6624}
sRGB: 12 102 36
HSL: 136 78 22
HSV: 136 88 40
CMYK: 88 0 64 60
HEX: 0c6624
Generic Color: Green
Specific Color: Camarone
```

## API

clr is based on two primitives `Color` and `ColorTable`. `Color` has three
backing types currently, Hex, RGB, and HSV. You can add your own types by
implementing the `Color` interface.

```golang
type Color interface {
	RGB() (uint8, uint8, uint8)
	HSL() (uint16, uint8, uint8)
	HSV() (uint16, uint8, uint8)
	CMYK() (uint8, uint8, uint8, uint8)
	XYZ() (float64, float64, float64)
	CIELAB() (l, a, b float64)
	Hex() string
	ColorName(colors ColorTable) ColorSpace
	Distance(c Color) float64
}

type ColorTable interface {
	Iterate() []Color
	Lookup(hexCode string) ColorSpace
    }
```
