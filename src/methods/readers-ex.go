package main

// Implement a Reader type that emits an infinite stream of the ASCII character 'A'.

import "golang.org/x/tour/reader"

type MyReader struct{}

func (mr MyReader) Read(b []byte) (int, error) {
	b[0] = 'A'
	return 1, nil
}

func main() {
	reader.Validate(MyReader{})
}
