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
	var	id int64

	if product.Id!=nil{
		id=*product.Id}

	if product.Name!=nil{
		sql:=`SELECT id FROM products WHERE name=?`
		err:=db.QueryRow(sql,product.Name).Scan(&id)
		if err!=nil{return 0, fmt.Errorf("Cant find product: %v",err)}}

	return id,nil
}

func (product Product) DeleteProduct(db *sql.DB) (int64, error){

	id,err:=product.GetProductId(db)
	if err!=nil{return 0,err}

	sql :=`DELETE * FROM products WHERE id=?`
	result, err:=db.Exec(sql,id)
	if err !=nil{return 0,err}
	return result.RowsAffected()
}

func (p *Product) GetProduct(db *sql.DB) (error){

	id,err:=p.GetProductId(db)
	if err!=nil{return err}

	sql := `SELECT * FROM products WHERE id=?`
	row := db.QueryRow(sql,id)
	err = row.Scan(&p.Id, &p.Name, &p.Production_time)
	if err!=nil{return err}
	return nil
}

func (newProd *Product) UpdateProduct(db *sql.DB) (int64, error){
	prod:=newProd
	prod.GetProduct(db)

	if newProd.Name!=nil{prod.Name=newProd.Name}
	if newProd.Production_time!=nil{prod.Production_time=newProd.Production_time}

	sql:=`UPDATE products SET name=?, production_time=?`
	res,err:=db.Exec(sql, prod.Name, prod.Production_time)
	if err!=nil{return 0,fmt.Errorf("cant update product: %v", err)}
	return res.RowsAffected()
}

