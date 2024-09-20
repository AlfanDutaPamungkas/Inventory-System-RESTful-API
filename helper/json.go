package helper

import (
	"encoding/json"
	"net/http"
)

func BodyToReq(r *http.Request, result any) {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(result)
	PanicError(err)
}

func WriteToBody(w http.ResponseWriter, response any) {
	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	err := encoder.Encode(response)
	PanicError(err)
}
