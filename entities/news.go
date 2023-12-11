package entities

import (
	"time"
)

type News struct {
	Uid				string `json:"id"`
	Title 			string `json:"title"`
	Description 	string `json:"desc"`
	UserId 			string `json:"user_id"`
	CreatedAt		time.Time `json:"created_at"`
	AppId 			string `json:"app_id"`
	AppName 		string `json:"app_name"`
	Application 	NewsApplicationResponse `json:"app"`
	User 			NewsUserResponse `json:"user"`
}

type NewsApplicationResponse struct {
	ApplicationId 	string `json:"id"`
	ApplicationName string `json:"name"`
}

type NewsImageForm struct {
	NewsId		string `json:"news_id"`
	Path 		string `json:"path"`
	Size		int    `json:"size"`
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
	Size		int    `json:"size"`
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