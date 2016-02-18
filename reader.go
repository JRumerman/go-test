package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func translate(v byte) byte {
	var diff int8
	if (v > 64 && v < 79) || (v > 96 && v < 109) {
		diff = 13
	}

	if (v > 78 && v < 90) || (v > 108 && v < 123) {
		diff = -13
	}

	return byte(int8(v) + diff)
}

func (a rot13Reader) Read(c []byte) (int, error) {
	n, err := a.r.Read(c)

	for i, v := range c {
		c[i] = translate(v)
	}

	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
