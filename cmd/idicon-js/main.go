package main

import (
	"github.com/GeorgeMac/idicon/icon"
	"github.com/gopherjs/gopherjs/js"
)

func Generate(g *icon.Generator) func(s string) string {
	return func(s string) string {
		return g.Generate([]byte(s)).String()
	}
}

func main() {
	props := icon.DefaultProps()
	props.Padding = 8
	props.Size = 30

	generator, err := icon.NewGenerator(8, 8, icon.With(props))
	if err != nil {
		panic(err)
	}

	js.Global.Set("icon", map[string]interface{}{
		"Generate": Generate(generator),
	})
}
