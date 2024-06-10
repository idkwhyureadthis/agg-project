package handlers

import (
	"net/http"

	"github.com/idkwhyureadthis/agg-project/pkg/fmt_respond"
)

func HandlerError(w http.ResponseWriter, r *http.Request) {
	fmt_respond.ErrorifyRespond(w, 400, "something went wrong")
}
