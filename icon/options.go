package icon

import "hash"

type option func(*Generator) error

// UseHash is used to provide an icon.Generator
// with an alternative hashing function to the
// default (sha1).
// e.g. NewGenerator(5, 5, UseHash(md5.New))
func UseHash(h func() hash.Hash) option {
	return func(g *Generator) error {
		g.hashfunc = h
		return nil
	}
}

// LinearDistribution is used to overide the
// normalisation distribution used to produce
// indexes when flipping bools in the pattern.
// By default idicon uses a normal distribution.
// This function swaps it for a linear one.
func LinearDistribution(g *Generator) error {
	g.distr = simple
	return nil
}

func SvgSize(w int) option {
	return func(g *Generator) error {
		g.svgwidth = w
		return nil
	}
}
