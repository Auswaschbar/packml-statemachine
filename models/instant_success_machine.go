package models

type InstantSuccessMachine struct {
	Machine
}

func (machine *InstantSuccessMachine) Clearing() {
	machine.SetState(Idle)
}

func (machine *InstantSuccessMachine) Stopping() {
	machine.SetState(Idle)
}

func (machine *InstantSuccessMachine) Aborting() {
	machine.SetState(Aborted)
}

func (machine *InstantSuccessMachine) Holding() {
	machine.SetState(Held)
}

func (machine *InstantSuccessMachine) Suspending() {
	machine.SetState(Suspended)
}

func (machine *InstantSuccessMachine) Starting() {
	machine.SetState(Execute)
}

func (machine *InstantSuccessMachine) Unholding() {
	machine.SetState(Execute)
}

func (machine *InstantSuccessMachine) Unsuspending() {
	machine.SetState(Execute)
}

func (machine *InstantSuccessMachine) Resetting() {
	machine.SetState(Idle)
}

func (machine *InstantSuccessMachine) Completing() {
	machine.SetState(Complete)
}
