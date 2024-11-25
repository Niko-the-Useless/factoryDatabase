package lib

import (
	"fmt"
	"database/sql"
	"encoding/json"
	"net/http"
)

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

		id,err:=Machine.InsertMachine(db)
		if err!=nil{http.Error(w,fmt.Sprintf("Cant insert product :%v",err),http.StatusInternalServerError)
			return}
		
		w.WriteHeader(http.StatusCreated)
		fmt.Fprint(w,fmt.Sprintf("Machine inserted with id: %d",id))
	}
}
func DeleteMachineHandler(db *sql.DB) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){

		var Target Target
		
		if r.Method != http.MethodPost{http.Error(w,"only post method allowed",http.StatusMethodNotAllowed)
			return}

		if json.NewDecoder(r.Body).Decode(&Target)!=nil{
			http.Error(w,"invalid JSON",http.StatusBadRequest)
			return}

		if Target.Name==nil&&Target.Id==nil{
			http.Error(w,"provide Name or Id ",http.StatusBadRequest)
			return}

		id,err:=Target.DeleteMachine(db)
		if err!=nil{http.Error(w,fmt.Sprintf("Cant delete machine :%v",err),http.StatusInternalServerError)
			return}
		w.WriteHeader(http.StatusNoContent)
		fmt.Fprint(w,fmt.Sprintf("Machine with id: %d was deleted",id))
	}
}
