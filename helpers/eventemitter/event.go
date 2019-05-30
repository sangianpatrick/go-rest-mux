package eventemitter

import (
	"gitlab.com/patricksangian/go-rest-mux/src/workers"
)

// Emitter contains channels and behavior for emitting listener
type Emitter struct {
	printCH chan<- interface{}
}

// NewEventEmitter initiate event emitter
func NewEventEmitter() Emitter {
	printCH := make(chan interface{})
	go workers.OnPrint(printCH)
	return Emitter{
		printCH: printCH,
	}
}

// EmitPrint will print data
func (e Emitter) EmitPrint(data interface{}) {
	ch := e.printCH
	ch <- data
}
