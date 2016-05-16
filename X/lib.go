package X

/*
#cgo CFLAGS: -std=c99 -pedantic -Wno-deprecated-declarations -Wno-unused -Wno-unused-parameter
#cgo CFLAGS: -Os -I/usr/X11R6/include
#cgo CFLAGS: -D_BSD_SOURCE -D_POSIX_C_SOURCE=2 -DXINERAMA

#cgo LDFLAGS: -s -L/usr/X11R6/lib -lX11 -lXinerama

#include <X11/cursorfont.h>
#include <X11/keysym.h>
#include <X11/Xatom.h>
#include <X11/Xlib.h>
#include <X11/Xproto.h>
#include <X11/Xutil.h>
*/
import "C"

type Display C.Display

func Sync(display *Display, discard bool) int {
	if discard {
		return int(C.XSync((*C.Display)(display), C.Bool(1)))
	} else {
		return int(C.XSync((*C.Display)(display), C.Bool(0)))
	}
}
