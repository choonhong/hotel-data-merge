package utils

import (
	"encoding/json"
	"net/http"
)

func ResponseHelper(w http.ResponseWriter, res any, status int) {
	w.Header().Set("Content-Type", "application/json")
	jsonOutput, _ := json.Marshal(res) //nolint:errchkjson
	w.WriteHeader(status)
	_, _ = w.Write(jsonOutput)
}
