package main

import (
	"flag"
	"fmt"

	"github.com/GeorgeMac/idicon/icon"
)

func main() {
	var width, height int
	flag.IntVar(&width, "w", 6, "Identicon Width")
	flag.IntVar(&height, "h", 6, "Identicon Height")
	flag.Parse()

	arg := []byte(flag.Arg(0))

	generator, err := icon.NewGenerator(width, height)
	if err != nil {
		panic(err)
	}

	fmt.Println(generator.Generate(arg))
}
