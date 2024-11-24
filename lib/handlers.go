package lib

import (
	"fmt"
	"database/sql"
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
