package models

type InstantSuccessMachine struct {
	State MachineState
}

func (machine *InstantSuccessMachine) StatusChanged() error {
	if machine.State == Clearing ||
		machine.State == Stopping {
		machine.State = Stopped
	}

	if machine.State == Aborting {
		machine.State = Aborted
	}

	if machine.State == Resetting {
		machine.State = Idle
	}

	if machine.State == Holding {
		machine.State = Held
	}

	if machine.State == Suspending {
		machine.State = Suspended
	}

	if machine.State == Starting ||
		machine.State == Unholding ||
		machine.State == Suspending {
		machine.State = Execute
	}

	if machine.State == Completing {
		machine.State = Complete
	}

	return nil
}
