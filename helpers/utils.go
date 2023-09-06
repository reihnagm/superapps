package helper

// import (	
// 	"net/mail"
// )

import (
    "regexp"
)

func IsValidEmail(email string) bool {
    // Optional 1
    // _, err := mail.ParseAddress(email)
    // return err == nil

    emailRegex := regexp.MustCompile(`^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|.(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$`)
    return emailRegex.MatchString(email)
}