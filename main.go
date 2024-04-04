package main

import (
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

	http.ListenAndServe(":3000", r)
}
