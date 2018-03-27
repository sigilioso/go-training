package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func isUpperLetter(b byte) bool {
	return b >= byte('A') && b <= byte('Z')
}

func isLowerLetter(b byte) bool {
	return b >= byte('a') && b <= byte('z')
}

func rot(b byte, min byte, max byte, l int) byte {
	dif := int(max) - int(b)
	if dif > l {
		return byte(int(b) + l)
	}
	return byte(int(min) + (l - dif) - 1)
}

func rot13(b byte) byte {
	switch {
	case isUpperLetter(b):
		v := rot(b, byte('A'), byte('Z'), 13)
		return v
	case isLowerLetter(b):
		v := rot(b, byte('a'), byte('z'), 13)
		return v
	default:
		return b
	}
}

func (rot rot13Reader) Read(b []byte) (int, error) {
	n, err := rot.r.Read(b)
	for i, elm := range b {
		b[i] = rot13(elm)
	}
	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, r)
}
