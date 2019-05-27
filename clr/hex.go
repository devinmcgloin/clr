package clr

import (
	"log"
	"strconv"
)

// Hex represents the Hex format
type Hex struct {
	Code string `json:"code"`
}

// getRGB provides the hatch into all the defined RGB Methods.
// This allows for lossless conversion to all other supported
// colorspaces.
func (c Hex) getRGB() RGB {
	r, g, b := c.RGB()
	return RGB{R: r, G: g, B: b}
}

// Valid checks if a given hex code is valid. Currently only
// 6 digit hex codes are supported.
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
	return err == nil
}

// RGB converts to rgb.
func (c Hex) RGB() (int, int, int) {
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

	return int(r), int(g), int(b)
}

// RGBA is an alias of RGB.RGBA
func (c Hex) RGBA() (r, g, b, a uint32) {
	return c.getRGB().RGBA()
}

// HSL is an alias of RGB.HSL
func (c Hex) HSL() (int, int, int) {
	return c.getRGB().HSL()
}

// HSV is an alias of RGB.HSV
func (c Hex) HSV() (int, int, int) {
	return c.getRGB().HSV()
}

// CMYK is an alias of RGB.CMYK
func (c Hex) CMYK() (int, int, int, int) {
	return c.getRGB().CMYK()
}

// Hex returns the hexcode in the Hex struct.
func (c Hex) Hex() string {
	return c.Code
}

// XYZ is an alias of RGB.XYZ
func (c Hex) XYZ() (x, y, z float64) {
	return c.getRGB().XYZ()
}

// CIELAB is an alias of RGB.CIELAB
func (c Hex) CIELAB() (l, a, b float64) {
	return c.getRGB().CIELAB()
}

// ColorName is an alias of RGB.ColorName
func (c Hex) ColorName(colors ColorTable) ColorSpace {
	return c.getRGB().ColorName(colors)
}

// Distance is an alias of RGB.Distance
func (c Hex) Distance(otherC Color) float64 {
	return c.getRGB().Distance(otherC)
}
