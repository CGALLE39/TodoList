package router

import (
	"github.com/gorilla/mux"
	"totolist/server/middleware"
)

func Router() *mux.Router {

	router := mux.NewRouter()
	router.HandleFunc("/api/task", middleware.GetAllTasks).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/tasks", middleware.CreateTask).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/tasks/{id}", middleware.TaskComplete).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/undoTask/{id}", middleware.UndoTask).Methods("PUTS", "OPTIONS")
	router.HandleFunc("/api/deleteTask{id}", middleware.DeleteTask).Methods("DELETE", "OPTIONS")
	router.HandleFunc("/api/deleteAllTasks", middleware.deleteAllTasks).Methods("DELETE", "OPTIONS")
	return router
}
