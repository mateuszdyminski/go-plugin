package main

import (
	"plugin"

	"github.com/mateuszdyminski/go-plugin/printer"
	"github.com/mateuszdyminski/go-plugin/processor"
)

func main() {
	lib, err := plugin.Open("printer-impl/printer.so")
	if err != nil {
		panic(err)
	}

	p, err := lib.Lookup("Impl")
	if err != nil {
		panic(err)
	}

	printerImpl, ok := p.(*printer.Printer)
	if !ok {
		panic("wrong type")
	}

	(*printerImpl).Print("test")

	lib2, err := plugin.Open("processor-impl/processor.so")
	if err != nil {
		panic(err)
	}

	p2, err := lib2.Lookup("Impl")
	if err != nil {
		panic(err)
	}

	processorImpl, ok := p2.(*processor.Processor)
	if !ok {
		panic("wrong type")
	}

	(*processorImpl).Process("test")
}
