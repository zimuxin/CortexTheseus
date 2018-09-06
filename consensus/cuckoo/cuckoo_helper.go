package cuckoo

/*
#cgo LDFLAGS: -L../../cminer -lgominer -lstdc++
#cgo CFLAGS: -I../../cminer

#include "gominer.h"
*/
import "C"
import (
	"unsafe"
)

func CuckooInit(threads uint) {
	C.CuckooInit(C.uint(threads))
}

func CuckooFinalize() {
	C.CuckooFinalize()
}

func CuckooSolve(hash *byte, hash_len int, nonce uint32, result *uint32, result_len *uint32, diff *byte, result_hash *byte) byte {

	r := C.CuckooSolve(
		(*C.uint8_t)(unsafe.Pointer(hash)),
		C.uint32_t(hash_len),
		C.uint32_t(nonce),
		(*C.uint32_t)(unsafe.Pointer(result)),
		(*C.uint32_t)(unsafe.Pointer(result_len)),
		(*C.uint8_t)(unsafe.Pointer(diff)),
		(*C.uint8_t)(unsafe.Pointer(result_hash)))

	return byte(r)
}

func CuckooVerify(hash *byte, hash_len int, nonce uint32, result *uint32, diff *byte, result_hash *byte) byte {
	r := C.CuckooVerify(
		(*C.uchar)(unsafe.Pointer(hash)),
		C.uint(hash_len),
		C.uint(uint32(nonce)),
		(*C.result_t)(unsafe.Pointer(result)),
		(*C.uchar)(unsafe.Pointer(diff)),
		(*C.uchar)(unsafe.Pointer(result_hash)))

	return byte(r)
}
