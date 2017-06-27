package main

import (
	"flag"
	"fmt"

	"github.com/sprioc/clr/clr"
)

func main() {
	var hex string
	flag.StringVar(&hex, "hex", "#000000", "Color represented in Hex")
	flag.Parse()
	rgb := &clr.Hex{
		Code: hex[1:7],
	}

	fmt.Printf("%+v\n", rgb)
	fmt.Print("HSL: ")
	fmt.Println(rgb.HSL())
	fmt.Print("HSV: ")
	fmt.Println(rgb.HSV())
	fmt.Print("CMYK: ")
	fmt.Println(rgb.CMYK())
	fmt.Println("HEX:", rgb.Hex())
	fmt.Println("Generic Color:", rgb.GenericColorSpace())
	fmt.Println("Specific Color:", rgb.SpecificColorSpace())
}
