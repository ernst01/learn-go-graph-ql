package songlist

import (
	"net/http"

	"github.com/graphql-go/graphql"
)

func (s *Server) appVersion() http.HandlerFunc {
	type version struct {
		Name    string `json:"name"`
		Version string `json:"version"`
		Note    string `json:"note"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		sendResponse(w, http.StatusOK, &version{
			"learn-go-graph-ql",
			"0.0.1-beta1",
			"Let's get ready to rumble!!",
		})
	}
}

func (s *Server) readSonglistHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query().Get("query")
		result := graphql.Do(graphql.Params{
			Schema:        SongListSchema,
			RequestString: query,
		})
		if result.HasErrors() {
			sendResponse(w, http.StatusUnprocessableEntity, result.Errors)
			return
		}
		sendResponse(w, http.StatusOK, result.Data)
	}
}
