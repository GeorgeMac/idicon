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
	var width, height int
	fset.IntVar(&width, "w", 6, "Identicon Width")
	fset.IntVar(&height, "h", 6, "Identicon Height")
	fset.Parse(os.Args[2:])

	arg := []byte(fset.Arg(1))

	generator, err := icon.NewGenerator(width, height)
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
