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
//close db
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
}
