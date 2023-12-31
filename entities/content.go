package entities

import (
	"time"
)

// CONTENT -------

type Content struct {
	Uid				string `json:"id"`
	Title 			string `json:"title"`
	Description 	string `json:"desc"`
	UserId 			string `json:"user_id"`
	CreatedAt		time.Time `json:"created_at"`
	AppId 			string `json:"app_id"`
	AppName 		string `json:"app_name"`
	Types			string `json:"types"`
	Application 	ContentApplicationResponse `json:"app"`
	User 			ContentUserResponse `json:"user"`
}


type ContentResponse struct {
	Uid			string `json:"id"`
	Title 		string `json:"title"`
	Description	string `json:"desc"`
	File		[]ContentMediaResponse `json:"files"`
	Like		[]ContentLikeResponse `json:"likes"`
	Unlike		[]ContentUnlikeResponse `json:"unlikes"`
	Comment 	[]ContentCommentResponse `json:"comments"`
	App 		ContentApplicationResponse `json:"app"`
	User 		ContentUserResponse `json:"user"`
	Types		string `json:"types"`
	CreatedAt   string `json:"created_at"`
}

// --------------

// Types --------

type Types struct {
	types 	string `json:"types"`
}

// --------------

type ReqContentLike struct {
	ContentId	string `json:"content_id"`
	UserId 		string 	`json:"user_id"`
}

type ReqContentUnlike struct {
	ContentId	string `json:"content_id"`
	UserId 		string 	`json:"user_id"`
}

type ContentLike struct {
	Uid		 	string `json:"id"`
	UserId 		string `json:"fullname"`
	Fullname 	string `json:"fullname"`
}

type ContentUnlike struct {
	Uid		 	string `json:"id"`
	UserId 		string `json:"fullname"`
	Fullname 	string `json:"fullname"`
}

type ContentLikeUser struct {
	UserId		string `json:"id"`
	Fullname	string `json:"name"`
}

type ContentUnlikeUser struct {
	UserId		string `json:"id"`
	Fullname	string `json:"name"`
}

type ContentLikeResponse struct {
	Uid		 	string `json:"id"`
	User	 	ContentLikeUserResponse `json:"user"`
}

type ContentUnlikeResponse struct {
	Uid		 	string `json:"id"`
	User	 	ContentUnlikeUserResponse `json:"user"`
}

type ContentLikeUserResponse struct {
	UserId	 	string `json:"id"`
	Fullname    string `json:"name"`
}

type ContentUnlikeUserResponse struct {
	UserId	 	string `json:"id"`
	Fullname    string `json:"name"`
}

type ReqContentComment struct {
	ContentId   string `json:"content_id"`
	UserId		string `json:"user_id"`
	Comment 	string `json:"comment"`
}

type DelContentComment struct {
	Uid 		string `json:"id"`
}

type ContentComment struct {
	Uid 		string `json:"id"`
	UserId		string `json:"user_id"`
	Fullname 	string `json:"fullname"`
	Comment	    string `json:"comment"`
}

type ContentCommentUser struct {
	UserId		string `json:"user_id"`
	Fullname    string `json:"name"`
}

type ContentCommentResponse struct {
	Uid 		string `json:"id"`
	Comment	    string `json:"comment"`
	User	 	ContentCommentUserResponse `json:"user"`
}

type ContentCommentUserResponse struct {
	UserId		string `json:"user_id"`
	Fullname    string `json:"name"`
}

type ContentApplicationResponse struct {
	ApplicationId 	string `json:"id"`
	ApplicationName string `json:"name"`
}

type ContentMedia struct {
	ContentId	string `json:"content_id"`
	Path 		string `json:"path"`
	Size		int    `json:"size"`
}

type AllCountContent struct {
	Uid			  string `json:"id"`
}

type ContentMediaResponse struct {
	ContentId	string `json:"content_id"`
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