package statemachine

import (
	"fmt"
	"testing"
)

func TestSameState(t *testing.T) {
	fmt.Println("Hello")
	b := NewFMSBuilder()
	b.AddStates(State{Initial: "SS", End: "SE", Intermedite: []string{"closed"}})
	b.AddEvents(Events{
		{Source: "SS", Target: "S1", Name: "E1", Action: func(e *Event) {}},
		{Source: "S1", Target: "S2", Name: "E2", Action: func(e *Event) {}},
		{Source: "S2", Target: "SE", Name: "E3", Action: func(e *Event) {}},
	})
	fms := b.Build()
	fms.SendEvent("E1")
	fmt.Println(">> ", fms.GetState())
}
