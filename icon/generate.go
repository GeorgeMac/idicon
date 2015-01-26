package icon

import (
	"crypto/sha1"
	"hash"
	"image/color"
	"math"
)

// Generator structure contains the schematics for generating
// identicons from strings.
type Generator struct {
	hashfunc              func() hash.Hash
	distr                 distribution
	props                 Props
	width, height, hwidth int
}

// NewGenerator returns a pointer to a generator or an error
// if there was a problem applying a provided option type.
func NewGenerator(width, height int, opts ...option) (*Generator, error) {
	hwidth := int(math.Ceil(float64(width) / 2.0))
	g := &Generator{
		hashfunc: sha1.New,
		width:    width,
		height:   height,
		hwidth:   hwidth,
		props:    DefaultProps(),
		distr:    normal,
	}

	for _, opt := range opts {
		if err := opt(g); err != nil {
			return nil, err
		}
	}

	return g, nil
}

// Generate returns a pointer to an icon for a
// provided slice of bytes.
func (g *Generator) Generate(data []byte) *Icon {
	icon := &Icon{
		props: g.props,
	}

	// hash input data
	hashed := g.hashfunc().Sum(data)

	// use first three bytes as colour data
	col := color.RGBA{uint8(hashed[0]), uint8(hashed[1]), uint8(hashed[2]), 0xff}
	icon.base, icon.main = g.props.BaseColour, g.props.Palette.Nearest(col)

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

func linear(lower, upper, scale float64) funcx {
	return func(x float64) float64 {
		return (x / (upper - lower)) * scale
	}
}

func normal(lower, upper, scale float64) funcx {
	mean := lower + ((upper - lower) / 2.0)
	stdv := mean / 3.0
	dist := gauss(mean, stdv)
	nscale := scale / dist(mean)
	return func(x float64) float64 {
		return dist(x) * nscale
	}
}

func gauss(mean, stdv float64) funcx {
	return func(x float64) float64 {
		expn := math.Pow(x-mean, 2.0) / (2.0 * math.Pow(stdv, 2.0))
		return (1.0 / (stdv * (math.Sqrt(2.0 * math.Pi))) * math.Exp(-1.0*expn))
	}
}
