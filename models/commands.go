package models

import "fmt"

// Abort a machine
func Abort(machine PackMLMachine) error {
	if machine.GetState() == Aborting || machine.GetState() == Aborted {
		return &InvalidCommand{machine.GetState(), "Abort"}
	}

	machine.SetState(Aborting)
	machine.Aborting()
	return nil
}

// Clear a machine
// Must be in Aborted state
func Clear(machine PackMLMachine) error {
	if machine.GetState() != Aborted {
		return &InvalidCommand{machine.GetState(), "Clear"}
	}
	machine.SetState(Clearing)
	machine.Clearing()
	return nil
}

// Reset a machine
// Must be in Completed or Stopped state
func Reset(machine PackMLMachine) error {
	if machine.GetState() != Complete &&
		machine.GetState() != Stopped {
		return &InvalidCommand{machine.GetState(), "Reset"}
	}
	machine.SetState(Resetting)
	machine.Resetting()
	return nil
}

// Stop a machine
// Machine must be running
func Stop(machine PackMLMachine) error {
	if machine.GetState() == Aborting ||
		machine.GetState() == Aborted ||
		machine.GetState() == Clearing ||
		machine.GetState() == Stopping ||
		machine.GetState() == Stopped {
		return &InvalidCommand{machine.GetState(), "Stop"}
	}

	machine.SetState(Stopping)
	machine.Stopping()
	return nil
}

// Start a machine
// Machine must be in Idle state
func Start(machine PackMLMachine) error {
	if machine.GetState() != Idle {
		return &InvalidCommand{machine.GetState(), "Start"}
	}
	machine.SetState(Starting)
	machine.Starting()
	return nil
}

// Hold a machine
// Machine must be in Execute state
func Hold(machine PackMLMachine) error {
	if machine.GetState() != Execute {
		return &InvalidCommand{machine.GetState(), "Hold"}
	}
	machine.SetState(Holding)
	machine.Holding()
	return nil
}

// Unhold a machine
// Machine must be in Holding state
func Unhold(machine PackMLMachine) error {
	if machine.GetState() != Held {
		return &InvalidCommand{machine.GetState(), "Unhold"}
	}
	machine.SetState(Unholding)
	machine.Unholding()
	return nil
}

// Suspend a machine
// Machine must be in Execute state
func Suspend(machine PackMLMachine) error {
	if machine.GetState() != Execute {
		return &InvalidCommand{machine.GetState(), "Suspend"}
	}
	machine.SetState(Suspending)
	machine.Suspending()
	return nil
}

// Unsuspend a machine
// Machine must be in Suspended state
func Unsuspend(machine PackMLMachine) error {
	if machine.GetState() != Suspended {
		return &InvalidCommand{machine.GetState(), "Unsuspend"}
	}
	machine.SetState(Unsuspending)
	machine.Unsuspending()
	return nil
}

// Error that is thrown when the issued command is invalid for the current machine state
type InvalidCommand struct {
	CurrentState MachineState
	Command      string
}

func (e *InvalidCommand) Error() string {
	return fmt.Sprintf("Cannot run %s in state %s", e.Command, +e.CurrentState)
}
