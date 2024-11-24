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

func InsertProduct(db *sql.DB, Product *Product) (int64, error){
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
func DeleteProduct(db *sql.DB, target Target) (int64, error){
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
/*
func UpdateProduct(db *sql.DB, target interface{}, ProductUpdate ProductUpdate, newValue interface{}) (int64, error){
	var(
		result sql.Result
		err error
	)
}
func LoadProductsCSV(filename string) ([]Product, error) {

    file, err := os.Open(filename)
    if err!=nil{return nil, err}
    defer file.Close()

    reader := csv.NewReader(file)
    data, err := reader.ReadAll()
    if err!=nil{return nil, err}

    var products []Product
    for _, dat := range data[1:] { 
        ProdTime, err := strconv.Atoi(dat[1])
        if err!=nil{return nil, err}

				products=append(products,Product{
					Name:dat[0],
					Production_time: ProdTime,
				})
    }
    return products, nil
}*/
