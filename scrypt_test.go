package multihashing

import (
	"bytes"
	"testing"
)

func TestScryptHash(t *testing.T) {
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
	scryptHash := ScryptHash(data, 0, 0)
	exp := [32]byte{
		0x34, 0x8c, 0xac, 0x6f, 0x07, 0x60, 0x0b, 0x7d,
		0xbb, 0x1e, 0xc8, 0x12, 0x46, 0x56, 0xfc, 0xd0,
		0xbd, 0x14, 0xca, 0xb1, 0xbe, 0x59, 0x28, 0xa9,
		0x7b, 0x0b, 0xb6, 0x1b, 0xbf, 0x62, 0x04, 0x73,
	}
	if bytes.Compare(scryptHash[:], exp[:]) != 0 {
		t.Error("Hashing produced", scryptHash, "expected", exp)
	}
}

func TestScryptnHash(t *testing.T) {
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
	scryptnHash := ScryptnHash(data, 1)
	exp := [32]byte{
		0x64, 0x7d, 0x9a, 0xb5, 0x3d, 0x94, 0xa1, 0xbf,
		0xd6, 0xab, 0xf3, 0x5a, 0xc0, 0xa4, 0x71, 0x04,
		0x93, 0xf3, 0x7f, 0x5a, 0x4d, 0xe2, 0x79, 0xc6,
		0xa6, 0x75, 0x62, 0x05, 0x6c, 0x26, 0xaf, 0x3c,
	}
	if bytes.Compare(scryptnHash[:], exp[:]) != 0 {
		t.Error("Hashing produced", scryptnHash, "expected", exp)
	}
}
