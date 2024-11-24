package main

import "net/http"

func main() {

 mux := http.NewServeMux()
 mux.Handle("/", &homeHandler{})
 mux.Handle("/no",&noHandler{})

 http.ListenAndServe(":8080", mux)
}

type homeHandler struct{}

func (hh *homeHandler) ServeHTTP(writer http.ResponseWriter, requester *http.Request) {
 writer.Write([]byte("wwwwwww"))
}

type noHandler struct{}
func (nh *noHandler) ServeHTTP(writer http.ResponseWriter, r *http.Request){
		writer.Write([]byte("aaaaa"))
}
