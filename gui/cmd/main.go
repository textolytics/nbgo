package main

import (
	"fmt"
	"os"

	"github.com/textolytics/nbgo/gui"
)

func main() {
	if err := gui.RunCLI(os.Args[1:]); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
