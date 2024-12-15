package routes

import(
	"net/http"
	"github.com/Niko-the-Useless/factoryDatabase/lib"
	"database/sql"
)

func RegisterRoutes(mux *http.ServeMux, db *sql.DB){
	//product
	mux.HandleFunc("/product/create-table",lib.CreateProductsTableHandler(db))
	mux.HandleFunc("/product/insert",lib.InsertProductHandler(db))
	mux.HandleFunc("/product/delete",lib.DeleteProductHandler(db))
	mux.HandleFunc("/product/get-id",lib.GetProductIdHandler(db))
	mux.HandleFunc("/product/get",lib.GetProductHandler(db))
	mux.HandleFunc("/product/update",lib.UpdateProductHandler(db))
	//machine
	mux.HandleFunc("/machines/create-table",lib.CreateMachinesTableHandler(db))
	mux.HandleFunc("/machine/insert",lib.InsertMachineHandler(db))
	mux.HandleFunc("/machine/delete",lib.DeleteMachineHandler(db))
	mux.HandleFunc("/machine/get-id",lib.GetMachineIdHandler(db))
	mux.HandleFunc("/machinee/update",lib.UpdateProductHandler(db))
	//misc
	mux.HandleFunc("/",lib.HomeHandler)

}
