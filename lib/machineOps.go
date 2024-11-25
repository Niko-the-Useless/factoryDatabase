package lib

import (
	"database/sql"
)

func CreateMachinesTable(db *sql.DB) (sql.Result, error){
	sql := `CREATE TABLE IF NOT EXISTS machines (
		id INTEGER PRIMARY KEY,
		name TEXT UNIQUE NOT NULL,
		crafting_speed REAL,
		polution REAL,
		module_slot INTEGER,
		q_coef_a REAL,
		q_coef_b REAL,
		q5_mod REAL,
		drain REAL,
		energy_consumption REAL
	);`
	return db.Exec(sql)
}

func ( Machine *Machine) InsertMachine(db *sql.DB) (int64, error){
	sql :=`INSERT INTO machines (
		name,
		crafting_speed,
		polution,
		module_slot,
		q_coef_a,
		q_coef_b,
		q5_mod,
		drain,
		energy_consumption)
		VALUES (?,?,?,?,?,?,?,?,?);`

	result, err :=db.Exec(sql,
		Machine.Name,
		Machine.Crafting_speed,
		Machine.Polution,
		Machine.Module_slot,
		Machine.Q_coef_a,
		Machine.Q_coef_b,
		Machine.Q5_mod,
		Machine.Drain,
		Machine.Energy_consumption)

	if err !=nil{
		return 0,err
	}
	return result.LastInsertId()
}

func (target Target) DeleteMachine(db *sql.DB) (int64, error){
	var result sql.Result
	
	id,err:=target.GetId(db)
	if err!=nil{return 0,err}

	sql :=`DELETE FROM machines WHERE id=?`
	result, err=db.Exec(sql,id)
	if err !=nil{return 0,err}
	return result.RowsAffected()
}
