package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/ernst01/learn-go-graph-ql/internal/songlist"
)

func main() {

	srv := songlist.Server{
		Router: mux.NewRouter(),
	}

	srv.Routes()

	srv.InitGraphQL()

	http.Handle("/", srv.Router)

	port := 8080

	if portStr := os.Getenv("APP_PORT"); portStr != "" {
		port, _ = strconv.Atoi(portStr)
	}

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
