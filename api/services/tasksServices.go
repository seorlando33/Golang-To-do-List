package services

import (
	"database/sql"
	"errors"

	auxiliarFunc "toDoList/api/lib"
	models "toDoList/api/models"
)

func CreateTaskService(db *sql.DB, task models.Task) error {

	if task.Description == "" {
		return errors.New("description is empty")
	}

	query := `INSERT INTO tasks(user_id, description) VALUES ($1, $2)`
	_, err := db.Query(query, task.User_ID, task.Description)

	if err != nil {
		return err
	}
	return nil
}

func GetTasksService(db *sql.DB, user models.User) ([]models.Task, error) {

	query := `SELECT * FROM tasks WHERE user_id = $1`
	rows, err := db.Query(query, user.ID)

	if err != nil {
		return nil, err
	}

	tasks, err := auxiliarFunc.GetList(rows)

	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func GetTaskService(db *sql.DB, user models.User, id string) (models.Task, error) {

	var task models.Task

	query := `SELECT * FROM tasks WHERE user_id = $1 AND id = $2`
	row := db.QueryRow(query, user.ID, id)

	err := row.Scan(&task.ID, &task.User_ID, &task.Description, &task.Fulfilled, &task.Created_At, &task.Updated_At)

	if err != nil {
		return task, err
	}

	return task, nil
}

func GetTasksByDescriptionService(db *sql.DB, task models.Task) ([]models.Task, error) {
	
	if task.Description == "" {
		return nil, errors.New("description is empty")
	}

	task.Description = "%" + task.Description + "%"

	query := `SELECT * FROM tasks WHERE user_id = $1 AND description ILIKE $2`
	rows, err := db.Query(query, task.User_ID, task.Description)

	if err != nil {
		return nil, err
	}

	tasks, err := auxiliarFunc.GetList(rows)

	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func GetTasksByStateService(db *sql.DB, task models.Task) ([]models.Task, error) {

	query := `SELECT * FROM tasks WHERE user_id = $1 AND fulfilled = $2`
	rows, err := db.Query(query, task.User_ID, task.Fulfilled)

	if err != nil {
		return nil, err
	}

	tasks, err := auxiliarFunc.GetList(rows)

	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func DeleteTaskService (db *sql.DB, task models.Task) (error) {

	query := `DELETE FROM tasks WHERE id = $1`
	_, err := db.Query(query, task.ID)

	if err != nil {
		return err
	}
	return nil
}

func UpdateTaskService(db *sql.DB, task models.Task) (error) {

	query := `UPDATE tasks SET fulfilled = $1 WHERE id = $2`
	_, err := db.Query(query, task.Fulfilled, task.ID)

	if err != nil {
		return err
	}
	return nil
}