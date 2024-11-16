package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)


func main(){ 
	//connect to db
	db, err := sql.Open("sqlite3", "./factorydb")
	if err != nil{
		fmt.Println(err)
		return
	}

	//schedule close db
	defer db.Close()

	fmt.Println("Connected to db")

	// get sqlute ver
	var sqliteVer string
	err = db.QueryRow("SELECT sqlite_version()").Scan(&sqliteVer)
	if err != nil{
		fmt.Println(err)
		return
	}
	fmt.Println(sqliteVer)
	
	//create products table
	_, err=createProductsTable(db)
	if err !=nil{
		fmt.Println(err)
		return
	}
	fmt.Println("created products table ^^")

	//create machines table
	_, err=createMachinesTable(db)
	if err !=nil{
		fmt.Println(err)
		return
	}
	fmt.Println("created machines table ^^")

	//create BOM table
	_, err=createBomTable(db)
	if err !=nil{
		fmt.Println(err)
		return
	}
	fmt.Println("created BOM table ^^")

	//example data
	productA := &Product{
		name:"productA",
		production_time: 2.5,
	}
	machineA := &Machine{
		name:"machineA",
		crafting_speed: 1,
		polution: 2.2,
		module_slot: 1,
		q_coef_a: 0.2,
		q_coef_b: 0.5,
		q5_mod: 2,
		drain: 2.5,
		energy_consumption: 75,
	}

	// inserting data
	productId, err := insertProduct(db, productA)
	if err != nil{
		fmt.Println(err)
		return
	}
	machineId, err := insertMachine(db, machineA)
	if err != nil{
		fmt.Println(err)
		return
	}

	//show inserted data
	fmt.Printf("Product %s was inserted with ID: %d\n",
		productA.name,productId,	
	)
	fmt.Printf("Machine %s was inserted with ID: %d\n",
		machineA.name,machineId,	
	)

}
