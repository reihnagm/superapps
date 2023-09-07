package models

import (
	"errors"
	// "fmt"
	"golang.org/x/crypto/bcrypt"
	// "github.com/jinzhu/gorm"
	// _ "github.com/jinzhu/gorm/dialects/mysql"
	helper "superapps/helpers"
	middleware "superapps/middlewares"
	uuid "github.com/satori/go.uuid"
)

type UserLogin struct {
	Password string
}

func (u *User) Login() (map[string]interface{}, error) {

	user := User{}

	uuid := uuid.NewV4()
	user.Id = uuid.String()
	user.Password = u.Password

	users := []UserLogin{}
	query := `SELECT password FROM users WHERE email = '`+u.Val+`' OR phone = '`+u.Val+`'`

	err := db.Debug().Raw(query).Scan(&users).Error

	if err != nil {
		helper.Logger("error", "In Server: "+err.Error())
		return nil, errors.New(err.Error())
	}

	isUserExist := len(users)

	if isUserExist == 0 {
		return nil, errors.New("Credentials is incorrect")
	} 

	passHashed := users[0].Password
	
	err = helper.VerifyPassword(passHashed, u.Password)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		helper.Logger("error", "In Server: "+err.Error())
		return nil, errors.New("Credentials is incorrect")
	}

	token, err := middleware.CreateToken(user.Id)
	if err != nil {
		helper.Logger("error", "In Server: "+err.Error())
		return nil, err
	}

	access := token["token"]

	return map[string]interface{}{"token": access}, nil
}

func (u *User) Register() (map[string]interface{}, error) {

	hashedPassword, err := helper.Hash(u.Password)
	if err != nil {
		helper.Logger("error", "In Server: "+err.Error())
		return nil, err
	}

	user := User{}

	uuid := uuid.NewV4()
	user.Id = uuid.String()

	user.Email = u.Email
	user.Phone = u.Phone
	user.Password = string(hashedPassword)

	err = db.Debug().Exec(`INSERT INTO users (uid, email, phone, password) 
	VALUES ('`+user.Id+`', '`+user.Email+`', '`+user.Phone+`', '`+user.Password+`')`).Error

	if err != nil {
		helper.Logger("error", "In Server: "+err.Error())
		return nil, errors.New(err.Error())
	}

	token, err := middleware.CreateToken(user.Id)
	if err != nil {
		helper.Logger("error", "In Server: "+err.Error())
	}

	// resp := map[string]interface{}{}
	// resp["id"] = user.Id

	access := token["token"]

	return map[string]interface{}{"token": access}, nil
}