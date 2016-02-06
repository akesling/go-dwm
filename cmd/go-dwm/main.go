package main

import (
	"fmt"
	"github.com/akesling/go-dwm/dwm"
	"os"
)

func main() {
	if len(os.Args) == 2 && os.Args[1] == "-v" {
		fmt.Printf("go-dwm-%s, © see LICENSE for details\n", dwm.Version())
		os.Exit(0)
	} else if len(os.Args) != 1 {
		fmt.Print("usage: dwm [-v]\n")
		os.Exit(1)
	}

	dwm.TestInitialization()
	dwm.CheckOtherWM()
	dwm.Setup()
	dwm.Scan()
	dwm.Run()
	dwm.Cleanup()
	os.Exit(dwm.CloseWM())
}
