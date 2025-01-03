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
func GetMachineHandler(db *sql.DB) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
	
		var machine Machine
		w.Header().Set("Content-Type","application/json")
		
		if r.Method != http.MethodGet{http.Error(w,"only get method allowed",http.StatusMethodNotAllowed)
			return}

		if json.NewDecoder(r.Body).Decode(&machine)!=nil{
			http.Error(w,"invalid JSON",http.StatusBadRequest)
			return}

		if machine.Name==nil&&machine.Id==nil{
			http.Error(w,"provide Name or Id ",http.StatusBadRequest)
			return}

		err:=machine.GetMachine(db)
		if err!=nil{http.Error(w,fmt.Sprintf("cant find machine: %v",err),
			http.StatusInternalServerError)
		return}

		if json.NewEncoder(w).Encode(machine)!=nil{http.Error(w,"failed to encode json ",http.StatusInternalServerError)
		return}
	}
}
func GetMachineIdHandler(db *sql.DB) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){

		var target Machine

		if r.Method!=http.MethodGet{http.Error(w,"only get method allowed",http.StatusMethodNotAllowed)
			return}
		
		if json.NewDecoder(r.Body).Decode(&target)!=nil{
			http.Error(w,"invalid JSON",http.StatusBadRequest)
			return
		}

		if target.Name==nil&&target.Id==nil{
			http.Error(w,"provide Name or Id ",http.StatusBadRequest)
			return}

		id,err:=target.GetMachineId(db)
		if err!=nil{http.Error(w,fmt.Sprintf("Cant get id: %v",err),http.StatusInternalServerError);return}
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w,"Id: ",id)
	}
}
func DeleteMachineHandler(db *sql.DB) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){

		var Target Machine
		
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

func UpdateMachineHandler(db *sql.DB) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		var newMach Machine
		w.Header().Set("Content-Type","application/json")

		if r.Method!=http.MethodPatch{http.Error(w,"only patch method allowed",http.StatusMethodNotAllowed)
		return}

		if json.NewDecoder(r.Body).Decode(&newMach)!=nil{http.Error(w,"invalid json",http.StatusBadRequest)
		return}

		if newMach.Name==nil&&newMach.Id==nil{
			http.Error(w,"provide name or id",http.StatusBadRequest)
			return}

		machId,err:=newMach.UpdateMachine(db)
		if err!=nil{http.Error(w,fmt.Sprintf("cant update machine %v",err),http.StatusInternalServerError)
		return}

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w,fmt.Sprintf("machine with id %d updated",machId))
	}
}
