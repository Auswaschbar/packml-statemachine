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
		"STOPPED",
		"STOPPING",
		"CLEARING",
		"RESETTING",
		"IDLE",
		"STARTING",
		"EXECUTE",
		"HOLDING",
		"HELD",
		"UNHOLDING",
		"SUSPENDING",
		"SUSPENDED",
		"UNSUSPENDING",
		"COMPLETING",
		"COMPLETE",
		"ABORTING",
		"ABORTED"}

	if state < Stopped || state > Aborted {
		return "Unknown"
	}
	return states[state]
}

// Interface for acting states
// Implement this to execute something when the machine switches its state
type PackMLActor interface {
	Resetting()
	Starting()
	Execute()
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

	// Check whether the machine is currently in a waiting state
	IsWaiting() bool

	// Check whether the machine is currently in an acting state
	IsActing() bool

	// finish the current acting state, and enter the following wait state
	StateCompletion() error
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

func (machine *Machine) IsWaiting() bool {
	switch machine.state {
	case Stopped:
	case Idle:
	case Held:
	case Suspended:
	case Complete:
	case Aborted:
		return true
	}
	return false
}

func (machine *Machine) IsActing() bool {
	return !machine.IsWaiting()
}

func (machine *Machine) StateCompletion() error {
	switch machine.state {
	case Resetting:
		machine.SetState(Idle)
	case Starting:
		machine.SetState(Execute)
		machine.Execute()
	case Holding:
		machine.SetState(Held)
	case Unholding:
		machine.SetState(Execute)
		machine.Execute()
	case Suspending:
		machine.SetState(Suspended)
	case Unsuspending:
		machine.SetState(Execute)
		machine.Execute()
	case Execute:
		machine.SetState(Completing)
		machine.Completing()
	case Completing:
		machine.SetState(Complete)
	case Aborting:
		machine.SetState(Aborted)
	case Clearing:
		machine.SetState(Stopped)
	case Stopping:
		machine.SetState(Stopped)
	default:
		return fmt.Errorf("State cannot be completed: %s", machine.state)
	}
	return nil
}

func NewMachine() Machine {
	something := Machine{}
	something.state = Stopped
	return something
}

func (machine *Machine) Resetting()    {}
func (machine *Machine) Starting()     {}
func (machine *Machine) Execute()      {}
func (machine *Machine) Holding()      {}
func (machine *Machine) Unholding()    {}
func (machine *Machine) Suspending()   {}
func (machine *Machine) Unsuspending() {}
func (machine *Machine) Completing()   {}
func (machine *Machine) Aborting()     {}
func (machine *Machine) Clearing()     {}
func (machine *Machine) Stopping()     {}
