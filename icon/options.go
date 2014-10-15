package icon

import (
	"crypto/md5"
	"crypto/sha1"
)

type option func(*Generator) error

func UseSha1(g *Generator) error {
	g.hashfunc = sha1.New
	return nil
}

func UseMd5(g *Generator) error {
	g.hashfunc = md5.New
	return nil
}

func SvgSize(w int) option {
	return func(g *Generator) error {
		g.svgwidth = w
		return nil
	}
}
