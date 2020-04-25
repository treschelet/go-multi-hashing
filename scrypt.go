package multihashing

// #include "scryptn.h"
import "C"
import "unsafe"

func ScryptHash(data []byte, nFactor uint, rFactor uint) [32]byte {
	var buffer [32]byte
	C.scrypt_N_R_1_256(
		(*C.char)(unsafe.Pointer(&data[0])),
		(*C.char)(unsafe.Pointer(&buffer[0])),
		C.uint(nFactor),
		C.uint32_t(rFactor),
		C.uint32_t(len(data)),
	)
	return buffer
}

func ScryptnHash(data []byte, nFactor uint) [32]byte {
	var buffer [32]byte
	N := 1 << nFactor
	C.scrypt_N_R_1_256(
		(*C.char)(unsafe.Pointer(&data[0])),
		(*C.char)(unsafe.Pointer(&buffer[0])),
		C.uint(N),
		C.uint32_t(1),
		C.uint32_t(len(data)),
	)
	return buffer
}