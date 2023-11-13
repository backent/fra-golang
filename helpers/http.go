package helpers

import (
	"encoding/json"
	"net/http"

	"github.com/backent/fra-golang/web"
)

func DecodeRequest(r *http.Request, requestVar interface{}) {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(requestVar)
	PanifIfError(err)
}

func ReturnReponseJSON(w http.ResponseWriter, response web.WebResponse) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(response.Code)
	encoder := json.NewEncoder(w)
	encoder.Encode(response)
}
