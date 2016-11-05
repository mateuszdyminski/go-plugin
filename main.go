package main

import (
	"fmt"
	"plugin"
	"reflect"

	"github.com/mateuszdyminski/go-plugin/printer"
	"github.com/mateuszdyminski/go-plugin/processor"
)

func main() {
	lib, err := plugin.Open("printer.so")
	if err != nil {
		panic(err)
	}

	p, err := lib.Lookup("Impl")
	if err != nil {
		panic(err)
	}

	printerImpl, ok := p.(*printer.Printer)
	if !ok {
		fmt.Printf("wrong type: %+v \n", reflect.TypeOf(p))
		panic("wrong type")
	}

	(*printerImpl).Print("test")

	lib2, err := plugin.Open("processor.so")
	if err != nil {
		panic(err)
	}

	p2, err := lib2.Lookup("Impl")
	if err != nil {
		panic(err)
	}

	processorImpl, ok := p2.(*processor.Processor)
	if !ok {
		fmt.Printf("wrong type: %+v \n", reflect.TypeOf(p))
		panic("wrong type")
	}

	(*processorImpl).Process("test")

}
