package dwm

/*
#cgo CFLAGS: -std=c99 -pedantic -Wno-deprecated-declarations -Wno-unused -Wno-unused-parameter
#cgo CFLAGS: -Os -I/usr/X11R6/include
#cgo CFLAGS: -D_BSD_SOURCE -D_POSIX_C_SOURCE=2 -DXINERAMA

#cgo LDFLAGS: -s -L/usr/X11R6/lib -lX11 -lXinerama

#include "dwm.h"

void checkothervm(void);

char* version() {
	return VERSION;
}

*/
import "C"

import (
	"bytes"
	"encoding/binary"
	"unsafe"
)

func invokeEventHandler(event_type C.int, event *C.XEvent) {
	switch event_type {
	case C.ButtonPress:
		C.buttonpress(event)
	case C.ClientMessage:
		C.clientmessage(event)
	case C.ConfigureRequest:
		C.configurerequest(event)
	case C.ConfigureNotify:
		C.configurenotify(event)
	case C.DestroyNotify:
		C.destroynotify(event)
	case C.EnterNotify:
		C.enternotify(event)
	case C.Expose:
		C.expose(event)
	case C.FocusIn:
		C.focusin(event)
	case C.KeyPress:
		C.keypress(event)
	case C.MappingNotify:
		C.mappingnotify(event)
	case C.MapRequest:
		C.maprequest(event)
	case C.MotionNotify:
		C.motionnotify(event)
	case C.PropertyNotify:
		C.propertynotify(event)
	case C.UnmapNotify:
		C.unmapnotify(event)
	}
}

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
		if event_type < C.LASTEvent {
			Handler(event_type, &ev)
		}
	}
}

func Cleanup() {
	C.cleanup()
}

func CloseWM() int {
	C.XCloseDisplay(C.dpy)
	return int(C.EXIT_SUCCESS)
}

func Version() string {
	return C.GoString(C.version())
}
