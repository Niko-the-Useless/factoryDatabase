package lib

import (
	"fmt"
	"database/sql"
	"encoding/json"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "factory database home page")
}

func CreateProductsTableHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request){
		if r.Method != http.MethodPost{http.Error(w,"only post method allowed",http.StatusMethodNotAllowed);return}
		_, err:=CreateProductsTable(db)
		if err!=nil{http.Error(w, fmt.Sprintf("Creating product table failed: %v",err),http.StatusInternalServerError);return}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("table created"))

	}
}
func InsertProductHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request){
		
		if r.Method != http.MethodPost{http.Error(w,"only post method allowed",http.StatusMethodNotAllowed)
			return}

		var Product Product
		if json.NewDecoder(r.Body).Decode(&Product)!=nil{
			http.Error(w,"invalid JSON",http.StatusBadRequest)
			return}

		id,err:=InsertProduct(db,&Product)
		if err!=nil{http.Error(w,fmt.Sprintf("Cant insert product :%v",err),http.StatusInternalServerError)
			return}
		
		w.WriteHeader(http.StatusCreated)
		fmt.Fprint(w,fmt.Sprintf("Product inserted with id: %d",id))
	}
}

func CreateMachinesTableHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request){
		
		if r.Method != http.MethodPost{http.Error(w,"only post method allowed",http.StatusMethodNotAllowed)
			return}
	
		_, err:=CreateMachinesTable(db)
		if err!=nil{http.Error(w, fmt.Sprintf("Creating machines table failed: %v",err),http.StatusInternalServerError)
			return}
	
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w,"table created")
	}
}

func InsertMachineHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request){
		
		if r.Method != http.MethodPost{http.Error(w,"only post method allowed",http.StatusMethodNotAllowed)
			return}

		var Machine Machine
		if json.NewDecoder(r.Body).Decode(&Machine)!=nil{
			http.Error(w,"invalid JSON",http.StatusBadRequest)
			return}

		id,err:=InsertMachine(db,&Machine)
		if err!=nil{http.Error(w,fmt.Sprintf("Cant insert product :%v",err),http.StatusInternalServerError)
			return}
		
		w.WriteHeader(http.StatusCreated)
		fmt.Fprint(w,fmt.Sprintf("Machine inserted with id: %d",id))
	}
}
