package lib

import (
	"database/sql"
	"fmt"
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

func InsertMachine(db *sql.DB, Machine *Machine) (int64, error){
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

func DeleteMachine(db *sql.DB, target Target) (int64, error){
	var(
		result sql.Result
		err error
	)
	if target.Id!=nil{
			sql :=`DELETE FROM machines WHERE id=?`
			result, err=db.Exec(sql,target.Id)
	}else if target.Name!=nil{
			sql :=`DELETE FROM machines WHERE name=?`
			result, err=db.Exec(sql,target.Name)
	}else{
			return 0, fmt.Errorf("wrong argument type: supported types int-id string-name")}
	
	if err !=nil{return 0,err}
	return result.RowsAffected()
}
