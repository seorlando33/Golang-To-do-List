package lib

import (
	"database/sql"

	models "toDoList/api/models"
)

func GetList(rows *sql.Rows) ([]models.Task, error) {
	var tasks []models.Task

	for rows.Next() {
		var task models.Task

		err := rows.Scan(&task.ID, &task.User_ID, &task.Description, &task.Fulfilled, &task.Created_At, &task.Updated_At)
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, task)
	}

	return tasks, nil
}
