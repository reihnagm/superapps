package services

import (
	"os"
	"strconv"
	"math"
	"errors"
	// "fmt"
	models "superapps/models"
	entities "superapps/entities"
	helper "superapps/helpers"
	// uuid "github.com/satori/go.uuid"
)

func GetContent(search, page, limit, appName string) (map[string]interface{}, error) {

	url := os.Getenv("API_URL")

	var appAssign entities.ContentApplicationResponse

	var contentMedia entities.ContentMedia
	var contentMediaAssign entities.ContentMediaResponse

	var content entities.Content

	var contentLike entities.ContentLike
	var contentLikeAssign entities.ContentLikeResponse
	var contentLikeUserAssign entities.ContentLikeUserResponse

	var contentComment entities.ContentComment
	var contentCommentAssign entities.ContentCommentResponse
	var contentCommentUserAssign entities.ContentCommentUserResponse

	var user entities.ContentUser
	var userAssign entities.ContentUserResponse

	var contentAssign entities.ContentResponse
	var appendContentAssign = make([]entities.ContentResponse, 0) 

	var allCountContent []models.AllCountContent

	pageinteger, _  := strconv.Atoi(page) 
	limitinteger, _ := strconv.Atoi(limit)

	var offset = strconv.Itoa((pageinteger - 1) * limitinteger)

	errAllContentQuery := db.Debug().Raw(`SELECT uid FROM contents`).Scan(&allCountContent).Error

	if errAllContentQuery != nil {
		helper.Logger("error", "In Server: "+errAllContentQuery.Error())
	}

	var resultTotal = len(allCountContent)

	var perPage = math.Ceil(float64(resultTotal) / float64(limitinteger))

	var prevPage int
	var nextPage int

	if pageinteger == 1 {
		prevPage = 1
	} else {
		prevPage = pageinteger - 1 
	}
	
	nextPage = pageinteger + 1

	rows, errContentQuery := db.Debug().Raw(`SELECT n.uid, n.title, n.description, n.user_id, n.created_at, 
	app.name AS app_name, app.uid AS app_id 
	FROM contents n 
	INNER JOIN applications app ON app.uid = n.app_id 
	WHERE n.title LIKE '%`+search+`%' AND app.username LIKE '%`+appName+`%'
	LIMIT `+offset+`, `+limit+``).Rows()

	if errContentQuery != nil {
		helper.Logger("error", "In Server: "+errContentQuery.Error())
		return nil, errors.New(errContentQuery.Error())
	}

	for rows.Next() {
		db.ScanRows(rows, &content)

		rows, errUserQuery := db.Debug().Raw(`SELECT email, phone, fullname FROM users u 
		INNER JOIN user_profiles p ON p.user_id = u.uid 
		WHERE u.uid = '`+content.UserId+`'`).Rows()

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

		rows.Close()

		var dataContentMedia = make([]entities.ContentMediaResponse, 0) 

		rows, errContentMediaQuery := db.Debug().Raw(`SELECT content_id, path, size FROM content_medias WHERE content_id = '`+content.Uid+`'`).Rows()

		if errContentMediaQuery != nil {
			helper.Logger("error", "In Server: "+errContentMediaQuery.Error())
			return nil, errors.New(errContentMediaQuery.Error())
		}

		for rows.Next() {
			db.ScanRows(rows, &contentMedia)
			
			if contentMedia.Path != "" {
				contentMediaAssign.ContentId = contentMedia.ContentId
				contentMediaAssign.Path = contentMedia.Path
				contentMediaAssign.Size = contentMedia.Size

				dataContentMedia = append(dataContentMedia, contentMediaAssign)
			}
		}

		rows.Close()

		var dataContentLike = make([]entities.ContentLikeResponse, 0)

		rows, errContentLikeQuery := db.Debug().Raw(`SELECT cl.uid, p.user_id, p.fullname FROM content_likes cl 
		INNER JOIN user_profiles p ON p.user_id = cl.user_id
		WHERE content_id = '`+content.Uid+`'`).Rows()

		if errContentLikeQuery != nil {
			helper.Logger("error", "In Server: "+errContentLikeQuery.Error())
			return nil, errors.New(errContentLikeQuery.Error())
		}

		for rows.Next() {
			db.ScanRows(rows, &contentLike)

			
			if contentLike.Uid != "" {

				contentLikeUserAssign.UserId = contentLike.UserId
				contentLikeUserAssign.Fullname = contentLike.Fullname

				contentLikeAssign.Uid = contentLike.Uid
				contentLikeAssign.User = contentLikeUserAssign

				dataContentLike = append(dataContentLike, contentLikeAssign)
			}
		}
		
		rows.Close()

		var dataContentComment = make([]entities.ContentCommentResponse, 0)

		rows, errContentCommentQuery := db.Debug().Raw(`SELECT cc.uid, cc.comment, cc.user_id, p.fullname FROM content_comments cc
		INNER JOIN user_profiles p ON p.user_id = cc.user_id
		WHERE content_id = '`+content.Uid+`'`).Rows()

		if errContentCommentQuery != nil {
			helper.Logger("error", "In Server: "+errContentCommentQuery.Error())
			return nil, errors.New(errContentCommentQuery.Error())
		}

		for rows.Next() {
			db.ScanRows(rows, &contentComment)
			
			if contentComment.Uid != "" {

				contentCommentUserAssign.UserId = contentComment.UserId
				contentCommentUserAssign.Fullname = contentComment.Fullname
				
				contentCommentAssign.Uid = contentComment.Uid
				contentCommentAssign.Comment = contentComment.Comment

				contentCommentAssign.User = contentCommentUserAssign

				dataContentComment = append(dataContentComment, contentCommentAssign)
			}
		}
		
		rows.Close()
		
		var createdAt = content.CreatedAt.Format("2006-01-02 15:04")

		appAssign.ApplicationId = content.AppId
		appAssign.ApplicationName = content.AppName

		contentAssign.Uid = content.Uid
		contentAssign.Title = content.Title
		contentAssign.Description = content.Description

		contentAssign.File = dataContentMedia
		contentAssign.Like = dataContentLike
		contentAssign.Comment = dataContentComment

		contentAssign.App = appAssign
		contentAssign.User = userAssign
		contentAssign.CreatedAt = createdAt

		appendContentAssign = append(appendContentAssign, contentAssign)
	}

	rows.Close()

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
		"data": &appendContentAssign,
	}, nil
}

func CreateMediaContent(n *models.ContentMedia) (map[string]interface{}, error) {
	// Uid  := n.ContentId 
	// Path := n.Path

	return map[string]interface{}{}, nil
}

func CreateContent(n *models.Content) (map[string]interface{}, error) {

	applications := []entities.Application{}
	errCheckApp := db.Debug().Raw(`SELECT uid, username FROM applications WHERE username = '`+n.AppName+`'`).Scan(&applications).Error
	
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

	errInsertContent := db.Debug().Exec(`INSERT INTO contents (uid, title, description, app_id, user_id) 
	VALUES ('`+n.Uid+`', '`+n.Title+`', '`+n.Description+`', '`+ApplicationId+`', '`+n.UserId+`')`).Error

	if errInsertContent != nil {
		helper.Logger("error", "In Server: "+errInsertContent.Error())
		return nil, errors.New(errInsertContent.Error())
	}	

	return map[string]interface{}{}, nil
}