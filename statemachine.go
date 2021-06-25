package statemachine

import "fmt"

type FMS struct {
	current     string
	States      States
	Events      Events
	Transitions map[eKey]string
	Action      map[string]ActionCall
}

type FMSBuilder struct {
	Fms *FMS
}

func NewFMSBuilder() *FMSBuilder {
	return &FMSBuilder{Fms: &FMS{
		current:     "nil",
		States:      States{},
		Events:      Events{},
		Transitions: make(map[eKey]string),
	}}
}

func (b *FMSBuilder) AddStates(s State) *FMSBuilder {
	b.Fms.States.Add(s)
	b.Fms.current = s.Initial
	return b
}

func (b *FMSBuilder) AddEvents(es Events) *FMSBuilder {
	for _, e1 := range es {
		b.Fms.Transitions[eKey{e1.Source, e1.Name}] = e1.Target
	}
	return b
}

func (b *FMSBuilder) Build() *FMS {
	return b.Fms
}

func (f *FMS) SendEvent(e string) error {
	dst, ok := f.Transitions[eKey{f.current, e}]
	if !ok {
		for ekey := range f.Transitions {
			if ekey.event == e {
				return InvalidEventError{e, f.current}
			}
		}
		return UnknownEventError{e}
	}

	if f.current == dst {
		return NoTransitionError{fmt.Errorf("NoTransitionError")}
	}

	f.current = dst
	return fmt.Errorf("Transition Error")
}

func (f *FMS) GetState() string {
	return f.current
}

type eKey struct {
	event string
	src   string
}

type InvalidEventError struct {
	Event string
	State string
}

func (e InvalidEventError) Error() string {
	return "event " + e.Event + " inappropriate in current state " + e.State
}

type UnknownEventError struct {
	Event string
}

func (e UnknownEventError) Error() string {
	return "event " + e.Event + " does not exist"
}

type NoTransitionError struct {
	Err error
}

func (e NoTransitionError) Error() string {
	if e.Err != nil {
		return "no transition with error: " + e.Err.Error()
	}
	return "no transition"
}
