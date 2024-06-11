package handlers

import (
	"net/http"

	"github.com/idkwhyureadthis/agg-project/internal/database"
)

type APIConfig struct {
	DB *database.Queries
}

type authedHandler func(http.ResponseWriter, *http.Request, database.User)
