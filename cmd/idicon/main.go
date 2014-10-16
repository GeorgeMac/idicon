package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"

	"github.com/GeorgeMac/idicon/icon"
)

var flags = flag.FlagSet{
	Usage: func() {
		fmt.Fprintf(os.Stderr, `Usage: idicon [print | svg] [-h "height" | -w "width" | -x "svg size"] <input string>\n`)
	},
}

func Print(icn *icon.Icon) string {
	buf := &bytes.Buffer{}
	for i := 0; i < len(icn.Data); i++ {
		line := make([]byte, len(icn.Data[i])+1)
		for j := 0; j < len(icn.Data[i]); j++ {
			line[j] = token(icn.Data[i][j])
		}
		line[len(line)-1] = '\n'
		buf.Write(line)
	}
	return buf.String()
}

func token(b bool) byte {
	if b {
		return '+'
	}
	return '-'
}

func main() {
	var width, height, svgx int
	flags.IntVar(&width, "w", 5, "Identicon Width")
	flags.IntVar(&height, "h", 5, "Identicon Height")
	flags.IntVar(&svgx, "x", 10, "SVG scale")

	if len(os.Args) < 3 {
		flags.Usage()
		return
	}

	funcn := os.Args[1]
	flags.Parse(os.Args[2:])
	arg := []byte(flags.Arg(0))

	props := icon.DefaultProps()
	props.Size = svgx

	generator, err := icon.NewGenerator(width, height, icon.With(props))
	if err != nil {
		panic(err)
	}

	icon := generator.Generate(arg)

	switch funcn {
	case "print":
		fmt.Print(Print(icon))
	case "svg":
		fmt.Println(icon)
	case "html":
		fmt.Printf("<html><body>%s</body></html>", icon)
	}
}
