package clr

import (
	"log"

	"github.com/BurntSushi/toml"
)

var genericColors colors
var specificColors colors

func init() {

	if _, err := toml.DecodeFile("./clr/colors.toml", &specificColors); err != nil {
		log.Panic(err)
	}

	if _, err := toml.DecodeFile("./clr/main-colors.toml", &genericColors); err != nil {
		log.Panic(err)
	}

}

type Color interface {
	RGB() (uint8, uint8, uint8)
	HSL() (uint16, uint8, uint8)
	HSV() (uint16, uint8, uint8)
	CMYK() (uint8, uint8, uint8, uint8)
	XYZ() (float64, float64, float64)
	CIELAB() (l, a, b float64)
	Hex() string
	GenericColorSpace() ColorSpace
	SpecificColorSpace() ColorSpace
	Distance(c Color) float64
}

type ColorSpace string

type color struct {
	Hex  string `toml:"hex"`
	Name string `toml:"name"`
}

type colors struct {
	Color []color
}
