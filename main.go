package main

// #cgo CFLAGS: -std=c99 -pedantic -Wall -Wno-deprecated-declarations -Os -I/usr/X11R6/include -D_BSD_SOURCE -D_POSIX_C_SOURCE=2 -DVERSION="6.1" -DXINERAMA
// #cgo LDFLAGS: -s -L/usr/X11R6/lib -lX11 -lXinerama
// #include "dwm.h"
//
// int main_impl(int argc, char* argv[]);
import "C"
import "os"

func main() {
	C.main_impl(len(os.Args), os.Args)
}
