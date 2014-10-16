package colour

import (
	"fmt"
	"image/color"
)

// Palette wraps a color.Palette
// It allows for ease of retrieving a
// colour.Colour using color.Palette.Convert
type Palette struct {
	color.Palette
}

func (p Palette) Nearest(c color.Color) *Colour {
	return &Colour{p.Palette.Convert(c)}
}

// Colour implements color.Color but
// can be printed in a web friendly
// hex value.
type Colour struct {
	color.Color
}

func NewColour(r, g, b uint8) *Colour {
	return &Colour{color.RGBA{r, g, b, 0xff}}
}

func (c *Colour) RGBA() (r, g, b, a uint32) {
	return c.Color.RGBA()
}

func (c *Colour) String() string {
	r, g, b, _ := c.RGBA()
	return fmt.Sprintf("#%s%s%s", hex(r), hex(g), hex(b))
}

func hex(u uint32) (s string) {
	v := uint8(u)
	if v < 0x10 {
		s += "0"
	}
	return fmt.Sprintf("%s%x", s, v)
}

var Default = Palette{
	color.Palette{
		NewColour(0x6e, 0xa1, 0xff),
		NewColour(0xf2, 0x80, 0x74),
		NewColour(0xf7, 0xb5, 0x6b),
		NewColour(0xfa, 0xeb, 0x78),
	},
}
