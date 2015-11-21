package main

// #cgo CFLAGS: -std=c99 -pedantic -Wall -Wno-deprecated-declarations -Os -I/usr/X11R6/include -D_BSD_SOURCE -D_POSIX_C_SOURCE=2 -DXINERAMA
// #cgo LDFLAGS: drw.o util.o -s -L/usr/X11R6/lib -lX11 -lXinerama
// #include "dwm.h"
//
// int main_impl(int argc, char* argv[]);
import "C"
import "os"

func main() {
	var c_args []*C.char
	for i := range os.Args {
		c_args = append(c_args, C.CString(os.Args[i]))
	}
	os.Exit(int(C.main_impl(C.int(len(os.Args)), &c_args[0])))
}
