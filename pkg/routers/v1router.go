package routers

import (
	"github.com/go-chi/chi"
	"github.com/idkwhyureadthis/agg-project/pkg/handlers"
)

func SetupV1Router(apiCfg handlers.APIConfig) chi.Router {
	v1Router := chi.NewRouter()
	v1Router.Get("/healthz", apiCfg.HandlerReadiness)
	v1Router.Get("/err", apiCfg.HandlerError)
	v1Router.Post("/users", apiCfg.HandlerCreateUser)
	v1Router.Post("/feeds", apiCfg.MiddlewareAuth(apiCfg.HandlerCreateFeed))
	v1Router.Get("/feeds", apiCfg.HandlerGetFeeds)
	v1Router.Get("/users", apiCfg.MiddlewareAuth(apiCfg.HandlerGetUser))
	v1Router.Post("/feed_follows", apiCfg.MiddlewareAuth(apiCfg.HandlerCreateFeedFollow))
	v1Router.Get("/feed_follows", apiCfg.MiddlewareAuth(apiCfg.HandlerGetFeedFollows))
	v1Router.Delete("/feed_follows/{feedFollowId}", apiCfg.MiddlewareAuth(apiCfg.HandlerDeleteFeedFollows))
	v1Router.Get("/getposts", apiCfg.MiddlewareAuth(apiCfg.HandeGetPostsForUser))
	return v1Router
}
