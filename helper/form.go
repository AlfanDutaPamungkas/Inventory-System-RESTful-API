package helper

import (
	"net/http"

	"github.com/gorilla/schema"
)

func FormToReq(r *http.Request, result any) {
	var decoder = schema.NewDecoder()

	err := r.ParseForm()
	PanicError(err)

	err = decoder.Decode(result, r.PostForm)
	PanicError(err)
}
