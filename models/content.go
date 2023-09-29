package models

import (
	"strconv"
	"math"
	"errors"
	helper "superapps/helpers"
)

type ApplicationEntity struct {
	ApplicationId string `json:"id"`
	Name 		  string `json:"name"`
}

type AllNewsEntity struct {
	Uid			  string `json:"id"`
}

type NewsEntity struct {
	Uid				string `json:"id"`
	Title 			string `json:"title"`
	Description 	string `json:"desc"`
	UserId 			string `json:"user_id"`
	ApplicationId 	string `json:"app_id"`
	ApplicationName string `json:"app_name"`
	Application ApplicationEntity `json:"app"`
	User UserEntity `json:"user"`
}

type NewsEntityResponse struct {
	Uid			string `json:"id"`
	Title 		string `json:"title"`
	Description string `json:"desc"`
	App ApplicationEntity `json:"app"`
	User UserEntity `json:"user"`
}

type UserEntity struct {
	Fullname string `json:"fullname"`
	Email 	 string `json:"email"`
	Phone 	 string `json:"phone"`
}

func (n *News) GetNews(page, limit string) (map[string]interface{}, error) {

	var appAssign ApplicationEntity

	var news NewsEntity
	var newsAssign NewsEntityResponse

	var user UserEntity
	var userAssign UserEntity

	var data = make([]NewsEntityResponse, 0) 

	allNews := []AllNewsEntity{}

	pageinteger, _  := strconv.Atoi(page) 
	limitinteger, _ := strconv.Atoi(limit)

	var offset = strconv.Itoa((pageinteger - 1) * limitinteger)

	allNewsQuery := `SELECT uid FROM news` 

	err := db.Debug().Raw(allNewsQuery).Scan(&allNews).Error

	if err != nil {
		helper.Logger("error", "In Server: "+err.Error())
	}

	var resultTotal = len(allNews)

	var perPage = math.Ceil(float64(resultTotal) / float64(limitinteger))

	var prevPage int
	var nextPage int

	if pageinteger == 1 {
		prevPage = 1
	} else {
		prevPage = pageinteger - 1 
	}

	if pageinteger == int(perPage) {
		nextPage = 1
	} else {
		nextPage = pageinteger + 1 
	}

	newsQuery := `SELECT n.uid, n.title, n.description, n.user_id, app.name AS application_name, app.uid AS application_id 
	FROM news n INNER JOIN applications app ON app.uid = n.application_id LIMIT 
	`+offset+`, `+limit+``

	rows, err := db.Debug().Raw(newsQuery).Rows()

	if err != nil {
		helper.Logger("error", "In Server: "+err.Error())
		return nil, errors.New(err.Error())
	}

	for rows.Next() {
		db.ScanRows(rows, &news)

		userQuery := `SELECT email, phone, fullname FROM users u 
		INNER JOIN user_profiles p ON p.user_id = u.uid 
		WHERE u.uid = '`+news.UserId+`'` 

		rows, err := db.Debug().Raw(userQuery).Rows()

		if err != nil {
			helper.Logger("error", "In Server: "+err.Error())
			return nil, errors.New(err.Error())
		}

		for rows.Next() {
			db.ScanRows(rows, &user)
			
			userAssign.Fullname = user.Fullname
			userAssign.Email = user.Email
			userAssign.Phone = user.Phone
		}

		appAssign.ApplicationId = news.ApplicationId
		appAssign.Name = news.ApplicationName

		newsAssign.Uid = news.Uid
		newsAssign.Title = news.Title
		newsAssign.Description = news.Description
		newsAssign.App = appAssign
		newsAssign.User = userAssign

		data = append(data, newsAssign)
	}

	return map[string]any {
		"total": resultTotal,
		"current_page": pageinteger,
		"per_page": int(perPage),
		"prev_page": prevPage,
		"next_page": nextPage,
		"news": &data,
	}, nil
}