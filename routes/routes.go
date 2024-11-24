package routes

import(
	"net/http"
	"github.com/Niko-the-Useless/factoryDatabase/lib"
	"database/sql"
)

func RegisterRoutes(mux *http.ServeMux, db *sql.DB){
	mux.HandleFunc("/",lib.HomeHandler)
	mux.HandleFunc("/create-product-table",lib.CreateProductsTableHandler(db))
	mux.HandleFunc("/create-machines-table",lib.CreateMachinesTableHandler(db))
}
