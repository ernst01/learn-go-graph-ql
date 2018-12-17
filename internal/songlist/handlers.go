package songlist

import "net/http"

func (s *Server) appVersion() http.HandlerFunc {
	type version struct {
		Name    string `json:"name"`
		Version string `json:"version"`
		Note    string `json:"note"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		sendSuccess(w, http.StatusOK, &version{"test", "test", "test"})
	}
}
