package main

import (
	"database/sql"
)

func createProductsTable(db *sql.DB) (sql.Result, error){
	sql := `CREATE TABLE IF NOT EXISTS products (
		id INTEGER PRIMARY KEY,
		name TEXT,
		production_time REAL
	);`
	return db.Exec(sql)
}

func createMachinesTable(db *sql.DB) (sql.Result, error){
	sql := `CREATE TABLE IF NOT EXISTS machines (
		id INTEGER PRIMARY KEY,
		name TEXT,
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
// parent = product produced
//child = product requred
func createBomTable(db *sql.DB) (sql.Result, error){
	sql := `CREATE TABLE IF NOT EXISTS BOM (
		parent_id INTEGER NOT NULL,
		parent_quantity INTEGER,
		child_id INTEGER,
		child_quantity INTEGER,
		byproduct_id INTEGER,
		byproduct_quantity INTEGER
		);`
	return db.Exec(sql)
}

func insertProduct(db *sql.DB, Product *Product) (int64, error){
	sql :=`INSERT INTO products (
		name,
		production_time)
		VALUES (?,?);`

	result, err :=db.Exec(sql,
		Product.name,
		Product.production_time)

	if err !=nil{
		return 0,err
	}
	return result.LastInsertId()
}

func insertMachine(db *sql.DB, Machine *Machine) (int64, error){
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
		Machine.name,
		Machine.crafting_speed,
		Machine.polution,
		Machine.module_slot,
		Machine.q_coef_a,
		Machine.q_coef_b,
		Machine.q5_mod,
		Machine.drain,
		Machine.energy_consumption)

	if err !=nil{
		return 0,err
	}
	return result.LastInsertId()
}
