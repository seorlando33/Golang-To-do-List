package test

import (

	"database/sql"

	models "toDoList/api/models"
)


var User1 models.User = models.User{
	ID:        1,
	Full_name: "user 1",
	Password:  "1234",
	Email:     "user1@gmail.com",
	Picture:   "html://fakeurl.com/1",
}
var User2 models.User = models.User{
	ID:        2,
	Full_name: "user 2",
	Password:  "1234",
	Email:     "user2@gmail.com",
	Picture:   "html://fakeurl.com/2",
}
var User3 models.User = models.User{
	ID: 3,
	Full_name: "user 3",
	Password: "1234",
	Email: "user3@gmail.com",
	Picture: "html://fakeurl.com/3",
}
var Task1 models.Task = models.Task{
	ID:          1,
	User_ID:     1,
	Description: "Do something",
	Fulfilled:   false,
}
var Task2 models.Task = models.Task{
	ID:          2,
	User_ID:     1,
	Description: "Do another thing",
	Fulfilled:   false,
}
var Task3 models.Task = models.Task{
	ID:          3,
	User_ID:     1,
	Description: "Do more thing",
	Fulfilled:   false,
}
var Task4 models.Task = models.Task{
	ID:          4,
	User_ID:     2,
	Description: "Do another thing",
	Fulfilled:   false,
}
var Task5 models.Task = models.Task{
	ID:          5,
	User_ID:     2,
	Description: "Do something different",
	Fulfilled:   false,
}

var DescriptionTask1 models.Task = models.Task{
	User_ID:     1,
	Description: " thing",
}

var DescriptionTask2 models.Task = models.Task{
	User_ID:     1,
	Description: "Take the pencil",
}

var UpdateTask models.Task = models.Task{
	ID:        4,
	User_ID:   2,
	Fulfilled: true,
}

var StateTasks models.Task = models.Task{
	User_ID:   1,
	Fulfilled: true,
}

func GetConnection() (*sql.DB, error) {

	connStr := "postgresql://@127.0.0.1:5432/test?sslmode=disable"
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		return nil, err
	}

	return db, nil
}