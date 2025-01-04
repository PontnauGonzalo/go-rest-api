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

// It initializes all dependencies and starts the HTTP server
func main() {
    // Load environment variables from .env file
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file: %v", err)
    }

    // db connection
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

    // Create a root context for the application
    ctx := context.Background()

    // Initialize HTTP handler with user endpoints
    h := handler.NewUserHTTPServer(user.MakeEndpoints(ctx, userService))

    // Configure server address from environment variables
    PORT := os.Getenv("PORT")
    address := fmt.Sprintf("localhost:%s", PORT)

    // Create and configure the HTTP server
    srv := &http.Server{
        Handler: enableCORS(h), // Enable CORS middleware
        Addr:    address,
    }

    fmt.Printf("Server listening on %s", PORT)

    log.Fatal(srv.ListenAndServe())
}

// enableCORS is a middleware that enables Cross-Origin Resource Sharing (CORS)
func enableCORS(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        
		// Set CORS headers
        w.Header().Set("Access-Control-Allow-Origin", "*")
        w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
        
        // Handle preflight requests
        if r.Method == http.MethodOptions {
            return
        }
        
        // Pass the request to the next handler
        next.ServeHTTP(w, r)
    })
}