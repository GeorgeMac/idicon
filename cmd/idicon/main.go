package main

import (
	"fmt"
	"os"

	"github.com/GeorgeMac/idicon/icon"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage idicon <name>")
		return
	}

	arg := []byte(os.Args[1])

	generator, err := icon.NewGenerator(7, 6)
	if err != nil {
		panic(err)
	}

	fmt.Println(generator.Generate(arg))
}
