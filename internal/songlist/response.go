package songlist

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func sendSuccess(w http.ResponseWriter, status int, data interface{}) {
	var buf bytes.Buffer

	if nil != data {
		if err := json.NewEncoder(&buf).Encode(data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if _, err := io.Copy(w, &buf); err != nil {
		log.Println("response:", err)
	}
}
