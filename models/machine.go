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

type PackMLMachine interface {
	Abort() error
	Clear() error
	Reset() error
	Stop() error
	Start() error

	Hold() error
	Unhold() error

	Suspend() error
	Unsuspend() error

	StatusChanged() error
}

type Machine struct {
	State MachineState
}

type InvalidTransitionError struct {
	CurrentState MachineState
	Command      string
}

func (e *InvalidTransitionError) Error() string {
	return fmt.Sprintf("Cannot run %s in state %s", e.Command, +e.CurrentState)
}

func (machine *Machine) Abort() error {
	if machine.State == Aborting || machine.State == Aborted {
		return &InvalidTransitionError{machine.State, "Abort"}
	}

	machine.State = Aborting
	return machine.StatusChanged()
}

func (machine *Machine) Clear() error {
	if machine.State != Aborted {
		return &InvalidTransitionError{machine.State, "Clear"}
	}
	machine.State = Clearing
	return nil
}

func (machine *Machine) Reset() error {
	if machine.State != Complete || machine.State != Stopped {
		return &InvalidTransitionError{machine.State, "Reset"}
	}
	machine.State = Resetting
	return machine.StatusChanged()
}

func (machine *Machine) Stop() error {
	if machine.State == Aborting ||
		machine.State == Aborted ||
		machine.State == Clearing ||
		machine.State == Stopping ||
		machine.State == Stopped {
		return &InvalidTransitionError{machine.State, "Stop"}
	}

	machine.State = Stopping
	return machine.StatusChanged()
}

func (machine *Machine) Start() error {
	if machine.State != Idle {
		return &InvalidTransitionError{machine.State, "Start"}
	}
	machine.State = Starting
	return machine.StatusChanged()
}

func (machine *Machine) Hold() error {
	if machine.State != Execute {
		return &InvalidTransitionError{machine.State, "Hold"}
	}
	machine.State = Holding
	return machine.StatusChanged()
}

func (machine *Machine) Unhold() error {
	if machine.State != Held {
		return &InvalidTransitionError{machine.State, "Unhold"}
	}
	machine.State = Unholding
	return machine.StatusChanged()
}

func (machine *Machine) Suspend() error {
	if machine.State != Execute {
		return &InvalidTransitionError{machine.State, "Suspend"}
	}
	machine.State = Suspending
	return machine.StatusChanged()
}

func (machine *Machine) Unsuspend() error {
	if machine.State != Suspended {
		return &InvalidTransitionError{machine.State, "Unsuspend"}
	}
	machine.State = Unsuspending
	return machine.StatusChanged()
}

func (machine *Machine) StatusChanged() error {
	return nil
}
