package middlewares

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"

	models "toDoList/api/models"
)

func JWTValidator(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file, en middleware")
		}

		cookie, err := r.Cookie("jwt")
		if err != nil {
			if err == http.ErrNoCookie {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		token := cookie.Value
		jwtStructure := &models.JWTStructure{}

		parsedToken, err := jwt.ParseWithClaims(token, jwtStructure, func(tkn *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("jwtKey")), nil
		})
		
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprintf(w, "%v", err)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			fmt.Println("entre acá", err)

			return
		}

		if !parsedToken.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w, "%v", err)
			return
		}

		next(w, r)
	})
}