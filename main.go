package main

import (
    "superapps/controllers"
	helper "superapps/helpers"
	middleware "superapps/middlewares"
    "net/http"
    "os"
    "fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		helper.Logger("error", "Error getting env")
	}

	router := mux.NewRouter()
	router.Use(middleware.JwtAuthentication)

    // Auth
	router.HandleFunc("/api/v1/login", controllers.Login).Methods("POST")
	router.HandleFunc("/api/v1/register", controllers.Register).Methods("POST")

	// Otp
	router.HandleFunc("/api/v1/verify-otp", controllers.VerifyOtp).Methods("POST")

	// News
	router.HandleFunc("/api/v1/news", controllers.All).Methods("GET")

	port := os.Getenv("PORT")
	handler := router
	server := new(http.Server)
	server.Handler = handler
	server.Addr = ":" + port
	fmt.Println("Starting server at", server.Addr)
    server.ListenAndServe()
}