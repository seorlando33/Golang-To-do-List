package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	connection "toDoList/api/db"
	auxiliarFunc "toDoList/api/lib"
	models "toDoList/api/models"
	"toDoList/api/services"

	"github.com/joho/godotenv"
)

func Register(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		auxiliarFunc.BadGateway(w)
		return
	}

	db, err := connection.DBConnection()
	if err != nil {
		auxiliarFunc.InternalServerError(w, err)
		return
	}

	user := &models.User{}

	err = json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		auxiliarFunc.BadRequest(w, err)
		return
	}

	err = services.RegisterService(db, *user)
	
	if err != nil {
		auxiliarFunc.BadRequest(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func Login(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		auxiliarFunc.BadGateway(w)
		return
	}

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := connection.DBConnection()
	if err != nil {
		auxiliarFunc.InternalServerError(w, err)
		return
	}

	user1 := &models.User{}

	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	err = json.NewDecoder(r.Body).Decode(user1)

	if err != nil {
		auxiliarFunc.BadRequest(w, err)
		return
	}

	signedToken, expirationTime, err := services.LoginService(db, *user1, os.Getenv("jwtKey"))

	if err != nil {
		auxiliarFunc.InternalServerError(w, err)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "jwt",
		Value:   signedToken,
		Expires: expirationTime,
	})
}
