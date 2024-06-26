package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/idkwhyureadthis/agg-project/internal/database"
	"github.com/idkwhyureadthis/agg-project/pkg/handlers"
	"github.com/idkwhyureadthis/agg-project/pkg/routers"
	"github.com/idkwhyureadthis/agg-project/pkg/scraping"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	godotenv.Load()
	router := chi.NewRouter()
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("Cannot find DB_URL in environment")
	}

	conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("failed to open a database: ", err)
	}

	apiCfg := handlers.APIConfig{
		DB: database.New(conn),
	}

	go scraping.StartScraping(apiCfg.DB, 10, time.Minute)

	portString := os.Getenv("PORT")
	if portString == "" {
		portString = "8080"
	}
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"POST", "PUT", "GET", "DELETE", "OPTIONS", "PATCH"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Router := routers.SetupV1Router(apiCfg)
	router.Mount("/v1", v1Router)

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}
	log.Printf("Server starting at :%v\n", portString)
	log.Fatal(srv.ListenAndServe())
}
