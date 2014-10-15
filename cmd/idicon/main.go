package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/GeorgeMac/idicon/icon"
)

var flags = flag.FlagSet{
	Usage: func() {
		fmt.Fprintf(os.Stderr, `Usage: idicon [print | svg] [-h "height" | -w "width" | -x "svg size"] <input string>`)
	},
}

func main() {
	var width, height, svgx int
	flags.IntVar(&width, "w", 6, "Identicon Width")
	flags.IntVar(&height, "h", 6, "Identicon Height")
	flags.IntVar(&svgx, "x", 30, "SVG scale")

	if len(os.Args) < 3 {
		flags.Usage()
		return
	}

	funcn := os.Args[1]
	flags.Parse(os.Args[2:])
	arg := []byte(flags.Arg(0))

	generator, err := icon.NewGenerator(width, height, icon.SvgSize(svgx))
	if err != nil {
		panic(err)
	}

	icon := generator.Generate(arg)

	switch funcn {
	case "print":
		fmt.Println(icon)
	case "svg":
		fmt.Println(icon.Svg())
	case "html":
		fmt.Printf("<html><body>%s</body></html>", icon.Svg())
	}
}
