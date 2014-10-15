[Identicons](http://en.wikipedia.org/wiki/Identicon) in Go
================

![georgemac](https://raw.githubusercontent.com/GeorgeMac/idicon/gh-pages/images/georgemac.png)
![gobug.me](https://raw.githubusercontent.com/GeorgeMac/idicon/gh-pages/images/gobugme.png)
![incisive.ly](https://raw.githubusercontent.com/GeorgeMac/idicon/gh-pages/images/incisively.png)

### About

> After checking with the official docker golang container releases, I can say with some confidence `idicon` works with golang versions 1.2+.

`go get github.com/GeorgeMac/idicon/cmd/idicon`

See: [idicon command](https://github.com/GeorgeMac/idicon/tree/master/cmd/idicon) for command usage.

### lib usage

Basically this:
```go
generator, err := icon.NewGenerator(5, 5, icon.UseMd5, icon.SvgSize(10))
if err != nil {
    // handle error
}

// use the generator for a given string input
icn := generator.Generate(`GeorgeMac`)

// string representation
fmt.Print(icn)

// svg string representation
fmt.Print(icn.Svg())
```

### NewGenerator with variadic options

```go
type option func(g *Generator) error

func NewGenerator(width, height int, ...option) *Generator { ... }
```

Current available options include:

```go
icon.UseSha1 // pointless because it does this by default but meh
icon.UseMd5 // use md5 hash function for generating identicons
icon.SvgSize(size int) // set the width/height of the outputted svg squares
```

Future options, currently **NOT** available:
```go
icon.SetPalette(basePalette, complPalette)
icon.Use* // any hash functions the community desire?
icon.UseHash(func() hash.Hash) // user supplied hash generator function
```
