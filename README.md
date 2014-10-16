[Identicons](http://en.wikipedia.org/wiki/Identicon) in Go
================

![idicon](https://raw.githubusercontent.com/GeorgeMac/idicon/gh-pages/images/idicons.png)

### About

> After checking with the official docker golang container releases, I can say with some confidence `idicon` works with golang versions 1.2+.

`go get github.com/GeorgeMac/idicon/cmd/idicon`

See: [idicon command](https://github.com/GeorgeMac/idicon/tree/master/cmd/idicon) for command usage.

### lib usage

Basically this:
```go
generator, err := icon.NewGenerator(5, 5)
if err != nil {
    // handle error
}

// use the generator for a given string input
icn := generator.Generate(`GeorgeMac`)

// print svg string representation
fmt.Print(icn)
```

### `icon.NewGenerator(...)` Usage

NewGenerator produces a idicon icon generator struct for producing `*icon.Icon`
It takes a width, height and an variadic set of options.

```go
type option func(g *Generator) error

func NewGenerator(width, height int, options ...option) *Generator { ... }
```

Current available options include:

```go
icon.UseHash(h hash.Hash) // Defaults to crypto.Sha1.New
icon.LinearDistribution // Different way of indexing flipped bits
icon.With(props icon.Props) // SVG output properties
```

### `icon.Props{}` Usage

```go
type Props struct {
	BaseColour   *colour.Colour
	Palette      colour.Palette
	Size         int
	Padding      int
	BorderRadius int
}
```

1. BaseColour - background colour
2. Palette - colour palette for deriving svg square colours
3. Size - width/height of an individual svg square
4. Padding - padding around the svg grid of squares
5. BorderRadius - for those pretty curved corners
