package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"github.com/Niko-the-Useless/factoryDatabase/lib"
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
	_, err=lib.CreateProductsTable(db)
	if err !=nil{
		fmt.Println(err)
		return
	}
	fmt.Println("created products table ^^")

	//create machines table
	_, err=lib.CreateMachinesTable(db)
	if err !=nil{
		fmt.Println(err)
		return
	}
	fmt.Println("created machines table ^^")

	//create BOM table
	_, err=lib.CreateBomTable(db)
	if err !=nil{
		fmt.Println(err)
		return
	}
	fmt.Println("created BOM table ^^")

	//example data
	productA := &lib.Product{
		Name:"productA",
		Production_time: 2.5,
	}
	machineA := &lib.Machine{
		Name:"machineA",
		Crafting_speed: 1,
		Polution: 2.2,
		Module_slot: 1,
		Q_coef_a: 0.2,
		Q_coef_b: 0.5,
		Q5_mod: 2,
		Drain: 2.5,
		Energy_consumption: 75,
	}

	// inserting data
	productId, err := lib.InsertProduct(db, productA)
	if err != nil{
		fmt.Println(err)
		return
	}
	machineId, err := lib.InsertMachine(db, machineA)
	if err != nil{
		fmt.Println(err)
		return
	}

	//show inserted data
	fmt.Printf("Product %s was inserted with ID: %d\n",
		productA.Name,productId,	
	)
	fmt.Printf("Machine %s was inserted with ID: %d\n",
		machineA.Name,machineId,	
	)

	//delete product with id 1
	_,err = lib.DeleteProduct(db,1)
	if err!=nil{fmt.Println(err)
	return
	}
	fmt.Printf("Product %s was deleted with ID: %d\n",
		productA.Name,1,	
	)
	//deleta machine with name machineA
	_,err = lib.DeleteMachine(db,"machineA")
	if err!=nil{fmt.Println(err)
	return
	}
	fmt.Printf("Machine %s was deleted with ID: %d\n",
		machineA.Name,1,	
	)

}
