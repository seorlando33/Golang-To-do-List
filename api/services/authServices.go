package services

import (
	"database/sql"
	"errors"
	"time"

	auxiliarFunc "toDoList/api/lib"
	models "toDoList/api/models"

	"golang.org/x/crypto/bcrypt"

)

func RegisterService(db *sql.DB, user models.User) error {

	if user.Full_name == "" || user.Email == "" {
		return errors.New("critical data required")
	}

	query := `INSERT INTO users(full_name, password, email, picture) values ($1,$2,$3,$4);`

	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	
	if err != nil {
		return err
	}
	_, err = db.Query(query, user.Full_name, password, user.Email, user.Picture)

	if err != nil {
		return err
	}
	return nil
}

func LoginService(db *sql.DB, user1 models.User, secret string) (string, time.Time, error) {

	var fakeTime time.Time
	var user2 models.User

	query := `SELECT * FROM users WHERE email = $1;`
	row := db.QueryRow(query, user1.Email)

	err := row.Scan(&user2.ID, &user2.Full_name, &user2.Password, &user2.Email, &user2.Picture, &user2.Created_At, &user2.Updated_At)

	if err != nil {
		return "", fakeTime, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user2.Password), []byte(user1.Password))

	if err != nil {
		return "", fakeTime, err
	}

	signedToken, expirationTime, _ := auxiliarFunc.JWTGenerator(user2, secret)

	return signedToken, expirationTime, nil
}
