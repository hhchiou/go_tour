package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r rot13Reader) Read(b []byte) (int, error) {
	n, e := r.r.Read(b)
	for i, j := range b {
		switch {
		case ('A' <= j && j <= 'M') || ('a' <= j && j <= 'm'):
			b[i] += 13
		case ('N' <= j && j <= 'Z') || ('n' <= j && j <= 'z'):
			b[i] -= 13
		default:
		}
	}
	return n, e
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
