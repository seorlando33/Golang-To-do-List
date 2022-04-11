package lib

import (
	"time"
	"strconv"

	models "toDoList/api/models"

	"github.com/dgrijalva/jwt-go"
)

func JWTGenerator(user models.User, secret string) (string, time.Time, error) {

	expirationTime := time.Now().Add(336 * time.Hour) // 14 days

	JWT := &models.JWTStructure{
		ID: strconv.Itoa(user.ID),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, JWT)

	JWTKey := []byte(secret)

	signedToken, err := token.SignedString(JWTKey)
	if err != nil {
		return "", time.Now(), err
	}

	return signedToken, expirationTime, nil
}
