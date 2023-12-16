package entities

import (
	"time"
)

type Content struct {
	Uid				string `json:"id"`
	Title 			string `json:"title"`
	Description 	string `json:"desc"`
	UserId 			string `json:"user_id"`
	CreatedAt		time.Time `json:"created_at"`
	AppId 			string `json:"app_id"`
	AppName 		string `json:"app_name"`
	Application 	ContentApplicationResponse `json:"app"`
	User 			ContentUserResponse `json:"user"`
}

type ContentApplicationResponse struct {
	ApplicationId 	string `json:"id"`
	ApplicationName string `json:"name"`
}

type ContentMediaForm struct {
	ContentId	string `json:"content_id"`
	Path 		string `json:"path"`
	Size		int    `json:"size"`
}

type AllCountContent struct {
	Uid			  string `json:"id"`
}

type ContentResponse struct {
	Uid			string `json:"id"`
	Title 		string `json:"title"`
	Description	string `json:"desc"`
	File		[]ContentMediaResponse `json:"files"`
	App 		ContentApplicationResponse `json:"app"`
	User 		ContentUserResponse `json:"user"`
	CreatedAt   string `json:"created_at"`
}

type ContentMediaResponse struct {
	Path 		string `json:"path"`
	Size		int    `json:"size"`
}

type ContentUser struct {
	Fullname string `json:"fullname"`
	Email 	 string `json:"email"`
	Phone 	 string `json:"phone"`
}

type ContentUserResponse struct {
	Fullname string `json:"fullname"`
	Email 	 string `json:"email"`
	Phone 	 string `json:"phone"`
}