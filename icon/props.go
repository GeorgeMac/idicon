package icon

import (
	"fmt"

	"github.com/GeorgeMac/idicon/colour"
)

// DefaultProps returns an example set of properties
// for use in a Generator.
func DefaultProps() Props {
	return Props{
		BaseColour:   colour.NewColour(0xc1, 0xc1, 0xc1),
		Palette:      colour.Default,
		Size:         50,
		Padding:      5,
		BorderRadius: 5,
	}
}

// Props structure contains SVG formatting properties
// for generator identicons via an icon.Generator.
type Props struct {
	BaseColour   *colour.Colour
	Palette      colour.Palette
	Size         int
	Padding      int
	BorderRadius int
}

// Style returns a valid CSS style string for use
// within style attributes on an SVG element.
func (p Props) Style() string {
	return fmt.Sprintf("background:%s;padding:%dpx;border-radius:%dpx;", p.BaseColour, p.Padding, p.BorderRadius)
}
