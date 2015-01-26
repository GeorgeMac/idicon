package icon

import (
	"bytes"
	"fmt"

	"github.com/GeorgeMac/idicon/colour"
)

const (
	openTag  string = `<svg width="%d" height="%d" style="%s">`
	closeTag string = `</svg>`
	rectFmt  string = `<rect x="%[1]d" y="%[2]d" width="%[3]d" height="%[3]d" style="fill:%[4]s"></rect>`
)

// Icon holds a representation of a generated identicon
// and can produce SVG string representation.
type Icon struct {
	Data  [][]bool
	base  *colour.Colour
	main  *colour.Colour
	props Props
}

// String returns the SVG representation of the
// underlying identicon.
func (icn *Icon) String() string {
	buf := &bytes.Buffer{}

	var (
		size = icn.props.Size
	)

	fmt.Fprintf(buf, openTag, len(icn.Data)*size, len(icn.Data[0])*size, icn.props.Style())

	for i := 0; i < len(icn.Data); i++ {
		for j := 0; j < len(icn.Data[i]); j++ {
			colour := icn.base
			if !icn.Data[i][j] {
				colour = icn.main
			}
			fmt.Fprintf(buf, rectFmt, j*size, i*size, size, colour)
		}
	}

	fmt.Fprintf(buf, closeTag)
	return buf.String()
}
