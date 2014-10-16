package icon

import (
	"fmt"

	"github.com/GeorgeMac/idicon/colour"
)

func DefaultProps() Props {
	return Props{
		BaseColour:   colour.NewColour(0xc1, 0xc1, 0xc1),
		Palette:      colour.Default,
		Size:         50,
		Padding:      5,
		BorderRadius: 5,
	}
}

type Props struct {
	BaseColour   *colour.Colour
	Palette      colour.Palette
	Size         int
	Padding      int
	BorderRadius int
}

func (p Props) Style() string {
	return fmt.Sprintf("background:%s;padding:%dpx;border-radius:%dpx;", p.BaseColour, p.Padding, p.BorderRadius)
}
