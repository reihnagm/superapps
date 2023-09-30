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

	dir, _ := os.Open("public")

	fileInfos, _ := dir.Readdir(-1)

	for _, fileInfo := range fileInfos {
		if fileInfo.IsDir() {
			router.PathPrefix("/"+fileInfo.Name()+"/").Handler(http.StripPrefix("/"+fileInfo.Name()+"/", http.FileServer(http.Dir("./public/"+ fileInfo.Name() +"/"))))
		}
	}
	
	// Serving static

    // Auth
	router.HandleFunc("/api/v1/login", controllers.Login).Methods("POST")
	router.HandleFunc("/api/v1/register", controllers.Register).Methods("POST")

	// Otp
	router.HandleFunc("/api/v1/verify-otp", controllers.VerifyOtp).Methods("POST")

	// News
	router.HandleFunc("/api/v1/news", controllers.All).Methods("GET")
	router.HandleFunc("/api/v1/news", controllers.CreateNews).Methods("POST")

	// Media 
	router.HandleFunc("/api/v1/media/upload", controllers.Upload).Methods("POST")

	port := os.Getenv("PORT")
	handler := router
	server := new(http.Server)
	server.Handler = handler
	server.Addr = ":" + port
	fmt.Println("Starting server at", server.Addr)
    server.ListenAndServe()
}