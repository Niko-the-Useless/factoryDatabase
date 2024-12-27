package lib

import (
	"fmt"
	"database/sql"
	"encoding/json"
	"net/http"
)

func CreateBomTableHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request){
		if r.Method != http.MethodPost{http.Error(w,"only post method allowed",http.StatusMethodNotAllowed);return}
		_, err:=CreateBomTable(db)
		if err!=nil{http.Error(w, fmt.Sprintf("Creating product table failed: %v",err),http.StatusInternalServerError);return}
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "Created bom table")
	}
}
func InsertBomHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		
		if r.Method != http.MethodPost {
			http.Error(w, "only POST method allowed", http.StatusMethodNotAllowed)
			return
		}

		var bom BOM
		if err := json.NewDecoder(r.Body).Decode(&bom); err != nil {
			http.Error(w, "invalid JSON", http.StatusBadRequest)
			return
		}

		if bom.Child_id == nil || bom.Child_quantity == nil {
			http.Error(w, "child_id and child_quantity are required", http.StatusBadRequest)
			return
		}

		id, err := bom.InsertBOM(db)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to insert BOM: %v", err), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, "BOM entry inserted with last inserted row ID: %d", id)
	}
}

/*
func GetProductIdHandler(db *sql.DB) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){

		var product Product

		if r.Method!=http.MethodGet{http.Error(w,"only get method allowed",http.StatusMethodNotAllowed)
			return}
		
		if json.NewDecoder(r.Body).Decode(&product)!=nil{
			http.Error(w,"invalid JSON",http.StatusBadRequest)
			return
		}

		if product.Name==nil&&product.Id==nil{
			http.Error(w,"provide Name or Id ",http.StatusBadRequest)
			return}

		id,err:=product.GetProductId(db)
		if err!=nil{http.Error(w,fmt.Sprintf("Cant get id: %v",err),http.StatusInternalServerError);return}
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w,"Id: ",id)
	}
}

func DeleteProductHandler(db *sql.DB) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){

		var product Product
		
		if r.Method != http.MethodDelete{http.Error(w,"only post method allowed",http.StatusMethodNotAllowed)
			return}

		if json.NewDecoder(r.Body).Decode(&product)!=nil{
			http.Error(w,"invalid JSON",http.StatusBadRequest)
			return}

		if product.Name==nil&&product.Id==nil{
			http.Error(w,"provide Name or Id ",http.StatusBadRequest)
			return}

		id,err:=product.DeleteProduct(db)
		if err!=nil{http.Error(w,fmt.Sprintf("Cant delete product :%v",err),http.StatusInternalServerError)
			return}
		
		w.WriteHeader(http.StatusNoContent)
		fmt.Fprint(w,fmt.Sprintf("Product with id: %d was deleted",id))
	}
}

func GetProductHandler(db *sql.DB) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
	
		var product Product
		w.Header().Set("Content-Type","application/json")
		
		if r.Method != http.MethodGet{http.Error(w,"only get method allowed",http.StatusMethodNotAllowed)
			return}

		if json.NewDecoder(r.Body).Decode(&product)!=nil{
			http.Error(w,"invalid JSON",http.StatusBadRequest)
			return}

		if product.Name==nil&&product.Id==nil{
			http.Error(w,"provide Name or Id ",http.StatusBadRequest)
			return}

		err:=product.GetProduct(db)
		if err!=nil{http.Error(w,fmt.Sprintf("cant find product: %v",err),
			http.StatusInternalServerError)
		return}

		if json.NewEncoder(w).Encode(product)!=nil{http.Error(w,"failed to encode json ",http.StatusInternalServerError)
		return}
	}
}

func UpdateProductHandler(db *sql.DB) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
	
		var newProduct Product
		w.Header().Set("Content-Type","application/json")
		
		if r.Method != http.MethodPatch{http.Error(w,"only patch method allowed",http.StatusMethodNotAllowed)
			return}

		if json.NewDecoder(r.Body).Decode(&newProduct)!=nil{
			http.Error(w,"invalid JSON",http.StatusBadRequest)
			return}

		if newProduct.Name==nil&&newProduct.Id==nil{
			http.Error(w,"provide Name or Id ",http.StatusBadRequest)
			return}

		prodId,err:=newProduct.UpdateProduct(db)
		if err!=nil{http.Error(w,fmt.Sprintf("cant update product %v",err),http.StatusInternalServerError)
			return}

		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w,fmt.Sprintf("product with id %d updated",prodId))
	}
}
*/
