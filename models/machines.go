package models

type MachineSchema struct {
	Id    int
	Name  string
	State MachineState
}

func (db *DB) AllMachines() ([]*MachineSchema, error) {
	rows, err := db.Query("SELECT id, name, state FROM machines")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	bks := make([]*MachineSchema, 0)
	for rows.Next() {
		bk := new(MachineSchema)
		err := rows.Scan(&bk.Id, &bk.Name, &bk.State)
		if err != nil {
			return nil, err
		}
		bks = append(bks, bk)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return bks, nil
}

/*func (db *DB) NewMachine(machineName string) (*Machine, error) {
	_, err := db.Exec("INSERT INTO machines (name, state) VALUES (?, ?)", machineName, Stopped)
	if err != nil {
		return nil, err
	}


}*/
