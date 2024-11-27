package lib

import (
	"fmt"
	"database/sql"
	"encoding/json"
	"net/http"
)
func CreateProductsTableHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request){
		if r.Method != http.MethodPost{http.Error(w,"only post method allowed",http.StatusMethodNotAllowed);return}
		_, err:=CreateProductsTable(db)
		if err!=nil{http.Error(w, fmt.Sprintf("Creating product table failed: %v",err),http.StatusInternalServerError);return}
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "Created product table")
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

		id,err:=Product.InsertProduct(db)
		if err!=nil{http.Error(w,fmt.Sprintf("Cant insert product :%v",err),http.StatusInternalServerError)
			return}
		
		w.WriteHeader(http.StatusCreated)
		fmt.Fprint(w,fmt.Sprintf("Product inserted with id: %d",id))
	}
}
func GetProductIdHandler(db *sql.DB) http.HandlerFunc{
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

		id,err:=target.GetMachineId(db)
		if err!=nil{http.Error(w,fmt.Sprintf("Cant get id: %v",err),http.StatusInternalServerError);return}
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w,"Id: ",id)
	}
}
func DeleteProductHandler(db *sql.DB) http.HandlerFunc{
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

		id,err:=Target.DeleteProduct(db)
		if err!=nil{http.Error(w,fmt.Sprintf("Cant delete product :%v",err),http.StatusInternalServerError)
			return}
		
		w.WriteHeader(http.StatusNoContent)
		fmt.Fprint(w,fmt.Sprintf("Product with id: %d was deleted",id))

	}
}
func GetProductHandler(db *sql.DB) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
	
		var Target Target
		w.Header().Set("Content-Type","application/json")
		
		if r.Method != http.MethodGet{http.Error(w,"only get method allowed",http.StatusMethodNotAllowed)
			return}

		if json.NewDecoder(r.Body).Decode(&Target)!=nil{
			http.Error(w,"invalid JSON",http.StatusBadRequest)
			return}

		if Target.Name==nil&&Target.Id==nil{
			http.Error(w,"provide Name or Id ",http.StatusBadRequest)
			return}

		product,err:=Target.GetProduct(db)
		if err!=nil{http.Error(w,fmt.Sprintf("cant find product: %v",err),
			http.StatusInternalServerError)
		return}

		if json.NewEncoder(w).Encode(product)!=nil{http.Error(w,"failed to encode json ",http.StatusInternalServerError)
		return}
	}
}
