package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/idkwhyureadthis/agg-project/internal/database"
	"github.com/idkwhyureadthis/agg-project/pkg/fmt_respond"
	"github.com/idkwhyureadthis/agg-project/pkg/models"
)

func (apiCfg *APIConfig) HandlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		fmt_respond.ErrorifyRespond(w, 400, fmt.Sprintf("Error parsing JSON: %s", err))
		return
	}

	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		Name:      params.Name,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	})
	if err != nil {
		fmt_respond.ErrorifyRespond(w, 400, fmt.Sprintf("failed to create user: %s", err))
		return
	}

	fmt_respond.JsonifyRespond(w, 201, models.DBUserToUser(user))

}

func (apiCfg *APIConfig) HandlerGetUser(w http.ResponseWriter, r *http.Request, user database.User) {
	fmt_respond.JsonifyRespond(w, 200, models.DBUserToUser(user))
}
