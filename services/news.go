package services

import (
	"os"
	"strconv"
	"math"
	"errors"
	entities "superapps/entities"
	helper "superapps/helpers"
)

func GetNews(search, page, limit string) (map[string]interface{}, error) {

	url := os.Getenv("API_URL")

	var appAssign entities.ApplicationResponseEntity

	var news entities.NewsFormEntity
	var newsAssign entities.NewsResponseEntity

	var user entities.UserEntity
	var userAssign entities.UserResponseEntity

	var data = make([]entities.NewsResponseEntity, 0) 

	allNews := []entities.AllCountNewsEntity{}

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
	
	nextPage = pageinteger + 1

	newsQuery := `SELECT n.uid, n.title, n.description, n.user_id, n.created_at, app.name AS application_name, app.uid AS application_id 
	FROM news n 
	INNER JOIN applications app ON app.uid = n.application_id 
	WHERE n.title LIKE '%`+search+`%'
	LIMIT `+offset+`, `+limit+``

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
		
		var createdAt = news.CreatedAt.Format("2006-01-02 15:04")

		appAssign.ApplicationId = news.ApplicationId
		appAssign.Name = news.ApplicationName

		newsAssign.Uid = news.Uid
		newsAssign.Title = news.Title
		newsAssign.Description = news.Description
		newsAssign.App = appAssign
		newsAssign.User = userAssign
		newsAssign.CreatedAt = createdAt

		data = append(data, newsAssign)
	}

	var nextUrl = strconv.Itoa(nextPage)
	var prevUrl = strconv.Itoa(prevPage)

	return map[string]any {
		"total": resultTotal,
		"current_page": pageinteger,
		"per_page": int(perPage),
		"prev_page": prevPage,
		"next_page": nextPage,
		"next_url": url + "?page=" + nextUrl,
		"prev_url": url + "?page=" + prevUrl,
		"news": &data,
	}, nil
}