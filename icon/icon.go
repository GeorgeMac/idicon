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
}

func (icn *Icon) String() string {
	buf := &bytes.Buffer{}
	fmt.Fprintf(buf, "Colours [%v, %v] \n", icn.BaseColour, icn.ComplColour)

	for i := 0; i < len(icn.Data); i++ {
		line := make([]byte, len(icn.Data[i])+1)
		for j := 0; j < len(icn.Data[i]); j++ {
			line[j] = token(icn.Data[i][j])
		}
		line[len(line)-1] = byte('\n')
		buf.Write(line)
	}

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
}

func NewGenerator(width, height int, opts ...option) (*Generator, error) {
	hwidth := int(math.Ceil(float64(width) / 2.0))
	g := &Generator{
		hashfunc: sha1.New,
		width:    width,
		height:   height,
		hwidth:   hwidth,
		distr:    simple(0, 255, float64(hwidth*height)),
	}

	for _, opt := range opts {
		if err := opt(g); err != nil {
			return nil, err
		}
	}

	return g, nil
}

func (g *Generator) Generate(data []byte) *Icon {
	icon := &Icon{}

	// hash input data
	hashed := g.hashfunc().Sum(data)

	// use first three bytes as colour data
	col := color.RGBA{uint8(hashed[0]), uint8(hashed[1]), uint8(hashed[2]), 0xff}
	idx := colour.Base.Index(col)
	icon.BaseColour, icon.ComplColour = colour.Base[idx], colour.Complements[idx]

	res := make([]bool, g.width*g.height)
	// calculate index from distribution and flip corresponding bit
	for _, d := range hashed[3:] {
		idx := int(math.Floor(g.distr(float64(d))))
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

type distribution func(x float64) float64

func normal(mean, stdv float64) distribution {
	return func(x float64) float64 {
		expn := math.Pow(x-mean, 2.0) / (2.0 * math.Pow(stdv, 2.0))
		return (1.0 / (stdv * (math.Sqrt(2.0 * math.Pi))) * math.Exp(-1.0*expn))
	}
}

func simple(lower, upper, factor float64) distribution {
	mean := lower + ((upper - lower) / 2.0)
	stdv := mean / 3.0
	dist := normal(mean, stdv)
	scale := factor / dist(mean)
	return func(x float64) float64 {
		return dist(x) * scale
	}
}
