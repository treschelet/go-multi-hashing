package multihashing

import (
	"bytes"
	"testing"
)

func TestNeoScryptHash(t *testing.T) {
	data := []byte{
		0x70, 0x00, 0x00, 0x00, 0x01, 0xe9, 0x80, 0x92,
		0x4e, 0x4e, 0x11, 0x09, 0x23, 0x03, 0x83, 0xe6,
		0x6d, 0x62, 0x94, 0x5f, 0xf8, 0xe7, 0x49, 0x90,
		0x3b, 0xea, 0x43, 0x36, 0x75, 0x5c, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x51, 0x92, 0x8a, 0xff,
		0x1b, 0x4d, 0x72, 0x41, 0x61, 0x73, 0xa8, 0xc3,
		0x94, 0x81, 0x59, 0xa0, 0x9a, 0x73, 0xac, 0x3b,
		0xb5, 0x56, 0xaa, 0x6b, 0xfb, 0xca, 0xd1, 0xa8,
		0x5d, 0xa7, 0xf4, 0xc1, 0xd1, 0x33, 0x50, 0x53,
		0x1e, 0x24, 0x03, 0x1b, 0x93, 0x9b, 0x9e, 0x2b,
	}
	neoscryptHash := NeoscryptHash(data, 0)
	exp := [32]byte{
		0x94, 0x36, 0x7f, 0x3f, 0x5b, 0x4d, 0x60, 0xe0,
		0x02, 0x1e, 0x5b, 0x73, 0xf6, 0x45, 0x00, 0x3e,
		0xbf, 0xd4, 0x10, 0xb0, 0xf6, 0x7f, 0x7b, 0xe9,
		0xf2, 0x13, 0x67, 0xa0, 0x2d, 0x68, 0x97, 0xae,
	}
	if bytes.Compare(neoscryptHash[:], exp[:]) != 0 {
		t.Error("Hashing produced", neoscryptHash, "expected", exp)
	}
}
