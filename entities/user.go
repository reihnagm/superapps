package entities

import (
	"time"
)

type User struct {
	Id       string `json:"id"`
	Val  	 string `json:"val"`
	Email    string `json:"email"`
	Phone	 string `json:"phone"`
	Otp	     string `json:"otp"`
	Password string `json:"password"`
}

type CheckAccount struct {
	Email string `json:"email"`
}

type UserLogin struct {
	EmailActive int `json:"email_active"`
	Password 	string `json:"password"`
}

type UserOtp struct {
	Uid 		string 	  `json:"uid"`
	EmailActive int	   	  `json:"email_active"`
	OtpDate 	time.Time `json:"otp_date"`
}