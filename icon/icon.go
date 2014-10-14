package icon

import (
	"bytes"
	"crypto/sha1"
	"fmt"
	"hash"
	"math"
)

type Icon struct {
	Data                  []bool
	Colour                []byte
	Width, Height, HWidth int
}

func (i *Icon) String() string {
	buf := &bytes.Buffer{}
	fmt.Fprintf(buf, "Base Colour [%x] \n", i.Colour)

	odd := i.Width%2 != 0
	lower, upper := 0, i.HWidth
	for x := 0; x < i.Height; x++ {
		seg := i.Data[lower:upper]
		for y := 0; y < len(seg); y++ {
			buf.Write(token(seg[y]))
		}
		if odd {
			seg = seg[0 : len(seg)-1]
		}
		for y := len(seg) - 1; y >= 0; y-- {
			buf.Write(token(seg[y]))
		}

		lower = upper
		upper += i.HWidth
		buf.Write([]byte("\n"))
	}

	return buf.String()
}

func token(b bool) []byte {
	if b {
		return []byte("+")
	}
	return []byte("-")
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
	icon := &Icon{
		Data:   make([]bool, g.width*g.height),
		Colour: make([]byte, 3),
		Width:  g.width,
		Height: g.height,
		HWidth: g.hwidth,
	}

	// hash input data
	hashed := g.hashfunc().Sum(data)

	// use first three bytes as colour data
	copy(icon.Colour, hashed[0:3])

	// calculate index from distribution and flip corresponding bit
	for _, d := range hashed[3:] {
		idx := int(math.Floor(g.distr(float64(d))))
		icon.Data[idx] = !icon.Data[idx]
	}

	return icon
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
