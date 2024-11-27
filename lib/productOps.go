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

func (product Product) GetProductId (db *sql.DB) (int64, error){
	var(
		err error
		id int64
	)

	if product.Id!=nil{
		id=*product.Id}

	if product.Name!=nil{
		sql:=`SELECT id FROM products WHERE name=?`
		err=db.QueryRow(sql,product.Name).Scan(&id)
		if err!=nil{return 0, fmt.Errorf("Cant find product: %v",err)}}

	return id,nil
}

func (product Product) DeleteProduct(db *sql.DB) (int64, error){
	var (
		result sql.Result
		err error
	)

	id,err:=product.GetProductId(db)
	if err!=nil{return 0,err}

	sql :=`DELETE * FROM products WHERE id=?`
	result, err=db.Exec(sql,id)
	if err !=nil{return 0,err}
	return result.RowsAffected()
}

func (p *Product) GetProduct(db *sql.DB) (error){
	var err error

	id,err:=p.GetProductId(db)
	if err!=nil{return err}

	sql := `SELECT * FROM products WHERE id=?`
	row := db.QueryRow(sql,id)
	err = row.Scan(&p.Id, &p.Name, &p.Production_time)
	if err!=nil{return err}
	return nil
}
