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
func (Machine *Machine) InsertMachine(db *sql.DB) (int64, error){
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

func (machine Machine) GetMachineId (db *sql.DB) (int64, error){
	var(
		err error
		id int64
	)

	if machine.Id!=nil{
		id=*machine.Id
	}
	if machine.Name!=nil{
		sql:=`SELECT id FROM machines WHERE name=?`
		err=db.QueryRow(sql,machine.Name).Scan(&id)
		if err!=nil{return 0, fmt.Errorf("Cant find product: %v",err)}
	}
	return id,nil
}

func (target Machine) DeleteMachine(db *sql.DB) (int64, error){
	var result sql.Result
	
	id,err:=target.GetMachineId(db)
	if err!=nil{return 0,err}

	sql :=`DELETE * FROM machines WHERE id=?`
	result, err=db.Exec(sql,id)
	if err !=nil{return 0,err}
	return result.RowsAffected()
}

func (m *Machine) GetMachine(db *sql.DB) (error){

	id,err:=m.GetMachineId(db)
	if err!=nil{return err}

	sql:=`SELECT * FROM machines WHERE id=?`
	row:=db.QueryRow(sql,id)
	err = row.Scan(&m.Id, &m.Name, &m.Crafting_speed, &m.Polution, &m.Module_slot, &m.Q_coef_a, &m.Q_coef_b, &m.Q5_mod, &m.Drain, &m.Energy_consumption)
	if err!=nil{return err}
	return nil
}

func (newMach *Machine) UpdateMachine(db *sql.DB) (int64, error){
	mach:=newMach
	mach.GetMachine(db)

	if newMach.Name!=nil{mach.Name=newMach.Name}
	if newMach.Crafting_speed!=nil{mach.Crafting_speed=newMach.Crafting_speed}
	if newMach.Polution!=nil{mach.Polution=newMach.Polution}
	if newMach.Module_slot!=nil{mach.Module_slot=newMach.Module_slot}
	if newMach.Q_coef_a!=nil{mach.Q_coef_a=newMach.Q_coef_a}
	if newMach.Q_coef_b!=nil{mach.Q_coef_b=newMach.Q_coef_b}
	if newMach.Q5_mod!=nil{mach.Q5_mod=newMach.Q5_mod}
	if newMach.Drain!=nil{mach.Drain=newMach.Drain}
	if newMach.Energy_consumption!=nil{mach.Energy_consumption=newMach.Energy_consumption}

	sql:=`UPDATE machines SET name=?,
	crafting_speed=?,
	Polution=?,
	Module_slot=?,
	Q_coef_a=?
	Q_coef_b=?,
	Q5_mod=?,
	Drain=?,
	Energy_consumption=?`

	res,err:=db.Exec(sql,mach.Name,
		mach.Crafting_speed,
		mach.Polution,
		mach.Module_slot,
		mach.Q_coef_a,
		mach.Q_coef_b,
		mach.Q5_mod,
		mach.Drain,
		mach.Energy_consumption)
		
	if err!=nil{return 0,fmt.Errorf("cant update machine %v",err)}
	return res.RowsAffected()
}
