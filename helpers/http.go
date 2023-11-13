package helpers

import (
	"encoding/json"
	"net/http"
)

func DecodeRequest(r *http.Request, requestVar interface{}) {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(requestVar)
	PanifIfError(err)
}
