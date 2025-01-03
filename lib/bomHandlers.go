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
			return}

		var bom BOM
		if err := json.NewDecoder(r.Body).Decode(&bom); err != nil {
			http.Error(w, "invalid JSON", http.StatusBadRequest)
			return}

		if bom.Child_id == nil || bom.Child_quantity == nil {
			http.Error(w, "child_id and child_quantity are required", http.StatusBadRequest)
			return}

		id, err := bom.InsertBOM(db)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to insert BOM: %v", err), http.StatusInternalServerError)
			return}

		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, "BOM entry inserted with last inserted row ID: %d", id)
	}
}

func GetBomHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var bom BOM
		w.Header().Set("Content-Type", "application/json")

		if r.Method != http.MethodGet {
			http.Error(w, "only GET method allowed", http.StatusMethodNotAllowed)
			return}

		if json.NewDecoder(r.Body).Decode(&bom) != nil {
			http.Error(w, "invalid JSON", http.StatusBadRequest)
			return}

		if bom.Child_id == nil {
			http.Error(w, "child_id is required", http.StatusBadRequest)
			return}

		err := bom.GetBom(db, *bom.Child_id)
		if err != nil {
			http.Error(w, fmt.Sprintf("failed to retrieve BOM: %v", err), http.StatusInternalServerError)
			return}

		if json.NewEncoder(w).Encode(bom) != nil {
			http.Error(w, "failed to encode JSON", http.StatusInternalServerError)
			return}
	}
}

func DeleteBomHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var bom BOM
		w.Header().Set("Content-Type", "application/json")

		if r.Method != http.MethodDelete {
			http.Error(w, "only DELETE method allowed", http.StatusMethodNotAllowed)
			return}

		if json.NewDecoder(r.Body).Decode(&bom) != nil {
			http.Error(w, "invalid JSON", http.StatusBadRequest)
			return}

		if bom.Child_id == nil {
			http.Error(w, "child_id is required", http.StatusBadRequest)
			return}

		err := bom.DeleteBom(db, *bom.Child_id)
		if err != nil {
			http.Error(w, fmt.Sprintf("failed to delete BOM entries: %v", err), http.StatusInternalServerError)
			return}

		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "BOM entries deleted successfully")
	}
}

func UpdateBomHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var bom BOM
		w.Header().Set("Content-Type", "application/json")

		if r.Method != http.MethodPut {
			http.Error(w, "only PUT method allowed", http.StatusMethodNotAllowed)
			return}

		if json.NewDecoder(r.Body).Decode(&bom) != nil {
			http.Error(w, "invalid JSON", http.StatusBadRequest)
			return}

		if bom.Child_id == nil {
			http.Error(w, "child_id is required", http.StatusBadRequest)
			return}

		err := bom.DeleteBom(db, *bom.Child_id)
		if err != nil {
			http.Error(w, fmt.Sprintf("failed to delete BOM entries: %v", err), http.StatusInternalServerError)
			return}

		_, err = bom.InsertBOM(db)
		if err != nil {
			http.Error(w, fmt.Sprintf("failed to insert BOM entries: %v", err), http.StatusInternalServerError)
			return}

		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "BOM entries updated successfully")
	}
}
