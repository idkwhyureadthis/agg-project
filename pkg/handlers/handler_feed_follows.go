package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/idkwhyureadthis/agg-project/internal/database"
	"github.com/idkwhyureadthis/agg-project/pkg/fmt_respond"
	"github.com/idkwhyureadthis/agg-project/pkg/models"
)

func (apiCfg *APIConfig) HandlerCreateFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		FeedId uuid.UUID `json:"feed_id"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		fmt_respond.ErrorifyRespond(w, 400, fmt.Sprintf("Error parsing JSON: %s", err))
		return
	}

	feedFollow, err := apiCfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    params.FeedId,
	})
	if err != nil {
		fmt_respond.ErrorifyRespond(w, 400, fmt.Sprintf("failed to create feed follow: %s", err))
		return
	}

	fmt_respond.JsonifyRespond(w, 201, models.DBFeedFollowToFeedFollow(feedFollow))

}

func (apiCfg *APIConfig) HandlerGetFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollows, err := apiCfg.DB.GetFeedFollows(r.Context(), user.ID)
	if err != nil {
		fmt_respond.ErrorifyRespond(w, 400, fmt.Sprintf("failed to get feed follows:%s", err))
		return
	}
	fmt_respond.JsonifyRespond(w, 200, models.DBFeedFollowsToFeedFollows(feedFollows))
}

func (apiCfg *APIConfig) HandlerDeleteFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFolowIdStr := chi.URLParam(r, "feedFollowId")
	feedFolowId, err := uuid.Parse(feedFolowIdStr)
	if err != nil {
		fmt_respond.ErrorifyRespond(w, 400, fmt.Sprintf("Failed to parse passed id:%s", err))
		return
	}
	err = apiCfg.DB.DeleteFeedFollows(r.Context(), database.DeleteFeedFollowsParams{
		UserID: user.ID,
		ID:     feedFolowId,
	})
	if err != nil {
		fmt_respond.ErrorifyRespond(w, 400, fmt.Sprintf("Falied to delete follow: %s", err))
		return
	}
	fmt_respond.JsonifyRespond(w, 200, struct{}{})
}
