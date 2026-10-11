package utils

import (
	"encoding/json"
	"net/http"
)

// MaxJSONBodyBytes caps JSON request bodies so a client can't exhaust memory.
const MaxJSONBodyBytes = 1 << 20 // 1 MB

func DecodeJSON(w http.ResponseWriter, r *http.Request, dst any) error {
	r.Body = http.MaxBytesReader(w, r.Body, MaxJSONBodyBytes)
	return json.NewDecoder(r.Body).Decode(dst)
}
