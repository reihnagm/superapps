package entities

import (
	"time"
)

type News struct {
	Id 	  			string `json:"id"`
	Title 			string `json:"title"`
	Desc  			string `json:"desc"`
	ApplicationId	string `json:"application_id"`
	CreatedAt		time.Time `json:"created_at"`
	UpdateAt        time.Time `json:"updated_at"` 
}