package handlers

import (
	"net/http"

	"github.com/idkwhyureadthis/agg-project/pkg/fmt_respond"
)

func (apiCfg *APIConfig) HandlerReadiness(w http.ResponseWriter, r *http.Request) {
	fmt_respond.JsonifyRespond(w, 200, struct{}{})
}
