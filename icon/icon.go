package icon

import (
	"bytes"
	"crypto/sha1"
	"fmt"
	"hash"
	"image/color"
	"math"

	"github.com/GeorgeMac/idicon/colour"
)

type Icon struct {
	Data        [][]bool
	BaseColour  color.Color
	ComplColour color.Color
	svgwidth    int
}

func (icn *Icon) String() string {
	buf := &bytes.Buffer{}
	fmt.Fprintf(buf, "Colours [%v, %v] \n", icn.BaseColour, icn.ComplColour)

	for i := 0; i < len(icn.Data); i++ {
		line := make([]byte, len(icn.Data[i])+1)
		for j := 0; j < len(icn.Data[i]); j++ {
			line[j] = token(icn.Data[i][j])
		}
		line[len(line)-1] = '\n'
		buf.Write(line)
	}

	return buf.String()
}

func (icn *Icon) Colours() (b, c string) {
	br, bg, bb, _ := icn.BaseColour.RGBA()
	cr, cg, cb, _ := icn.ComplColour.RGBA()
	return fmt.Sprintf("%x%x%x", uint8(br), uint8(bg), uint8(bb)), fmt.Sprintf("%x%x%x", uint8(cr), uint8(cg), uint8(cb))
}

func (icn *Icon) Svg() string {
	buf := &bytes.Buffer{}
	fmt.Fprintf(buf, `<svg width="%d" height="%d">`, len(icn.Data)*icn.svgwidth, len(icn.Data[0])*icn.svgwidth)

	rectstr := `<rect x="%[1]d" y="%[2]d" width="%[3]d" height="%[3]d" style="fill:#%[4]s"></rect>`
	basecol, complcol := icn.Colours()

	for i := 0; i < len(icn.Data); i++ {
		for j := 0; j < len(icn.Data[i]); j++ {
			colour := complcol
			if icn.Data[i][j] {
				colour = basecol
			}
			fmt.Fprintf(buf, rectstr, j*icn.svgwidth, i*icn.svgwidth, icn.svgwidth, colour)
		}
	}

	fmt.Fprintf(buf, `</svg>`)
	return buf.String()
}

func token(b bool) byte {
	if b {
		return '+'
	}
	return '-'
}

type Generator struct {
	hashfunc              func() hash.Hash
	distr                 distribution
	width, height, hwidth int
	svgwidth              int
}

func NewGenerator(width, height int, opts ...option) (*Generator, error) {
	hwidth := int(math.Ceil(float64(width) / 2.0))
	g := &Generator{
		hashfunc: sha1.New,
		width:    width,
		height:   height,
		hwidth:   hwidth,
		svgwidth: 50,
		distr:    gauss,
	}

	for _, opt := range opts {
		if err := opt(g); err != nil {
			return nil, err
		}
	}

	return g, nil
}

func (g *Generator) Generate(data []byte) *Icon {
	icon := &Icon{
		svgwidth: g.svgwidth,
	}

	// hash input data
	hashed := g.hashfunc().Sum(data)

	// use first three bytes as colour data
	col := color.RGBA{uint8(hashed[0]), uint8(hashed[1]), uint8(hashed[2]), 0xff}
	idx := colour.Base.Index(col)
	icon.BaseColour, icon.ComplColour = colour.Base[idx], colour.Complements[idx]

	res := make([]bool, g.width*g.height)
	// calculate index from distribution and flip corresponding bit
	for _, d := range hashed[3:] {
		f := g.distr(0, 255, float64(g.hwidth*g.height))
		idx := int(math.Floor(f(float64(d))))
		res[idx] = !res[idx]
	}

	icon.Data = g.expand(res)
	return icon
}

func (g *Generator) expand(b []bool) (data [][]bool) {
	data = make([][]bool, g.height)

	odd := g.width%2 != 0
	lower, upper := 0, g.hwidth
	for x := 0; x < g.height; x++ {
		data[x] = make([]bool, g.width)
		j := 0

		seg := b[lower:upper]
		for y := 0; y < len(seg); y++ {
			data[x][j] = seg[y]
			j++
		}
		if odd {
			seg = seg[0 : len(seg)-1]
		}
		for y := len(seg) - 1; y >= 0; y-- {
			data[x][j] = seg[y]
			j++
		}

		lower = upper
		upper += g.hwidth
	}

	return
}

// distribution funcs
type funcx func(x float64) float64

type distribution func(lower, upper, scale float64) funcx

func simple(lower, upper, scale float64) funcx {
	return func(x float64) float64 {
		return (x / (upper - lower)) * scale
	}
}

func gauss(lower, upper, scale float64) funcx {
	mean := lower + ((upper - lower) / 2.0)
	stdv := mean / 3.0
	dist := normal(mean, stdv)
	nscale := scale / dist(mean)
	return func(x float64) float64 {
		return dist(x) * nscale
	}
}

func normal(mean, stdv float64) funcx {
	return func(x float64) float64 {
		expn := math.Pow(x-mean, 2.0) / (2.0 * math.Pow(stdv, 2.0))
		return (1.0 / (stdv * (math.Sqrt(2.0 * math.Pi))) * math.Exp(-1.0*expn))
	}
}
