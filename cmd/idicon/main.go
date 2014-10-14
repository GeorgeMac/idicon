package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/GeorgeMac/idicon/icon"
)

func main() {
	funcn := os.Args[1]
	fset := flag.FlagSet{}
	var width, height, svgx int
	fset.IntVar(&width, "w", 6, "Identicon Width")
	fset.IntVar(&height, "h", 6, "Identicon Height")
	fset.IntVar(&svgx, "x", 30, "SVG scale")
	fset.Parse(os.Args[2:])

	arg := []byte(fset.Arg(1))

	generator, err := icon.NewGenerator(width, height, icon.SetWidth(svgx))
	if err != nil {
		panic(err)
	}

	icon := generator.Generate(arg)

	switch funcn {
	case "print":
		fmt.Println(icon)
	case "svg":
		fmt.Println(icon.Svg())
	}
}
