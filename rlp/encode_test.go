package rlp

import (
	"bytes"
	"testing"
)

type foo struct {
	a string
	b uint64
}

func TestEncode(t *testing.T) {
	b := new(bytes.Buffer)

	val := foo{
		a: "abcd",
		b: 10,
	}

	Encode(b, val)
}
