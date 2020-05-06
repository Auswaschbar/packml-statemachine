package models

// An instant success machine for testing purposes
// This machine instantly finishes all acting states,
// and immediately enters the next waiting state
type InstantSuccessMachine struct {
	Machine
}

func (machine *InstantSuccessMachine) Clearing() {
	defer machine.StateCompletion()
}

func (machine *InstantSuccessMachine) Stopping() {
	defer machine.StateCompletion()
}

func (machine *InstantSuccessMachine) Aborting() {
	defer machine.StateCompletion()
}

func (machine *InstantSuccessMachine) Holding() {
	defer machine.StateCompletion()
}

func (machine *InstantSuccessMachine) Suspending() {
	defer machine.StateCompletion()
}

func (machine *InstantSuccessMachine) Starting() {
	defer machine.StateCompletion()
}

func (machine *InstantSuccessMachine) Unholding() {
	defer machine.StateCompletion()
}

func (machine *InstantSuccessMachine) Unsuspending() {
	defer machine.StateCompletion()
}

func (machine *InstantSuccessMachine) Resetting() {
	defer machine.StateCompletion()
}

func (machine *InstantSuccessMachine) Completing() {
	defer machine.StateCompletion()
}
