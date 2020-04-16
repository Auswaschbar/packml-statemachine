package models

import (
	"fmt"
)

type MachineState int

const (
	Stopped  MachineState = 0
	Stopping MachineState = 1
	Clearing MachineState = 2

	Resetting MachineState = 3
	Idle      MachineState = 4
	Starting  MachineState = 5
	Execute   MachineState = 6

	Holding   MachineState = 7
	Held      MachineState = 8
	Unholding MachineState = 9

	Suspending   MachineState = 10
	Suspended    MachineState = 11
	Unsuspending MachineState = 12

	Completing MachineState = 13
	Complete   MachineState = 14

	Aborting MachineState = 15
	Aborted  MachineState = 16
)

func (state MachineState) String() string {
	states := [...]string{
		"Stopped",
		"Stopping",
		"Clearing",
		"Resetting",
		"Idle",
		"Starting",
		"Execute",
		"Holding",
		"Held",
		"Unholding",
		"Suspending",
		"Suspended",
		"Unsuspending",
		"Completing",
		"Complete",
		"Aborting",
		"Aborted"}

	if state < Stopped || state > Aborted {
		return "Unknown"
	} // return the name of a Weekday
	// constant from the names array
	// above.
	return states[state]
}

type PackMLActor interface {
	Resetting()
	Starting()
	Holding()
	Unholding()
	Suspending()
	Unsuspending()
	Completing()
	Aborting()
	Clearing()
	Stopping()
}

type PackMLStateMachine interface {
	GetState() MachineState
	SetState(MachineState)
}

type PackMLMachine interface {
	PackMLActor
	PackMLStateMachine
}

type Machine struct {
	state MachineState
}

func (machine *Machine) GetState() MachineState {
	return machine.state
}

func (machine *Machine) SetState(new_state MachineState) {
	machine.state = new_state
}

func NewMachine() Machine {
	something := Machine{}
	something.state = Stopped
	return something
}

func (machine *Machine) Resetting()    {}
func (machine *Machine) Starting()     {}
func (machine *Machine) Holding()      {}
func (machine *Machine) Unholding()    {}
func (machine *Machine) Suspending()   {}
func (machine *Machine) Unsuspending() {}
func (machine *Machine) Completing()   {}
func (machine *Machine) Aborting()     {}
func (machine *Machine) Clearing()     {}
func (machine *Machine) Stopping()     {}

type InvalidTransitionError struct {
	CurrentState MachineState
	Command      string
}

func (e *InvalidTransitionError) Error() string {
	return fmt.Sprintf("Cannot run %s in state %s", e.Command, +e.CurrentState)
}

func Abort(machine PackMLMachine) error {
	if machine.GetState() == Aborting || machine.GetState() == Aborted {
		return &InvalidTransitionError{machine.GetState(), "Abort"}
	}

	machine.SetState(Aborting)
	machine.Aborting()
	return nil
}

func Clear(machine PackMLMachine) error {
	if machine.GetState() != Aborted {
		return &InvalidTransitionError{machine.GetState(), "Clear"}
	}
	machine.SetState(Clearing)
	machine.Clearing()
	return nil
}

func Reset(machine PackMLMachine) error {
	if machine.GetState() != Complete &&
		machine.GetState() != Stopped {
		return &InvalidTransitionError{machine.GetState(), "Reset"}
	}
	machine.SetState(Resetting)
	machine.Resetting()
	return nil
}

func Stop(machine PackMLMachine) error {
	if machine.GetState() == Aborting ||
		machine.GetState() == Aborted ||
		machine.GetState() == Clearing ||
		machine.GetState() == Stopping ||
		machine.GetState() == Stopped {
		return &InvalidTransitionError{machine.GetState(), "Stop"}
	}

	machine.SetState(Stopping)
	machine.Stopping()
	return nil
}

func Start(machine PackMLMachine) error {
	if machine.GetState() != Idle {
		return &InvalidTransitionError{machine.GetState(), "Start"}
	}
	machine.SetState(Starting)
	machine.Starting()
	return nil
}

func Hold(machine PackMLMachine) error {
	if machine.GetState() != Execute {
		return &InvalidTransitionError{machine.GetState(), "Hold"}
	}
	machine.SetState(Holding)
	machine.Holding()
	return nil
}

func Unhold(machine PackMLMachine) error {
	if machine.GetState() != Held {
		return &InvalidTransitionError{machine.GetState(), "Unhold"}
	}
	machine.SetState(Unholding)
	machine.Unholding()
	return nil
}

func Suspend(machine PackMLMachine) error {
	if machine.GetState() != Execute {
		return &InvalidTransitionError{machine.GetState(), "Suspend"}
	}
	machine.SetState(Suspending)
	machine.Suspending()
	return nil
}

func Unsuspend(machine PackMLMachine) error {
	if machine.GetState() != Suspended {
		return &InvalidTransitionError{machine.GetState(), "Unsuspend"}
	}
	machine.SetState(Unsuspending)
	machine.Unsuspending()
	return nil
}
