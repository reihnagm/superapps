package services

import (
	"errors"
	"time"
	"golang.org/x/crypto/bcrypt"
	entities "superapps/entities"
	models "superapps/models"
	helper "superapps/helpers"
	middleware "superapps/middlewares"
	uuid "github.com/satori/go.uuid"
)

func VerifyOtp(u *models.User) (map[string]interface{}, error) {

	user := entities.User{}

	user.Email = u.Email
	user.Otp = u.Otp

	users := []entities.UserOtp{}
	query := `SELECT uid, email_active, otp_date FROM users 
	WHERE (email = '`+u.Email+`' OR phone = '`+u.Email+`') AND otp = '`+u.Otp+`'`

	err := db.Debug().Raw(query).Scan(&users).Error

	if err != nil {
		helper.Logger("error", "In Server: "+err.Error())
		return nil, errors.New(err.Error())
	}

	isUserExist := len(users)

	if isUserExist == 0 {
		return nil, errors.New("User not found")
	} 

	uid := users[0].Uid
	emailActive := users[0].EmailActive
	otpDate := users[0].OtpDate

	if emailActive == 1 {
		helper.Logger("error", "In Server: Account is already active")
		return nil, errors.New("Account is already active")
	}

	currentTime := time.Now()
    elapsed := currentTime.Sub(otpDate)

	if elapsed >= 1*time.Minute {
		helper.Logger("error", "In Server: Otp is expired")
		return nil, errors.New("Otp is expired")
    } 

	errUpdateEmailActive := db.Debug().Exec(`UPDATE users SET email_active = 1, email_active_date = NOW()
		WHERE email = '`+u.Email+`'
	`).Error

	if errUpdateEmailActive != nil {
		helper.Logger("error", "In Server: "+errUpdateEmailActive.Error())
		return nil, errors.New(errUpdateEmailActive.Error())
	}

	token, err := middleware.CreateToken(uid)
	if err != nil {
		helper.Logger("error", "In Server: "+err.Error())
		return nil, err
	}

	access := token["token"]

	return map[string]interface{}{"token": access}, nil
}

func ResendOtp(u *models.User) (map[string]interface{}, error) {

	users := []entities.UserOtp{}
	query := `SELECT email_active, otp_date FROM users 
	WHERE (email = '`+u.Email+`' OR phone = '`+u.Email+`')`

	err := db.Debug().Raw(query).Scan(&users).Error

	if err != nil {
		helper.Logger("error", "In Server: "+err.Error())
		return nil, errors.New(err.Error())
	}

	isUserExist := len(users)

	if isUserExist == 0 {
		return nil, errors.New("User not found")
	} 

	emailActive := users[0].EmailActive
	otpDate := users[0].OtpDate

	if emailActive == 1 {
		helper.Logger("error", "In Server: Account is already active")
		return nil, errors.New("Account is already active")
	}

	currentTime := time.Now()
    elapsed := currentTime.Sub(otpDate)

	otp := helper.CodeOtp()

	if elapsed >= 1*time.Minute {
		errUpdateResendOtp:= db.Debug().Exec(`UPDATE users SET otp = '`+otp+`', otp_date = NOW() WHERE email = '`+u.Email+`'`).Error

		if errUpdateResendOtp != nil {
			helper.Logger("error", "In Server: "+errUpdateResendOtp.Error())
			return nil, errors.New(errUpdateResendOtp.Error())
		}
    } 

	return map[string]interface{}{
		"otp": otp,
	}, nil
}

func Login(u *models.User) (map[string]interface{}, error) {

	user := entities.User{}

	user.Password = u.Password

	users := []entities.UserLogin{}
	query := `SELECT uid, email_active, password FROM users WHERE email = '`+u.Val+`' OR phone = '`+u.Val+`'`

	err := db.Debug().Raw(query).Scan(&users).Error

	if err != nil {
		helper.Logger("error", "In Server: "+err.Error())
		return nil, errors.New(err.Error())
	}

	isUserExist := len(users)

	if isUserExist == 0 {
		return nil, errors.New("User not found")
	} 
	
	emailActive := users[0].EmailActive
	user.Id = users[0].Uid

	if emailActive == 0 { 
		err := db.Debug().Exec(`UPDATE users SET otp = '`+helper.CodeOtp()+`', otp_date = NOW() 
		WHERE email = '`+u.Val+`' OR phone = '`+u.Val+`'`).Error
		
		if err != nil {
			helper.Logger("error", "In Server: "+err.Error())
			return nil, errors.New(err.Error())
		}

		helper.Logger("error", "In Server: Please activate your account")
		return nil, errors.New("Please activate your account")
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

func Register(u *models.User) (map[string]interface{}, error) {

	hashedPassword, err := helper.Hash(u.Password)
	if err != nil {
		helper.Logger("error", "In Server: "+err.Error())
		return nil, err
	}

	user := entities.User{}

	user.Id = uuid.NewV4().String()

	profileUuid := uuid.NewV4().String()

	user.Email = u.Email
	user.Phone = u.Phone
	user.Password = string(hashedPassword)

	otp := helper.CodeOtp()

	users := []entities.CheckAccount{}
	errCheckAccount := db.Debug().Raw(`SELECT email FROM users WHERE email = '`+u.Email+`'`).Scan(&users).Error

	if errCheckAccount != nil {
		helper.Logger("error", "In Server: "+errCheckAccount.Error())
		return nil, errors.New(errCheckAccount.Error())
	}

	isUserExist := len(users)

	if isUserExist == 1 {
		return nil, errors.New("User already exist")
	} 
	
	applications := []entities.Application{}
	errCheckApp := db.Debug().Raw(`SELECT uid, username FROM applications WHERE username = '`+u.AppName+`'`).Scan(&applications).Error
	
	if errCheckApp != nil {
		helper.Logger("error", "In Server: "+errCheckApp.Error())
		return nil, errors.New(errCheckApp.Error())
	}

	isAppExist := len(applications)

	if isAppExist == 0 {
		return nil, errors.New("App not found")
	} 

	ApplicationId := applications[0].Uid

	errInsertUser := db.Debug().Exec(`INSERT INTO users (uid, email, phone, password, otp, application_id) 
	VALUES ('`+user.Id+`', '`+user.Email+`', '`+user.Phone+`', '`+user.Password+`', '`+otp+`', '`+ApplicationId+`')`).Error

	if errInsertUser != nil {
		helper.Logger("error", "In Server: "+errInsertUser.Error())
		return nil, errors.New(errInsertUser.Error())
	}

	errUserProfile := db.Debug().Exec(`INSERT INTO user_profiles (uid, user_id) VALUES('`+profileUuid+`', '`+user.Id+`')`).Error

	if errUserProfile != nil {
		helper.Logger("error", "In Server: "+errUserProfile.Error())
		return nil, errors.New(errUserProfile.Error())
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