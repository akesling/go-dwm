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
*/
import "C"

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
	C.run()
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
