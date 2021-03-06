package dwm

/*
#include <X11/Xlib.h>

extern void Handler(XEvent* event);
*/
import "C"

import (
	"github.com/akesling/gox/X"
)

// This function exists in its own file to allow import into C.  The preamble
// for a Cgo file is included multiple times, so it can't have definitions (as
// we need to in order to interact with dwm.h).
//export Handler
func Handler(event *C.XEvent) {
	invokeEventHandler((*X.Event)(event))
}

//export wm_version_name
func wm_version_name() *C.char {
	return C.CString(Name() + "-" + Version())
}
