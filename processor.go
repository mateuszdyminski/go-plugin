package main

import (
	"fmt"

	"github.com/mateuszdyminski/go-plugin/processor"
)

func main() {
	Impl.Process("test")
}

type ProcessorImpl struct{}

func (p ProcessorImpl) Process(text string) {
	fmt.Printf("[ProcessorImpl] %s\n", text)
}

var Impl processor.Processor = ProcessorImpl{}
