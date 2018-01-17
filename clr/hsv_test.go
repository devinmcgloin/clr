package clr

import "testing"

func TestConversionToRGB(t *testing.T) {
	table := map[HSV]RGB{
		HSV{13, 81, 47}:  RGB{119, 43, 23},
		HSV{159, 51, 84}: RGB{104, 214, 175},
		HSV{222, 92, 87}: RGB{17, 78, 221},
		HSV{222, 74, 47}: RGB{31, 57, 119},
		HSV{138, 74, 47}: RGB{31, 119, 57},
		HSV{207, 79, 84}: RGB{44, 138, 214},
	}

	for hsv, rgb := range table {
		if !colorEquality(hsv.toRGB(), rgb) {
			t.Errorf("Expected: %+v, Got: %+v\n", rgb, hsv.toRGB())
		}
	}
}

func colorEquality(a RGB, b RGB) bool {
	return abs(a.R, b.R) <= 1 && abs(a.G, b.G) <= 1 && abs(a.B, b.B) <= 1
}

func abs(a uint8, b uint8) int {
	c := int(a) - int(b)
	if c < 0 {
		c *= -1
	}
	return c
}
