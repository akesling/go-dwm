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

void set_start_x_error_handler() {
	xerrorxlib = XSetErrorHandler(xerrorstart);
}

void set_x_error_handler() {
	XSetErrorHandler(xerror);
}

Window get_default_root_window(Display* display) {
	return DefaultRootWindow(display);
}

*/
import "C"

import (
	"bytes"
	"encoding/binary"
	"github.com/akesling/gox/X"
	"os"
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
	if C.setlocale(C.LC_CTYPE, C.CString("")) == nil || C.XSupportsLocale() == 0 {
		os.Stderr.WriteString("warning: no locale support\n")
	}

	if C.dpy = C.XOpenDisplay(nil); C.dpy == nil {
		panic("dwm: cannot open display\n")
	}
}

func CheckOtherWM() {
	C.set_start_x_error_handler()
	// this causes an error if some other window manager is running
	C.XSelectInput(C.dpy, C.get_default_root_window(C.dpy), C.SubstructureRedirectMask)
	X.Sync((*X.Display)(C.dpy), false)
	C.set_x_error_handler()
	X.Sync((*X.Display)(C.dpy), false)
}

func Setup() {
	C.setup()
}

func Scan() {
	C.scan()
}

func Run() {
	var ev C.XEvent
	X.Sync((*X.Display)(C.dpy), false)

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
