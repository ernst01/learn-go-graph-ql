package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/ernst01/learn-go-graph-ql/internal/songlist"
)

func main() {

	srv := songlist.Server{
		Router: mux.NewRouter(),
	}

	srv.Routes()

	http.Handle("/", srv.Router)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
