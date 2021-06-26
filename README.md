# Statemachine
It is a standalone implementation of FSM.
Finite state machine, is a mathematical model of computation. It is an abstract machine that can be in exactly one of a finite number of states at any given time. The FSM can change from one state to another in response to some inputs; the change from one state to another is called a transition. An FSM is defined by a list of its states, its initial state, and the inputs that trigger each transition. 

## Note
As of now concurrency is not supported. Expected soon.

## Usage

## Create State machine:
```go
b := NewFMSBuilder()
```

## Define States and Events:
```go
b.AddStates(State{Initial: "SS", End: "SE", Intermedite: []string{"S1","S2"}})
	b.AddEvents(Events{
		{Source: "SS", Target: "S1", Name: "E1", Action: func(e *Event) {}},
		{Source: "S1", Target: "S2", Name: "E2", Action: func(e *Event) {}},
		{Source: "S2", Target: "SE", Name: "E3", Action: func(e *Event) {}},
	})
	fms := b.Build()
```

## Trigger an Event and get status:
```go
fms.SendEvent("E1")
fmt.Println(">> ", fms.GetState())
```

## License
[CJ](https://github.com/chandrasekar-om)
