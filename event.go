package statemachine

type Event struct {
	Source string
	Target string
	Name   string
	Action ActionCall
}

type ActionCall func(*Event)

type Events []Event
