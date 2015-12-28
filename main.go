package main

/*
#cgo CFLAGS: -std=c99 -pedantic -Wall -Wno-deprecated-declarations
#cgo CFLAGS: -Os -I/usr/X11R6/include
#cgo CFLAGS: -D_BSD_SOURCE -D_POSIX_C_SOURCE=2 -DXINERAMA

#cgo LDFLAGS: -s -L/usr/X11R6/lib -lX11 -lXinerama

#include "dwm.h"

int main_impl(int argc, char* argv[]);
*/
import "C"

import (
	"os"
	"unsafe"
)

func main() {
	c_args := goToCArgumentList(os.Args)
	defer freeCArgs(c_args)

	os.Exit(int(C.main_impl(C.int(len(os.Args)), c_args)))
}

func goToCArgumentList(go_args []string) **C.char {
	c_args := (**C.char)(C.malloc(C.size_t(len(go_args)) * sizeOfChar()))
	for i := range os.Args {
		*c_args = C.CString(os.Args[i])
	}
	return c_args
}

func freeCArgs(c_args **C.char) {
	C.free(unsafe.Pointer(c_args))
}

func sizeOfChar() C.size_t {
	var b *C.char
	return C.size_t(unsafe.Sizeof(b))
}
