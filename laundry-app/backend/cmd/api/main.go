package main

import (
	"log"
	"net/http"
	"os"

	"laundry/internal/db"
	"laundry/internal/handler"
	"laundry/internal/repository"
	"laundry/internal/service"
)

func main() {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "host=localhost user=postgres password=postgres dbname=laundry port=5432 sslmode=disable"
	}
	dbConn := db.NewPostgres(dsn)

	userRepo := repository.NewUserRepository(dbConn)
	authService := service.NewAuthService(userRepo)
	authHandler := handler.NewAuthHandler(authService)

	transactionRepo := repository.NewTransactionRepository(dbConn)
	transactionService := service.NewTransactionService(transactionRepo)
	transactionHandler := handler.NewTransactionHandler(transactionService)

	mux := http.NewServeMux()
	mux.HandleFunc("/api/login", authHandler.Login)
	mux.HandleFunc("/api/google-login", authHandler.GoogleLogin)
	mux.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("API Running"))
	})

	mux.HandleFunc("/api/transactions", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			transactionHandler.Create(w, r)
		case http.MethodGet:
			transactionHandler.GetByEmailAndDate(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	finalHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		if origin == "" { origin = "*" }
		w.Header().Set("Access-Control-Allow-Origin", origin)
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		mux.ServeHTTP(w, r)
	})

	port := os.Getenv("PORT")
	if port == "" { port = "4000" }

	log.Printf("API running at http://0.0.0.0:%s", port)
	log.Fatal(http.ListenAndServe(":"+port, finalHandler))
}
