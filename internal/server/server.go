package server

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"mymod/internal/api"
	"mymod/internal/data"
	"net/http"
)

func Server() {
	fmt.Println("Server run...")
	r := mux.NewRouter()
	r.HandleFunc("/", handleConnection)
	log.Fatal(http.ListenAndServe(data.UrlServer, r))
}

func handleConnection(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write(api.GetApi())
}
