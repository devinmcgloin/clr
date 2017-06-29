package main

import (
	"flag"
	"fmt"

	"github.com/devinmcgloin/clr/clr"
)

func main() {

	var hex string
	flag.StringVar(&hex, "hex", "#000000", "Color represented in Hex")
	flag.Parse()
	color := &clr.Hex{
		Code: hex[1:7],
	}

	fmt.Printf("%+v\n", color)
	fmt.Print("sRGB: ")
	fmt.Println(color.RGB())
	fmt.Print("HSL: ")
	fmt.Println(color.HSL())
	fmt.Print("HSV: ")
	fmt.Println(color.HSV())
	fmt.Print("CMYK: ")
	fmt.Println(color.CMYK())
	fmt.Println("HEX:", color.Hex())
}
