package dwm

/*
#cgo CFLAGS: -std=c99 -pedantic -Wno-deprecated-declarations -Wno-unused -Wno-unused-parameter
#cgo CFLAGS: -Os -I/usr/X11R6/include
#cgo CFLAGS: -D_BSD_SOURCE -D_POSIX_C_SOURCE=2 -DXINERAMA

#cgo LDFLAGS: -s -L/usr/X11R6/lib -lX11 -lXinerama

#include "dwm.h"

void checkothervm(void);

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

type EventHandler func(*X.Event)

//go:generate bash ../update_version.sh ${GOFILE}
//go:generate git add ${GOFILE}
var version string = "7.1.16"
var name string = "go-dwm"

var currentEventHandler EventHandler = dwmEventHandler

func dwmEventHandler(event *X.Event) {
	var cEventType C.int = C.int(event.EventType())
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

func invokeEventHandler(event *X.Event) {
	currentEventHandler(event)
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

	for C.running == C.True {
		var ev X.Event
		if (dpy.NextEvent(&ev) != 0) {
			break
		}

		eventType := ev.EventType()
		if eventType < int(C.LASTEvent) {
			invokeEventHandler(&ev);
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
	return version
}

func SetVersion(v string) {
	version = v
}

func Name() string {
	return name
}

func SetName(n string) {
	name = n
}
