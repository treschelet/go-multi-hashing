package gomultihashing

// #include "neoscrypt.h"
import "C"

func NeoscryptHash(data []byte, profile uint) [32]byte {
	var buffer [32]byte
	C.neoscrypt((*C.uchar)(&data[0]), (*C.uchar)(&buffer[0]), C.uint(profile))
	return buffer
}
