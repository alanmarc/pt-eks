package main

import (
	"net/http"

	"github.com/alanmarc/backend-rest-go/db"
	"github.com/alanmarc/backend-rest-go/models"
	"github.com/alanmarc/backend-rest-go/routes"
	"github.com/gorilla/mux"
)

func main() {

	db.DBConnection()
	db.DB.AutoMigrate(models.User{})
	db.DB.AutoMigrate(models.Food{})
	r := mux.NewRouter()
	//cors
	headers := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	}
	r.HandleFunc("/", routes.HomeHandler)
	//Users
	r.HandleFunc("/users", headers).Methods(http.MethodOptions)
	r.HandleFunc("/users", routes.GetUsersHandler).Methods(http.MethodOptions)
	r.HandleFunc("/users/{id}", headers).Methods(http.MethodOptions)
	r.HandleFunc("/users/{id}", routes.GetUserHandler).Methods("GET")
	r.HandleFunc("/users", routes.PostUserHandler).Methods("POST")
	r.HandleFunc("/users/{id}", routes.DeleteUserHandler).Methods("DELETE")

	//Foods
	r.HandleFunc("/users", headers).Methods(http.MethodOptions)
	r.HandleFunc("/foods", routes.GetFoodsHandler).Methods("GET")
	r.HandleFunc("/foods/{id}", routes.GetFoodHandler).Methods("GET")
	r.HandleFunc("/foods", routes.PostFoodHandler).Methods("POST")
	r.HandleFunc("/foods/{id}", routes.DeleteFoodHandler).Methods("DELETE")

	http.ListenAndServe(":3000", r)
}
