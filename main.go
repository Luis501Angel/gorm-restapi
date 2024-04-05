package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/luis501angel/go-gorm-restapi/db"
	"github.com/luis501angel/go-gorm-restapi/models"
	"github.com/luis501angel/go-gorm-restapi/routes"
)

func main() {

	db.DBConnection()

	db.DB.AutoMigrate(models.Task{})
	db.DB.AutoMigrate(models.User{})

	r := mux.NewRouter()

	r.HandleFunc("/", routes.HomeHandler)

	r.HandleFunc("/users", routes.GetUsersHandler).Methods("GET")
	r.HandleFunc("/users/{id}", routes.GetUserHandler).Methods("GET")
	r.HandleFunc("/users", routes.PostUserHandler).Methods("POST")
	r.HandleFunc("/users", routes.DeleteUserHandler).Methods("DELETE")

	err := http.ListenAndServe(":4000", r)
	if err != nil {
		fmt.Println(err.Error())
	}
}
