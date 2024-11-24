package lib

import (
	"database/sql"
	"fmt"
)

func CreateProductsTable(db *sql.DB) (sql.Result, error){
	sql := `CREATE TABLE IF NOT EXISTS products (
		id INTEGER PRIMARY KEY,
		name TEXT UNIQE NOT NULL,
		production_time REAL
	);`
	return db.Exec(sql)
}

func (Product *Product) InsertProduct(db *sql.DB) (int64, error){
	sql :=`INSERT INTO products (
		name,
		production_time)
		VALUES (?,?);`

	result, err :=db.Exec(sql,
		Product.Name,
		Product.Production_time)

	if err !=nil{
		return 0,err
	}
	return result.LastInsertId()
}

func (target Target) DeleteProduct(db *sql.DB) (int64, error){
	var(
		result sql.Result
		err error
	)

	if target.Id!=nil{
			sql :=`DELETE FROM products WHERE id=?`
			result, err=db.Exec(sql,target.Id)
	}else if target.Name!=nil{
			sql :=`DELETE FROM products WHERE name=?`
			result, err=db.Exec(sql,target.Name)
	}else{
			return 0, fmt.Errorf("wrong argument type: supported types int-id string-name")}
	
	if err !=nil{return 0,err}
	return result.RowsAffected()
}
