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
func GetIdHandler(db *sql.DB) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){

		var target Target

		if r.Method!=http.MethodGet{http.Error(w,"only get method allowed",http.StatusMethodNotAllowed)
			return}
		
		if json.NewDecoder(r.Body).Decode(&target)!=nil{
			http.Error(w,"invalid JSON",http.StatusBadRequest)
			return
		}

		if target.Name==nil&&target.Id==nil{
			http.Error(w,"provide Name or Id ",http.StatusBadRequest)
			return}

		id,err:=target.GetId(db)
		if err!=nil{http.Error(w,fmt.Sprintf("Cant get id: %v",err),http.StatusInternalServerError);return}
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w,"Id: ",id)
	}
}
