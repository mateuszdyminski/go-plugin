package main

import (
	"fmt"

	"github.com/mateuszdyminski/go-plugin/printer"
)

func main() {
	Impl.Print("test")
}

type PrinterImpl struct{}

func (p PrinterImpl) Print(text string) {
	fmt.Printf("[PrinterImpl] %s\n", text)
}

var Impl printer.Printer = PrinterImpl{}
