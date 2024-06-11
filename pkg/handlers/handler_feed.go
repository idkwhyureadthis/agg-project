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

func (apiCfg *APIConfig) HandlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		fmt_respond.ErrorifyRespond(w, 400, fmt.Sprintf("Error parsing JSON: %s", err))
		return
	}

	feed, err := apiCfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		Name:      params.Name,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Url:       params.URL,
		UserID:    user.ID,
	})
	if err != nil {
		fmt_respond.ErrorifyRespond(w, 400, fmt.Sprintf("failed to create feed: %s", err))
		return
	}

	fmt_respond.JsonifyRespond(w, 201, models.DBFeedToFeed(feed))

}

func (apiCfg *APIConfig) HandlerGetFeeds(w http.ResponseWriter, r *http.Request) {
	feeds, err := apiCfg.DB.GetFeeds(r.Context())
	if err != nil {
		fmt_respond.ErrorifyRespond(w, 400, fmt.Sprintf("failed to fetch feeds: %s", err))
		return
	}

	fmt_respond.JsonifyRespond(w, 200, models.DBFeedsToFeeds(feeds))
}
