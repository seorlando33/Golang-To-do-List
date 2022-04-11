package api

import (
	"net/http"

	auth "toDoList/api/middlewares"
	controllers "toDoList/api/controllers"
)

func initRoutes() {

	go http.HandleFunc("/register", controllers.Register)
	go http.HandleFunc("/login", controllers.Login)


	go http.HandleFunc("/createTask", auth.JWTValidator(controllers.CreateTask))

	go http.HandleFunc("/getTasks", auth.JWTValidator(controllers.GetTasks))
	go http.HandleFunc("/getTask", auth.JWTValidator(controllers.GetTask))
	go http.HandleFunc("/getTasksByDescription", auth.JWTValidator(controllers.GetTasksByDescription))
	go http.HandleFunc("/getTasksByState", auth.JWTValidator(controllers.GetTasksByState))

	go http.HandleFunc("/deleteTask", auth.JWTValidator(controllers.DeleteTask))

	go http.HandleFunc("/updateTask", auth.JWTValidator(controllers.UpdateTask))

}
