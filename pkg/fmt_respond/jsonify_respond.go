package fmt_respond

import (
	"encoding/json"
	"log"
	"net/http"
)

func JsonifyRespond(w http.ResponseWriter, code int, payload interface{}) {
	dat, err := json.Marshal(payload)
	if err != nil {
		log.Printf("can't unmarshall JSON: %v", dat)
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(code)
	w.Header().Add("Content-Type", "application/json")
	w.Write(dat)
}
