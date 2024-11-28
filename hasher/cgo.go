//go:build cgo

package main

import "C"

//export DoubleMd5Sorted
func DoubleMd5Sorted(input *C.char) *C.char {
	// Convert C string (input) to Go string
	goInput := C.GoString(input)

	// Convert the Go string result back to a C string and return it
	result := doubleMd5Sorted(goInput)

	// Convert the Go string result back to a C string and return it
	return C.CString(result)
}

//export Sha256Raw
func Sha256Raw(input *C.char) *C.char {
	// Convert C string (input) to Go string
	goInput := C.GoString(input)

	// Perform the sha256Raw function in Go
	result := sha256Raw(goInput)

	// Convert the Go string result back to a C string and return it
	return C.CString(result)
}
