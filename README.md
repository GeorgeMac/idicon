Identicons in Go
================

![georgemac](https://raw.githubusercontent.com/GeorgeMac/idicon/gh-pages/images/georgemac.png)
![gobug.me](https://raw.githubusercontent.com/GeorgeMac/idicon/gh-pages/images/gobugme.png)
![incisive.ly](https://raw.githubusercontent.com/GeorgeMac/idicon/gh-pages/images/incisively.png)

Use the following to get the libs:

`go get github.com/GeorgeMac/idicon`

Use the following to get the `idicon` command

`go get github.com/GeorgeMac/idicon/cmd/idicon`

See: [idicon command](https://github.com/GeorgeMac/idicon/tree/master/cmd/idicon)

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
