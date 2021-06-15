package util

import (
	"auth-service/model"
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"net/http"
	"strings"
	"time"
)

const secret = "someSecret"

type TokenClaims struct {
	Username string `json:"userId"`
	Role model.Role `json:"role"`
	jwt.StandardClaims
}

func CreateJWT(username string, role *model.Role) (string, error) {
	tokenClaims := TokenClaims{Username: username, Role: *role, StandardClaims: jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Minute * 50).Unix(),
		IssuedAt: time.Now().Unix(),
	}}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims)
	tokenString, err := token.SignedString([]byte(secret))

	if err != nil {
		fmt.Errorf("Something went wrong: %s", err.Error())
		return "", err
	}

	return tokenString, nil
}

func GetJWT(r http.Header) string {
	bearToken := r.Get("Authorization")
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

func GetUsernameFromToken(r *http.Request) (string){
	tokenString := GetJWT(r.Header)
	token, err := jwt.ParseWithClaims(tokenString, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if claims, ok := token.Claims.(*TokenClaims); ok && token.Valid {
		return claims.Username
	} else {
		fmt.Println(err)
		return ""
	}
}
