package main

import (
	"fmt"
	"plugin"
	"reflect"

	"github.com/mateuszdyminski/go-plugin/printer"
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

	printerImpl, ok := p.(printer.Printer)
	if !ok {
		fmt.Printf("wrong type: %+v \n", reflect.TypeOf(p))
		panic("wrong type")
	}

	printerImpl.Print("test")
}
