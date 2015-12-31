package main

import (
	"github.com/akesling/go-dwm/dwm"
	"os"
)

func main() {
	c_args := dwm.GoToCArgumentList(os.Args)
	defer dwm.FreeCArgs(c_args)

	os.Exit(dwm.MainImpl(len(os.Args), c_args))
}
