package entities

import (
	"time"
)

type NewsFormEntity struct {
	Uid				string `json:"id"`
	Title 			string `json:"title"`
	Description 	string `json:"desc"`
	UserId 			string `json:"user_id"`
	CreatedAt		time.Time `json:"created_at"`
	ApplicationId 	string `json:"app_id"`
	ApplicationName string `json:"app_name"`
	Application 	ApplicationResponseEntity `json:"app"`
	User 			UserResponseEntity `json:"user"`
}

type ApplicationResponseEntity struct {
	ApplicationId string `json:"id"`
	Name 		  string `json:"name"`
}

type AllCountNewsEntity struct {
	Uid			  string `json:"id"`
}

type NewsResponseEntity struct {
	Uid			string `json:"id"`
	Title 		string `json:"title"`
	Description	string `json:"desc"`
	App 		ApplicationResponseEntity `json:"app"`
	User 		UserResponseEntity `json:"user"`
	CreatedAt   string `json:"created_at"`
}

type UserEntity struct {
	Fullname string `json:"fullname"`
	Email 	 string `json:"email"`
	Phone 	 string `json:"phone"`
}

type UserResponseEntity struct {
	Fullname string `json:"fullname"`
	Email 	 string `json:"email"`
	Phone 	 string `json:"phone"`
}