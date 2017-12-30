package clr

import "testing"

func TestConversionToRGB(t *testing.T) {
	table := map[HSV]RGB{
		HSV{13, 81, 47}:  RGB{119, 43, 23},
		HSV{159, 51, 84}: RGB{104, 214, 175},
	}

	for hsv, rgb := range table {
		if !colorEquality(hsv.toRGB(), rgb) {
			t.Errorf("Expected: %+v, Got: %+v\n", rgb, hsv.toRGB())
		}
	}
}

func colorEquality(a RGB, b RGB) bool {
	return a.R == b.R && a.G == b.G && a.B == b.B
}
