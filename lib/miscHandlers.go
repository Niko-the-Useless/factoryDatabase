package lib

import (
	"fmt"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "factory database home page")
}
