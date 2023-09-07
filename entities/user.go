package entities

type User struct {
	Id          string `json:"id"`
	Val  		string `json:"val"`
	Email       string `json:"email"`
	Phone	    string `json:"phone"`
	Password 	string `json:"password"`
}