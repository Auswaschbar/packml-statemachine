package main

import (
	"errors"
	"test/models"
	"testing"
)

func TestVanilla(t *testing.T) {
	env := models.Machine{State: models.Stopped}
	err := env.Abort()
	if err != nil {
		t.Errorf("Abort(): %s\n", err)
	}

	err = env.Clear()
	var e *models.InvalidTransitionError
	if !errors.As(err, &e) {
		t.Errorf("Clear(): %s\n", err)
	}

	err = env.Reset()
	if !errors.As(err, &e) {
		t.Errorf("Reset(): %s\n", err)
	}
}
