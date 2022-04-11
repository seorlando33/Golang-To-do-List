package controllers

import (
	//"database/sql"
	"encoding/json"
	"net/http"

	connection "toDoList/api/db"
	auxiliarFunc "toDoList/api/lib"
	models "toDoList/api/models"
	services "toDoList/api/services"
)

func CreateTask(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		auxiliarFunc.BadGateway(w)
		return
	}

	db, err := connection.DBConnection()
	if err != nil {
		auxiliarFunc.InternalServerError(w, err)
		return
	}

	task := &models.Task{}

	err = json.NewDecoder(r.Body).Decode(task)
	if err != nil {
		auxiliarFunc.BadRequest(w, err)
		return
	}

	err = services.CreateTaskService(db, *task)

	if err != nil {
		auxiliarFunc.BadRequest(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func GetTasks(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
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

	tasks, err := services.GetTasksService(db, *user)

	if err != nil {
		auxiliarFunc.InternalServerError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)

}

func GetTask(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		auxiliarFunc.BadGateway(w)
		return
	}

	db, err := connection.DBConnection()
	if err != nil {
		auxiliarFunc.InternalServerError(w, err)
		return
	}

	url := r.URL.Query()
	id := url["id"][0]

	user := &models.User{}

	err = json.NewDecoder(r.Body).Decode(user)

	if err != nil {
		auxiliarFunc.BadRequest(w, err)
		return
	}

	task, err := services.GetTaskService(db, *user, id)

	if err != nil {
		auxiliarFunc.InternalServerError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)

}

func GetTasksByDescription(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		auxiliarFunc.BadGateway(w)
		return
	}

	db, err := connection.DBConnection()
	if err != nil {
		auxiliarFunc.InternalServerError(w, err)
		return
	}

	task := &models.Task{}

	err = json.NewDecoder(r.Body).Decode(task)

	if err != nil {
		auxiliarFunc.BadRequest(w, err)
		return
	}

	tasks, err := services.GetTasksByDescriptionService(db, *task)

	if err != nil {
		auxiliarFunc.InternalServerError(w, err)
		return
	}

	if len(tasks) == 0 {
		w.Write([]byte("Any task match the pattern!"))
	} else {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(tasks)
	}
}

func GetTasksByState(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		auxiliarFunc.BadGateway(w)
		return
	}

	db, err := connection.DBConnection()
	if err != nil {
		auxiliarFunc.InternalServerError(w, err)
		return
	}

	task := &models.Task{}

	err = json.NewDecoder(r.Body).Decode(task)
	if err != nil {
		auxiliarFunc.BadRequest(w, err)
		return
	}

	tasks, err := services.GetTasksByStateService(db, *task)

	if err != nil {
		auxiliarFunc.InternalServerError(w, err)
		return
	}

	if len(tasks) == 0 {
		w.Write([]byte("Any task match the pattern!"))
	} else {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(tasks)
	}

}

func DeleteTask(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodDelete {
		auxiliarFunc.BadGateway(w)
		return
	}

	db, err := connection.DBConnection()
	if err != nil {
		auxiliarFunc.InternalServerError(w, err)
		return
	}

	task := &models.Task{}

	err = json.NewDecoder(r.Body).Decode(task)

	if err != nil {
		auxiliarFunc.BadRequest(w, err)
		return
	}

	err = services.DeleteTaskService(db, *task)

	if err != nil {
		auxiliarFunc.InternalServerError(w, err)
		return
	}

	w.WriteHeader(http.StatusAccepted)
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPut {
		auxiliarFunc.BadGateway(w)
		return
	}

	db, err := connection.DBConnection()
	if err != nil {
		auxiliarFunc.InternalServerError(w, err)
		return
	}

	task := &models.Task{}

	err = json.NewDecoder(r.Body).Decode(task)

	if err != nil {
		auxiliarFunc.BadRequest(w, err)
		return
	}

	err = services.UpdateTaskService(db, *task)

	if err != nil {
		auxiliarFunc.InternalServerError(w, err)
		return
	}

	w.WriteHeader(http.StatusAccepted)
}
