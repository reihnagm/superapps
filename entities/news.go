package entities

import (
	"time"
)

type NewsForm struct {
	Uid				string `json:"id"`
	Title 			string `json:"title"`
	Description 	string `json:"desc"`
	UserId 			string `json:"user_id"`
	CreatedAt		time.Time `json:"created_at"`
	ApplicationId 	string `json:"app_id"`
	ApplicationName string `json:"app_name"`
	Application 	NewsApplicationResponse `json:"app"`
	User 			NewsUserResponse `json:"user"`
}

type NewsImageForm struct {
	NewsId		string `json:"news_id"`
	Path 		string `json:"path"`
	Size		any    `json:"size"`
}

type NewsApplicationResponse struct {
	ApplicationId string `json:"id"`
	Name 		  string `json:"name"`
}

type AllCountNews struct {
	Uid			  string `json:"id"`
}

type NewsResponse struct {
	Uid			string `json:"id"`
	Title 		string `json:"title"`
	Description	string `json:"desc"`
	Images		[]NewsImageResponse `json:"images"`
	App 		NewsApplicationResponse `json:"app"`
	User 		NewsUserResponse `json:"user"`
	CreatedAt   string `json:"created_at"`
}

type NewsImageResponse struct {
	Path 		string `json:"path"`
	Size		any    `json:"size"`
}

type NewsUser struct {
	Fullname string `json:"fullname"`
	Email 	 string `json:"email"`
	Phone 	 string `json:"phone"`
}

type NewsUserResponse struct {
	Fullname string `json:"fullname"`
	Email 	 string `json:"email"`
	Phone 	 string `json:"phone"`
}