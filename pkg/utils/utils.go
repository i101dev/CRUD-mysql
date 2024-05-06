package utils

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func ParseBody(r *http.Request, x interface{}) {

	body, err := io.ReadAll(r.Body)

	if err != nil {
		log.Fatalf("Error reading request body: %s", err)
		return
	}

	if err := json.Unmarshal([]byte(body), x); err != nil {
		log.Fatalf("Error unmarshalling JSON: %s", err)
		return
	}
}
