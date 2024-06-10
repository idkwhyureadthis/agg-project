package fmt_respond

import (
	"log"
	"net/http"
)

func ErrorifyRespond(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Println("responsed with 5XX error: ", msg)
	}

	type errResponse struct {
		Error string `json:"error"`
	}

	JsonifyRespond(w, code, errResponse{
		Error: msg,
	})
}
