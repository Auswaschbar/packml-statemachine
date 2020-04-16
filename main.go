package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"test/models"
)

type MachineStatus struct {
	models.InstantSuccessMachine
}

func main() {
	env := MachineStatus{}

	http.HandleFunc("/state", env.MachineState)
	http.HandleFunc("/commands", env.MachineCommands)
	http.ListenAndServe(":8090", nil)
}

func (m *MachineStatus) MachineState(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), 405)
		return
	}

	fmt.Fprintf(w, "%s", m.GetState())
}

func (m *MachineStatus) MachineCommands(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(405), 405)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		http.Error(w, "can't read body", http.StatusBadRequest)
		return
	}

	switch string(body[:]) {
	case "Abort":
		err = models.Abort(m)
	case "Clear":
		err = models.Clear(m)
	case "Reset":
		err = models.Reset(m)
	case "Stop":
		err = models.Stop(m)
	case "Start":
		err = models.Start(m)
	case "Hold":
		err = models.Hold(m)
	case "Unhold":
		err = models.Unhold(m)
	case "Suspend":
		err = models.Suspend(m)
	case "Unsuspend":
		err = models.Unsuspend(m)
	default:
		log.Printf("Invalid command name: \"%s\"\n", string(body[:]))
		http.Error(w,
			fmt.Sprintf("Invalid command name: \"%s\"\n", string(body[:])),
			http.StatusBadRequest)
	}

	if err != nil {
		log.Printf("Command error: %v", err)
		http.Error(w,
			fmt.Sprintf("Command error: %s\n", err),
			http.StatusBadRequest)
		return
	}
}
