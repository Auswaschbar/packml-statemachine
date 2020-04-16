package models

import "fmt"

func Abort(machine PackMLMachine) error {
	if machine.GetState() == Aborting || machine.GetState() == Aborted {
		return &InvalidCommand{machine.GetState(), "Abort"}
	}

	machine.SetState(Aborting)
	machine.Aborting()
	return nil
}

func Clear(machine PackMLMachine) error {
	if machine.GetState() != Aborted {
		return &InvalidCommand{machine.GetState(), "Clear"}
	}
	machine.SetState(Clearing)
	machine.Clearing()
	return nil
}

func Reset(machine PackMLMachine) error {
	if machine.GetState() != Complete &&
		machine.GetState() != Stopped {
		return &InvalidCommand{machine.GetState(), "Reset"}
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
		return &InvalidCommand{machine.GetState(), "Stop"}
	}

	machine.SetState(Stopping)
	machine.Stopping()
	return nil
}

func Start(machine PackMLMachine) error {
	if machine.GetState() != Idle {
		return &InvalidCommand{machine.GetState(), "Start"}
	}
	machine.SetState(Starting)
	machine.Starting()
	return nil
}

func Hold(machine PackMLMachine) error {
	if machine.GetState() != Execute {
		return &InvalidCommand{machine.GetState(), "Hold"}
	}
	machine.SetState(Holding)
	machine.Holding()
	return nil
}

func Unhold(machine PackMLMachine) error {
	if machine.GetState() != Held {
		return &InvalidCommand{machine.GetState(), "Unhold"}
	}
	machine.SetState(Unholding)
	machine.Unholding()
	return nil
}

func Suspend(machine PackMLMachine) error {
	if machine.GetState() != Execute {
		return &InvalidCommand{machine.GetState(), "Suspend"}
	}
	machine.SetState(Suspending)
	machine.Suspending()
	return nil
}

func Unsuspend(machine PackMLMachine) error {
	if machine.GetState() != Suspended {
		return &InvalidCommand{machine.GetState(), "Unsuspend"}
	}
	machine.SetState(Unsuspending)
	machine.Unsuspending()
	return nil
}

type InvalidCommand struct {
	CurrentState MachineState
	Command      string
}

func (e *InvalidCommand) Error() string {
	return fmt.Sprintf("Cannot run %s in state %s", e.Command, +e.CurrentState)
}
