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
	"github.com/akesling/gox/X"
	"os"
)

type EventHandler func(int, *X.Event)

var currentEventHandler EventHandler = dwmEventHandler

func dwmEventHandler(event_type int, event *X.Event) {
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

func invokeEventHandler(eventType int, event *X.Event) {
	currentEventHandler(eventType, event)
}

func SetEventHandler(newHandler EventHandler) {
	currentEventHandler = newHandler
}

func GetEventHandler() EventHandler {
	return currentEventHandler
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
	dpy := (*X.Display)(C.dpy);
	// this causes an error if some other window manager is running
	dpy.SelectInput(
		(X.Window)(C.get_default_root_window(C.dpy)),
		int64(C.SubstructureRedirectMask))
	dpy.Sync(false)
	C.set_x_error_handler()
	dpy.Sync(false)
=======
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
	dpy := (*X.Display)(C.dpy);

	dpy.Sync(false)

	var ev X.Event
	for C.running == C.True && dpy.NextEvent(&ev) == 0 {
		eventType := ev.EventType()
		if eventType < int(C.LASTEvent) {
			invokeEventHandler(eventType, &ev);
		}
	}
}

func Cleanup() {
	C.cleanup()
}

func CloseWM() int {
	(*X.Display)(C.dpy).CloseDisplay()
	return int(C.EXIT_SUCCESS)
}

func Version() string {
	return C.GoString(C.version())
}
