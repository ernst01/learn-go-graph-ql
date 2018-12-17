package songlist

// Routes defines all the available routes
func (s *Server) Routes() {
	s.Router.Path("/").
		Methods("GET").
		Name("get-api-info").
		HandlerFunc(s.appVersion())
}
