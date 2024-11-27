package lib

import (
	"database/sql"
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
	var (
		result sql.Result
		err error
	)

	id,err:=target.GetId(db)
	if err!=nil{return 0,err}

	sql :=`DELETE FROM products WHERE id=?`
	result, err=db.Exec(sql,id)
	if err !=nil{return 0,err}
	return result.RowsAffected()
}

func (target Target) GetProduct(db *sql.DB) (Product, error){
	var (
		result sql.Result
		err error
	)
	product:=&Product{}

	id,err:=target.GetId(db)
	if err!=nil{return *product,err}

	sql :=`DELETE FROM products WHERE id=?`
	row :=db.QueryRow(db,id)
	return *product,nil
}

