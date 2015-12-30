package dwm

/*
#cgo CFLAGS: -std=c99 -pedantic -Wall -Wno-deprecated-declarations
#cgo CFLAGS: -Os -I/usr/X11R6/include
#cgo CFLAGS: -D_BSD_SOURCE -D_POSIX_C_SOURCE=2 -DXINERAMA

#cgo LDFLAGS: -s -L/usr/X11R6/lib -lX11 -lXinerama

#include "dwm.h"

int main_impl(int argc, char* argv[]);
*/
import "C"
