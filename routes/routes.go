package routes

import(
	"net/http"
	"github.com/Niko-the-Useless/factoryDatabase/lib"
	"database/sql"
)

func RegisterRoutes(mux *http.ServeMux, db *sql.DB){
	mux.HandleFunc("/",lib.HomeHandler)
	mux.HandleFunc("/create-product-table",lib.CreateProductsTableHandler(db))
	mux.HandleFunc("/insert-product",lib.InsertProductHandler(db))
	mux.HandleFunc("/delete-product",lib.DeleteProductHandler(db))
	mux.HandleFunc("/create-machines-table",lib.CreateMachinesTableHandler(db))
	mux.HandleFunc("/insert-machine",lib.InsertMachineHandler(db))
	mux.HandleFunc("/delete-machine",lib.DeleteMachineHandler(db))
}
