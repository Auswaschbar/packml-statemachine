package main

import (
	"errors"
	"test/models"
	"testing"
)

func TestVanilla(t *testing.T) {
	env := models.Machine{}
	err := models.Abort(&env)
	if err != nil {
		t.Errorf("Abort(): %s\n", err)
	}

	err = models.Clear(&env)
	var e *models.InvalidCommand
	if !errors.As(err, &e) {
		t.Errorf("Clear(): %s\n", err)
	}

	err = models.Reset(&env)
	if !errors.As(err, &e) {
		t.Errorf("Reset(): %s\n", err)
	}
}

func TestInstantSuccessMachine(t *testing.T) {
	env := models.InstantSuccessMachine{}
	err := models.Reset(&env)
	if err != nil {
		t.Errorf("Reset(): %s\n", err)
	}

	err = models.Start(&env)
	if err != nil {
		t.Errorf("Start(): %s\n", err)
	}

	var e *models.InvalidCommand
	err = models.Reset(&env)
	if !errors.As(err, &e) {
		t.Errorf("Reset(): %s\n", err)
	}

	err = models.Hold(&env)
	if err != nil {
		t.Errorf("Hold(): %s\n", err)
	}

	err = models.Unhold(&env)
	if err != nil {
		t.Errorf("Unhold(): %s\n", err)
	}

	err = models.Suspend(&env)
	if err != nil {
		t.Errorf("Suspend(): %s\n", err)
	}

	err = models.Unsuspend(&env)
	if err != nil {
		t.Errorf("Unsuspend(): %s\n", err)
	}

	err = models.Stop(&env)
	if err != nil {
		t.Errorf("Stop(): %s\n", err)
	}
}
