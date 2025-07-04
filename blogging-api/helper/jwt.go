package helper

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var accesskey = os.Getenv("ACCESS_SECRET_KEY")
var refreshkey = os.Getenv("REFRESH_SECRET_KEY")

func GenerateToken(username string, id int64) (string, string, error) {
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"userID":   id,
		"exp":      time.Now().Add(time.Minute * 15).Unix(),
	})
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"userID":   id,
		"exp":      time.Now().Add(time.Hour * 2).Unix(),
	})

	access_T, err := accessToken.SignedString([]byte(accesskey))
	if err != nil {
		return "", "", err
	}
	refresh_T, err := refreshToken.SignedString([]byte(refreshkey))
	if err != nil {
		return "", "", err
	}

	return access_T, refresh_T, nil
}

func VerifyToken(token string) (int64, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(accesskey), nil
	})

	if err != nil {
		return 0, errors.New("could not parse the token")
	}

	validToken := parsedToken.Valid
	if !validToken {
		return 0, errors.New("invalid token")
	}
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("invalid token claims")
	}
	// username := claims["username"].(string)
	userid := int64(claims["userID"].(float64))

	return userid, nil
}
