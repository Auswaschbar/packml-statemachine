package main

import (
	"fmt"
	"log"
	"net/http"
	"test/models"
)

type Env struct {
	db models.Datastore
}

func main() {
	db, err := models.NewDB("root:root@tcp(localhost)/machines")
	if err != nil {
		log.Panic(err)
	}

	env := &Env{db}

	http.HandleFunc("/machines", env.machines_index)
	http.ListenAndServe(":8090", nil)
}

func (env *Env) machines_index(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), 405)
		return
	}
	machines, err := env.db.AllMachines()
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	for _, mach := range machines {
		fmt.Fprintf(w, "%d, %s, %d\n", mach.Id, mach.Name, mach.State)
	}
}
