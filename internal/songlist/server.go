package songlist

import "github.com/gorilla/mux"

// Server represents our server
type Server struct {
	Router *mux.Router
}
