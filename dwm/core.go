package dwm

/*
#cgo CFLAGS: -std=c99 -pedantic -Wno-deprecated-declarations
#cgo CFLAGS: -Os -I/usr/X11R6/include
#cgo CFLAGS: -D_BSD_SOURCE -D_POSIX_C_SOURCE=2 -DXINERAMA

#cgo LDFLAGS: -s -L/usr/X11R6/lib -lX11 -lXinerama

#include "dwm.h"

void checkothervm(void);

char* version() {
	return VERSION;
}

XEvent* handlerWrapper(int);

*/
import "C"

import (
	"bytes"
	"encoding/binary"
	"unsafe"
)

func TestInitialization() {
	C.test_initialization()
}

func CheckOtherWM() {
	C.checkotherwm()
}

func Setup() {
	C.setup()
}

func Scan() {
	C.scan()
}

func Run() {
	var ev C.XEvent
	C.XSync(C.dpy, C.False)

	for C.running == C.True && C.XNextEvent(C.dpy, &ev) == 0 {
		var event_type C.int
		binary.Read(bytes.NewBuffer(ev[:unsafe.Sizeof(event_type)]), binary.LittleEndian, &event_type)
		if C.handler[event_type] != nil {
			C.handler[event_type](&ev)
		}
	}
}

func Cleanup() {
	C.cleanup()
}

func CloseWM() int {
	return int(C.close_wm())
}

func Version() string {
	return C.GoString(C.version())
}
