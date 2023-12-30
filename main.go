package main

import (
	helper "superapps/helpers"
	middleware "superapps/middlewares"
    "superapps/controllers"
    "net/http"
    "os"
	"time"
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
			// Serving static
			router.PathPrefix("/"+fileInfo.Name()+"/").Handler(http.StripPrefix("/"+fileInfo.Name()+"/", http.FileServer(http.Dir("./public/"+ fileInfo.Name() +"/"))))
		}
	}
	
    // Auth
	router.HandleFunc("/api/v1/login", controllers.Login).Methods("POST")
	router.HandleFunc("/api/v1/register", controllers.Register).Methods("POST")

	// Otp
	router.HandleFunc("/api/v1/resend-otp", controllers.ResendOtp).Methods("POST")
	router.HandleFunc("/api/v1/verify-otp", controllers.VerifyOtp).Methods("POST")

	// Content
	router.HandleFunc("/api/v1/content", controllers.All).Methods("GET")
	router.HandleFunc("/api/v1/content-upload", controllers.CreateMediaContent).Methods("POST")
	router.HandleFunc("/api/v1/content", controllers.CreateContent).Methods("POST")

	// Content Comment
	router.HandleFunc("/api/v1/content/comment", controllers.CreateContentComment).Methods("POST")
	router.HandleFunc("/api/v1/content/comment/delete", controllers.DeleteContentComment).Methods("POST")

	// Content Like
	router.HandleFunc("/api/v1/content/like", controllers.CreateContentLike).Methods("POST")

	// Content Unlike
	router.HandleFunc("/api/v1/content/unlike", controllers.CreateContentUnlike).Methods("POST")

	// Membernear
	router.HandleFunc("/api/v1/membernear/all", controllers.GetMembernear).Methods("GET")

	// Media 
	router.HandleFunc("/api/v1/media/upload", controllers.Upload).Methods("POST")

	portEnv := os.Getenv("PORT")
	port	:= ":" + portEnv

	// NOT SECURE FOR USE
	// server := new(http.Server)
	// server.Handler = router
	// server.Addr = ":" + port

	fmt.Println("Starting server at", port)

	server := &http.Server{
        Addr:              port,
		Handler:  		   router,
        ReadHeaderTimeout: 3 * time.Second,
    }

    errListenAndServe := server.ListenAndServe()
    if err != nil {
		fmt.Println("Error starting server:", errListenAndServe)
    }

    // errs := http.ListenAndServe(port, router)
	// if errs != nil {
	// 	fmt.Println("Error starting server:", errs)
	// }
}