package middleware

import (
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
)

var collection *mongo.Collection

func init() {
	loadTheEnv()
	createDBInstance()
}

func loadTheEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading the .env file")
	}
}

func GetAllTasks(w http.ResponseWriter, r *http.Request) {

}

func CreateTask() {

}

func TaskComplete() {

}

func UndoTask() {

}

func DeleteTask() {

}

func DeleteAllTasks() {

}
