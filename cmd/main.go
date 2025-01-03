package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/PontnauGonzalo/go-rest-api/internal/user"
	"github.com/PontnauGonzalo/go-rest-api/pkg/boostrap"
	"github.com/PontnauGonzalo/go-rest-api/pkg/handler"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	db, err := boostrap.NewDB()

	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	logger := boostrap.NewLogger()
	userRepository := user.NewRepository(db, logger)
	userService := user.NewService(logger, userRepository)

	ctx := context.Background()

	h := handler.NewUserHTTPServer(user.MakeEndpoints(ctx, userService))

	PORT := os.Getenv("PORT")
	address := fmt.Sprintf("localhost:%s", PORT)

	srv := &http.Server{
		Handler: enableCORS(h),
		Addr:    address,
	}

	fmt.Printf("Server listening on %s", PORT)

	log.Fatal(srv.ListenAndServe())
}


func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if r.Method == http.MethodOptions {
			return
		}
		next.ServeHTTP(w, r)
	})
}
