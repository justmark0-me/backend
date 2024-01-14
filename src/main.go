package main

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
	"log"
	"main/src/handlers"
	"net/http"
	"os"
)

func main() {
	ctx := context.Background()
	// DATABASE_URL example DATABASE_URL=postgres://postgres:postgres@localhost:5432/postgres
	conn, err := pgx.Connect(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	mux := http.NewServeMux()

	backendHandlers := handlers.NewClient(conn)
	mux.HandleFunc("/monkeytype_background.jpg", backendHandlers.GetMonkeytypeBackground)
	mux.HandleFunc("/v1/new_request/", backendHandlers.SaveRequest)
	mux.HandleFunc("/", backendHandlers.GetRootPage)

	handler := corsMiddleware(mux)
	log.Println("Server starting on port 8080")
	err = http.ListenAndServe(":8080", handler)
	if err != nil {
		panic(errors.Wrap(err, "listen and serve"))
	}
}

// corsMiddleware adds CORS headers to the response
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		// If it's a preflight OPTIONS request, send a 200 response without further processing
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Otherwise, pass the request to the next handler
		next.ServeHTTP(w, r)
	})
}
