package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	_ "github.com/mattn/go-sqlite3"
	"github.com/Niko-the-Useless/factoryDatabase/routes"
)


func main(){ 
	//connect to db
	db, err := sql.Open("sqlite3", "./data/factorydb")
	if err != nil{log.Fatal("cant connect to db: ",err)}
	//schedule close db
	defer db.Close()
	log.Println("Connected to db")
	//start server and register routes
	mux:=http.NewServeMux()
	routes.RegisterRoutes(mux,db)
	port:=":8080"
	log.Printf("Starting server on port: %s",port)
	err=http.ListenAndServe(port,mux)
	if err!=nil{fmt.Println("cant start server:",err)}

}
