package icon

import (
	"bytes"
	"fmt"

	"github.com/GeorgeMac/idicon/colour"
)

const (
	SVG_OPEN  string = `<svg width="%d" height="%d" style="%s">`
	SVG_CLOSE string = `</svg>`
	RECT_STR  string = `<rect x="%[1]d" y="%[2]d" width="%[3]d" height="%[3]d" style="fill:%[4]s"></rect>`
)

type Icon struct {
	Data  [][]bool
	base  *colour.Colour
	main  *colour.Colour
	props Props
}

func (icn *Icon) String() string {
	buf := &bytes.Buffer{}

	var (
		size = icn.props.Size
	)

	fmt.Fprintf(buf, SVG_OPEN, len(icn.Data)*size, len(icn.Data[0])*size, icn.props.Style())

	for i := 0; i < len(icn.Data); i++ {
		for j := 0; j < len(icn.Data[i]); j++ {
			colour := icn.base
			if !icn.Data[i][j] {
				colour = icn.main
			}
			fmt.Fprintf(buf, RECT_STR, j*size, i*size, size, colour)
		}
	}

	fmt.Fprintf(buf, SVG_CLOSE)
	return buf.String()
}
