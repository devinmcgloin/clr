package clr

import (
	"testing"
)

func TestHSL(t *testing.T) {
	tables := []struct {
		rgb RGB
		h   uint16
		s   uint8
		l   uint8
	}{
		{RGB{R: 78, G: 91, B: 112}, 217, 17, 37},
		{RGB{R: 165, G: 23, B: 139}, 310, 75, 36},
		{RGB{R: 221, G: 132, B: 90}, 19, 65, 60},
		{RGB{R: 89, G: 1, B: 55}, 323, 97, 17},
		{RGB{R: 132, G: 135, B: 132}, 120, 1, 52},
	}

	for _, table := range tables {
		h, s, l := table.rgb.HSL()
		if h-table.h > 1 || table.h-h > 1 {
			t.Errorf("H of %v was incorrect, got: %d wanted: %d\n", table.rgb, h, table.h)
		}
		if s != table.s {
			t.Errorf("S of %v was incorrect, got: %d wanted: %d\n", table.rgb, s, table.s)
		}
		if l != table.l {
			t.Errorf("L of %v was incorrect, got: %d wanted: %d\n", table.rgb, l, table.l)
		}
	}
}

func TestHSV(t *testing.T) {
	tables := []struct {
		rgb RGB
		h   uint16
		s   uint8
		v   uint8
	}{
		{RGB{R: 78, G: 91, B: 112}, 217, 30, 43},
		{RGB{R: 165, G: 23, B: 139}, 310, 86, 64},
		{RGB{R: 221, G: 132, B: 90}, 19, 59, 86},
		{RGB{R: 89, G: 1, B: 55}, 323, 98, 34},
		{RGB{R: 132, G: 135, B: 132}, 120, 2, 52},
	}

	for _, table := range tables {
		h, s, v := table.rgb.HSV()
		if h-table.h > 1 || table.h-h > 1 {
			t.Errorf("H of %v was incorrect, got: %d wanted: %d\n", table.rgb, h, table.h)
		}
		if s != table.s {
			t.Errorf("S of %v was incorrect, got: %d wanted: %d\n", table.rgb, s, table.s)
		}
		if v != table.v {
			t.Errorf("V of %v was incorrect, got: %d wanted: %d\n", table.rgb, v, table.v)
		}
	}
}

func BenchmarkColorDistance(b *testing.B) {
	colors := []RGB{
		RGB{R: 78, G: 91, B: 112},
		RGB{R: 165, G: 23, B: 139},
		RGB{R: 221, G: 132, B: 90},
		RGB{R: 89, G: 1, B: 55},
		RGB{R: 132, G: 135, B: 132},
	}
	for i := 0; i < b.N; i++ {
		for _, c1 := range colors {
			for _, c2 := range colors {
				c1.Distance(c2)
			}
		}
	}
}
