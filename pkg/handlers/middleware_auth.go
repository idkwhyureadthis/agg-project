package handlers

import (
	"fmt"
	"net/http"

	"github.com/idkwhyureadthis/agg-project/internal/auth"
	"github.com/idkwhyureadthis/agg-project/internal/database"
	"github.com/idkwhyureadthis/agg-project/pkg/fmt_respond"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *APIConfig) MiddlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			fmt_respond.ErrorifyRespond(w, 403, fmt.Sprintf("Auth error: %v", err))
			return
		}

		user, err := apiCfg.DB.GetUserByKey(r.Context(), apiKey)
		if err != nil {
			fmt_respond.ErrorifyRespond(w, 400, fmt.Sprintf("Couldn't get user: %v", err))
			return
		}
		handler(w, r, user)
	}
}
