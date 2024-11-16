package main

import (
	"database/sql"
)

func createProductTable(db *sql.DB) (sql.Result, error){
	sql := `CREATE TABLE IF NOT EXISTS product (
		id INTEGER PRIMARY KEY,
		name TEXT,
		production_time REAL
	);`
	return db.Exec(sql)
}

func createMachineTable(db *sql.DB) (sql.Result, error){
	sql := `CREATE TABLE IF NOT EXISTS product (
		id INTEGER PRIMARY KEY,
		name TEXT,
		crafting_speed REAL,
		polution REAL,
		module_slots INTEGER,
		quality_coefficient_a REAL,
		quality_coefficient_b REAL,
		quality5_modifier REAL,
		drain REAL,
		energy_consumption REAL
	);`
	return db.Exec(sql)
}
// parent = product produced
//child = product requred
func createProductBomTable(db *sql.DB) (sql.Result, error){
	sql := `CREATE TABLE IF NOT EXISTS product (
		parent_id INTEGER NOT NULL,
		parent_quantity INTEGER,
		child_id INTEGER,
		child_quantity INTEGER,
		byproduct_id INTEGER,
		byproduct_quantity INTEGER
		);`
	return db.Exec(sql)
}
