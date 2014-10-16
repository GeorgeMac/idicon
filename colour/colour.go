package colour

import (
	"fmt"
	"image/color"
)

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
	return fmt.Sprintf("#%x%x%x", uint8(r), uint8(g), uint8(b))
}

var Base = color.Palette([]color.Color{
	color.RGBA{0x6e, 0xa1, 0xff, 0xff},
	color.RGBA{0xf2, 0x80, 0x74, 0xff},
	color.RGBA{0xf7, 0xb5, 0x6b, 0xff},
	color.RGBA{0xfa, 0xeb, 0x78, 0xff},
})

var Complements = color.Palette([]color.Color{
	color.RGBA{0xff, 0xb1, 0x71, 0xff},
	color.RGBA{0x76, 0xff, 0xc4, 0xff},
	color.RGBA{0x81, 0x6b, 0xff, 0xff},
	color.RGBA{0xed, 0x76, 0xff, 0xff},
})
