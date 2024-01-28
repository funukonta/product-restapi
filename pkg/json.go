package pkg

import (
	"encoding/json"
	"net/http"
)

// Func to decode from json Request
func DecodeJsonReq(r *http.Request, v any) error {
	err := json.NewDecoder(r.Body).Decode(v)
	defer r.Body.Close()
	return err
}

func WriteJsonRes(w http.ResponseWriter, statusCode int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	return json.NewEncoder(w).Encode(v)
}
