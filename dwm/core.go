package dwm

/*
#cgo CFLAGS: -std=c99 -pedantic -Wno-deprecated-declarations
#cgo CFLAGS: -Os -I/usr/X11R6/include
#cgo CFLAGS: -D_BSD_SOURCE -D_POSIX_C_SOURCE=2 -DXINERAMA

#cgo LDFLAGS: -s -L/usr/X11R6/lib -lX11 -lXinerama

#include "dwm.h"

int main_impl(int argc, char* argv[]);
void checkothervm(void);

*/
import "C"

import (
	"os"
	"unsafe"
)

func MainImpl(argc int, argv **C.char) int {
	TestInitialization(argc, argv)
	CheckOtherWM()
	Setup()
	Scan()
	Run()
	Cleanup()
	return CloseWM()
}

func GoToCArgumentList(go_args []string) **C.char {
	c_args := (**C.char)(C.malloc(C.size_t(len(go_args)) * sizeOfChar()))
	for i := range os.Args {
		*c_args = C.CString(os.Args[i])
	}
	return c_args
}

func FreeCArgs(c_args **C.char) {
	C.free(unsafe.Pointer(c_args))
}

func sizeOfChar() C.size_t {
	var b *C.char
	return C.size_t(unsafe.Sizeof(b))
}

func TestInitialization(argc int, argv **C.char) {
	C.test_initialization(C.int(argc), argv)
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
