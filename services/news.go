package services

import (
	"os"
	"strconv"
	"math"
	"errors"
	"fmt"
	models "superapps/models"
	entities "superapps/entities"
	helper "superapps/helpers"
	// uuid "github.com/satori/go.uuid"
)

func GetNews(search, page, limit, appName string) (map[string]interface{}, error) {

	url := os.Getenv("API_URL")

	var appAssign entities.NewsApplicationResponse

	var newsImage entities.NewsImageForm
	var newsImageAssign entities.NewsImageResponse

	var news entities.News

	var user entities.NewsUser
	var userAssign entities.NewsUserResponse

	var newsAssign entities.NewsResponse
	var appendNewsAssign = make([]entities.NewsResponse, 0) 

	var allCountNews []models.AllCountNews

	pageinteger, _  := strconv.Atoi(page) 
	limitinteger, _ := strconv.Atoi(limit)

	var offset = strconv.Itoa((pageinteger - 1) * limitinteger)

	errAllNewsQuery := db.Debug().Raw(`SELECT uid FROM news`).Scan(&allCountNews).Error

	if errAllNewsQuery != nil {
		helper.Logger("error", "In Server: "+errAllNewsQuery.Error())
	}

	var resultTotal = len(allCountNews)

	var perPage = math.Ceil(float64(resultTotal) / float64(limitinteger))

	var prevPage int
	var nextPage int

	if pageinteger == 1 {
		prevPage = 1
	} else {
		prevPage = pageinteger - 1 
	}
	
	nextPage = pageinteger + 1

	rows, errNewsQuery := db.Debug().Raw(`SELECT n.uid, n.title, n.description, n.user_id, n.created_at, app.name AS application_name, app.uid AS app_id 
	FROM news n 
	INNER JOIN applications app ON app.uid = n.app_id 
	WHERE n.title LIKE '%`+search+`%' AND app.username LIKE '%`+appName+`%'
	LIMIT `+offset+`, `+limit+``).Rows()

	if errNewsQuery != nil {
		helper.Logger("error", "In Server: "+errNewsQuery.Error())
		return nil, errors.New(errNewsQuery.Error())
	}

	for rows.Next() {
		db.ScanRows(rows, &news)

		rows, errUserQuery := db.Debug().Raw(`SELECT email, phone, fullname FROM users u 
		INNER JOIN user_profiles p ON p.user_id = u.uid 
		WHERE u.uid = '`+news.UserId+`'`).Rows()

		if errUserQuery != nil {
			helper.Logger("error", "In Server: "+errUserQuery.Error())
			return nil, errors.New(errUserQuery.Error())
		}

		for rows.Next() {
			db.ScanRows(rows, &user)
			
			userAssign.Fullname = user.Fullname
			userAssign.Email = user.Email
			userAssign.Phone = user.Phone
		}

		var dataNewsImage = make([]entities.NewsImageResponse, 0) 

		rows, errNewsImageQuery := db.Debug().Raw(`SELECT path, size FROM news_images WHERE news_id = '`+news.Uid+`'` ).Rows()

		if errNewsImageQuery != nil {
			helper.Logger("error", "In Server: "+errNewsImageQuery.Error())
			return nil, errors.New(errNewsImageQuery.Error())
		}

		for rows.Next() {
			db.ScanRows(rows, &newsImage)
			
			if newsImage.Path != "" {
				newsImageAssign.Path = newsImage.Path
				newsImageAssign.Size = newsImage.Size

				dataNewsImage = append(dataNewsImage, newsImageAssign)
			}
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

		newsAssign.Images = dataNewsImage

		newsAssign.App = appAssign
		newsAssign.User = userAssign
		newsAssign.CreatedAt = createdAt

		appendNewsAssign = append(appendNewsAssign, newsAssign)
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
		"data": &appendNewsAssign,
	}, nil
}

func CreateImageNews(n *models.NewsImageForm) (map[string]interface{}, error) {
	Uid  := n.NewsId 
	Path := n.Path

	fmt.Println(Uid)
	fmt.Println(Path)

	return map[string]interface{}{}, nil
}

func CreateNews(n *models.News) (map[string]interface{}, error) {

	applications := []entities.Application{}
	errCheckApp := db.Debug().Raw(`SELECT uid, username FROM applications WHERE username = '`+n.ApplicationName+`'`).Scan(&applications).Error
	
	if errCheckApp != nil {
		helper.Logger("error", "In Server: "+errCheckApp.Error())
		return nil, errors.New(errCheckApp.Error())
	}

	isAppExist := len(applications)

	if isAppExist == 0 {
		return nil, errors.New("App not found")
	} 

	ApplicationId := applications[0].Uid

	// n.Uid = uuid.NewV4().String()

	errInsertNews := db.Debug().Exec(`INSERT INTO news (uid, title, description, app_id, user_id) 
	VALUES ('`+n.Uid+`', '`+n.Title+`', '`+n.Description+`', '`+ApplicationId+`', '`+n.UserId+`')`).Error

	if errInsertNews != nil {
		helper.Logger("error", "In Server: "+errInsertNews.Error())
		return nil, errors.New(errInsertNews.Error())
	}	

	return map[string]interface{}{}, nil
}