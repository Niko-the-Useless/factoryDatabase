package lib

import (
	"database/sql"
	"fmt"
)

func (target Target) GetId (db *sql.DB) (int64, error){
	var(
		err error
		id int64
	)

	if target.Id!=nil{
		id=*target.Id
	}
	if target.Name!=nil{
		sql:=`SELECT id FROM products WHERE name=?`
		err=db.QueryRow(sql,target.Name).Scan(&id)
		if err!=nil{return 0, fmt.Errorf("Cant find product: %v",err)}
	}
	return id,nil
}
