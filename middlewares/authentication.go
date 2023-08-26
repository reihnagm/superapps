package middlewares

import (
	helper "superapps/helpers"
	"context"
	"net/http"
	"os"
	"strings"
	"time"
	"github.com/dgrijalva/jwt-go"
)

func JwtAuthentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		tokenHeader := r.Header.Get("Authorization")

		if r.URL.Path == "/api/v1/login" || r.URL.Path == "/api/v1/register" {
			next.ServeHTTP(w, r)
			return
		}

		if tokenHeader == "" {
			helper.Logger("error", "In Server: Missing auth token")
			resp := map[string]interface{}{}
			helper.Response(w, http.StatusUnauthorized, true, "Missing auth token", resp)
			return
		}

		splitted := strings.Split(tokenHeader, " ")

		if len(splitted) != 2 {
			helper.Logger("error", "In Server: Missing auth token")
			resp := map[string]interface{}{}
			helper.Response(w, http.StatusUnauthorized, true, "Missing auth token", resp)
			return
		}

		tokenPart := splitted[1]

		claims := jwt.MapClaims{}

		token, err := jwt.ParseWithClaims(tokenPart, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil {
			helper.Logger("error", "In Server: "+err.Error())
			resp := map[string]interface{}{}
			helper.Response(w, http.StatusUnauthorized, true, "Token expired", resp)
			return
		}

		if !token.Valid {
			helper.Logger("error", "In Server: Token Expired")
			resp := map[string]interface{}{}
			helper.Response(w, http.StatusUnauthorized, true, "Token expired", resp)
			return
		}

		ctx := context.WithValue(r.Context(), "user", claims)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

func CreateToken(userId string, fullname string, username string, publicId string) (map[string]string, error) {

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["uid"] = userId
	claims["fullname"] = fullname
	claims["username"] = username
	claims["publicId"] = publicId
	claims["exp"] = time.Now().Add(time.Hour * 168).Unix()

	access, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		helper.Logger("error", "In Server: "+err.Error())
		return nil, err
	}

	refToken := jwt.New(jwt.SigningMethodHS256)
	refClaims := refToken.Claims.(jwt.MapClaims)
	refClaims["authorized"] = true
	refClaims["uid"] = userId
	refClaims["fullname"] = fullname
	refClaims["username"] = username
	refClaims["publicId"] = publicId
	refClaims["exp"] = time.Now().Add(time.Hour * 192).Unix()

	refresh, err := refToken.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		helper.Logger("error", "In Server: "+err.Error())
		return nil, err
	}

	return map[string]string{"access_token": access, "refresh_token": refresh}, nil
}
