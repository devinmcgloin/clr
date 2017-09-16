package clr

import (
	"log"
	"strconv"
)

type Hex struct {
	Code string `json:"code"`
}

func (c Hex) getRGB() RGB {
	r, g, b := c.RGB()
	return RGB{R: r, G: g, B: b}
}

func (c Hex) Valid() bool {
	var err error
	_, err = strconv.ParseUint(c.Code[0:2], 16, 8)
	if err != nil {
		return false
	}
	_, err = strconv.ParseUint(c.Code[2:4], 16, 8)
	if err != nil {
		return false
	}
	_, err = strconv.ParseUint(c.Code[4:6], 16, 8)
	if err != nil {
		return false
	}
	return true
}

func (c Hex) RGB() (uint8, uint8, uint8) {
	r, err := strconv.ParseUint(c.Code[0:2], 16, 8)
	if err != nil {
		log.Fatal(err)
	}
	g, err := strconv.ParseUint(c.Code[2:4], 16, 8)
	if err != nil {
		log.Fatal(err)
	}
	b, err := strconv.ParseUint(c.Code[4:6], 16, 8)
	if err != nil {
		log.Fatal(err)
	}

	return uint8(r), uint8(g), uint8(b)
}

func (c Hex) RGBA() (r, g, b, a uint8) {
	return c.getRGB().RGBA()
}

func (c Hex) HSL() (uint16, uint8, uint8) {
	return c.getRGB().HSL()
}

func (c Hex) HSV() (uint16, uint8, uint8) {
	return c.getRGB().HSV()
}

func (c Hex) CMYK() (uint8, uint8, uint8, uint8) {
	return c.getRGB().CMYK()
}

func (c Hex) Hex() string {
	return c.Code
}

func (c Hex) XYZ() (x, y, z float64) {
	return c.getRGB().XYZ()
}

func (c Hex) CIELAB() (l, a, b float64) {
	return c.getRGB().CIELAB()
}

func (c Hex) ColorName(colors ColorTable) ColorSpace {
	return c.getRGB().ColorName(colors)
}

func (hex Hex) Distance(c Color) float64 {
	return hex.getRGB().Distance(c)
}
