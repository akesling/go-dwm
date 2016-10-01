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

func invokeEventHandler(event_type int, event *X.Event) {
	var cEventType C.int = C.int(event_type)
	switch cEventType {
	case C.ButtonPress:
		C.buttonpress((*C.XEvent)(event))
	case C.ClientMessage:
		C.clientmessage((*C.XEvent)(event))
	case C.ConfigureRequest:
		C.configurerequest((*C.XEvent)(event))
	case C.ConfigureNotify:
		C.configurenotify((*C.XEvent)(event))
	case C.DestroyNotify:
		C.destroynotify((*C.XEvent)(event))
	case C.EnterNotify:
		C.enternotify((*C.XEvent)(event))
	case C.Expose:
		C.expose((*C.XEvent)(event))
	case C.FocusIn:
		C.focusin((*C.XEvent)(event))
	case C.KeyPress:
		C.keypress((*C.XEvent)(event))
	case C.MappingNotify:
		C.mappingnotify((*C.XEvent)(event))
	case C.MapRequest:
		C.maprequest((*C.XEvent)(event))
	case C.MotionNotify:
		C.motionnotify((*C.XEvent)(event))
	case C.PropertyNotify:
		C.propertynotify((*C.XEvent)(event))
	case C.UnmapNotify:
		C.unmapnotify((*C.XEvent)(event))
	}
}

func TestInitialization() {
	if C.setlocale(C.LC_CTYPE, C.CString("")) == nil || X.SupportsLocale() == 0 {
		os.Stderr.WriteString("warning: no locale support\n")
	}

	if C.dpy = (*C.Display)(X.OpenDisplay(nil)); C.dpy == nil {
		panic("dwm: cannot open display\n")
	}
}

func CheckOtherWM() {
	C.set_start_x_error_handler()
	// this causes an error if some other window manager is running
	X.SelectInput(
		(*X.Display)(C.dpy),
		(X.Window)(C.get_default_root_window(C.dpy)),
		int64(C.SubstructureRedirectMask))
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
	var ev X.Event
	X.Sync((*X.Display)(C.dpy), false)

	for C.running == C.True && X.NextEvent((*X.Display)(C.dpy), &ev) == 0 {
		var event_type C.int
		binary.Read(bytes.NewBuffer(ev[:unsafe.Sizeof(event_type)]), binary.LittleEndian, &event_type)
		if event_type < C.LASTEvent {
			invokeEventHandler(int(event_type), &ev);
		}
	}
}

func Cleanup() {
	C.cleanup()
}

func CloseWM() int {
	X.CloseDisplay((*X.Display)(C.dpy))
	return int(C.EXIT_SUCCESS)
}

func Version() string {
	return C.GoString(C.version())
}
