package test

import (
	"database/sql"
	"strings"
	"testing"

	_ "github.com/lib/pq"

	auxiliarFunc "toDoList/api/lib"
	models "toDoList/api/models"
	services "toDoList/api/services"
)

var DB *sql.DB

func TestRegisterService(t *testing.T) {

	db, err := GetConnection()

	if err != nil {
		t.FailNow()
	}
	DB = db

	_, _ = db.Query(InicialTableUsersQuery)
	_, _ = db.Query(InicialTableTasksQuery)

	_ = services.RegisterService(db, User1)
	_ = services.RegisterService(db, User2)

	rows, _ := db.Query(GetAllUsersQuery)

	var users []models.User

	for rows.Next() {
		var user models.User

		_ = rows.Scan(&user.ID, &user.Full_name, &user.Password, &user.Email, &user.Picture, &user.Created_At, &user.Updated_At)

		users = append(users, user)
	}

	if len(users) != 2 {
		t.Fail()
	}
}

func TestLoginService(t *testing.T) {
	var secret string = "Super-Difficult-Password"
	signedToken, _, _ := services.LoginService(DB, User1, secret)

	splitedToken := strings.Split(signedToken, ".")

	if len(splitedToken) != 3 {
		t.Fail()
	}
}

func TestCreateTaskService(t *testing.T) {

	_ = services.CreateTaskService(DB, Task1)
	_ = services.CreateTaskService(DB, Task2)
	_ = services.CreateTaskService(DB, Task3)
	_ = services.CreateTaskService(DB, Task4)
	_ = services.CreateTaskService(DB, Task5)

	rows, _ := DB.Query(GetAllTasksQuery)

	tasks, _ := auxiliarFunc.GetList(rows)

	if len(tasks) != 5 {
		t.Fail()
	}
}

func TestGetTasksService(t *testing.T) {

	user1Tasks, _ := services.GetTasksService(DB, User1)
	user2Tasks, _ := services.GetTasksService(DB, User2)
	user3Tasks, _ := services.GetTasksService(DB, User3) // user 3 is not in DB

	if len(user1Tasks) != 3 {
		t.Fail()
	}
	if len(user2Tasks) != 2 {
		t.Fail()
	}
	if len(user3Tasks) != 0 {
		t.Fail()
	}
}

func TestGetTaskService(t *testing.T) {

	task1, _ := services.GetTaskService(DB, User1, "1")
	task2, _ := services.GetTaskService(DB, User1, "8")

	if task1.Description != "Do something" {
		t.Fail()
	}
	if task2.Description != "" {
		t.Fail()
	}
}

func TestGetTasksByDescriptionService(t *testing.T) {

	tasksByDescription1, _ := services.GetTasksByDescriptionService(DB, DescriptionTask1)
	tasksByDescription2, _ := services.GetTasksByDescriptionService(DB, DescriptionTask2)

	if len(tasksByDescription1) != 2 {
		t.Fail()
	}
	if len(tasksByDescription2) != 0 {
		t.Fail()
	}

}

func TestGetTasksByStateService(t *testing.T) {

	_, _ = DB.Query(FulfillTask, 2)
	_, _ = DB.Query(FulfillTask, 3)

	tasksByState, _ := services.GetTasksByStateService(DB, StateTasks)

	if len(tasksByState) != 2 {
		t.Fail()
	}
}

func TestDeleteTasksService(t *testing.T) {

	_ = services.DeleteTaskService(DB, Task5)

	rows, _ := DB.Query(GetAllTasksQuery)

	tasks, _ := auxiliarFunc.GetList(rows)

	if len(tasks) != 4 {
		t.Fail()
	}
}

func TestUpdateTaskService(t *testing.T) {

	_ = services.UpdateTaskService(DB, UpdateTask)

	taskUpdated, _ := services.GetTaskService(DB, User2, "4")

	if taskUpdated.Fulfilled != true {
		t.Fail()
	}

	_, _ = DB.Query(DestroyTables)

}
