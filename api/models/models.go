package api

import (
	"time"
	"github.com/dgrijalva/jwt-go"
)

type User struct {
	ID				int			`json:"id,omitempty"`
	Full_name		string		`json:"full_name,omitempty"`
	Password		string		`json:"password,omitempty"`
	Email			string		`json:"email,omitempty"`
	Picture			string		`json:"picture,omitempty"`
	Created_At		time.Time	`json:"created_at,omitempty"`
	Updated_At		time.Time	`json:"updated_at,omitempty"`
}

type Task struct {
	ID				int			`json:"id,omitempty"`
	User_ID			int			`json:"user_id,omitempty"`
	Description		string		`json:"description,omitempty"`
	Fulfilled		bool		`json:"fulfilled,omitempty"`
	Created_At		time.Time	`json:"created_at,omitempty"`
	Updated_At		time.Time	`json:"updated_at,omitempty"`
}

type JWTStructure struct {
	ID					string	`json:"id"`
	jwt.StandardClaims			
}