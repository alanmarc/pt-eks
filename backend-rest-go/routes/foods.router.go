package routes

import (
	"encoding/json"
	"net/http"

	"github.com/alanmarc/backend-rest-go/db"
	"github.com/alanmarc/backend-rest-go/models"
	"github.com/gorilla/mux"
)

func GetFoodsHandler(w http.ResponseWriter, r *http.Request) {
	var Foods []models.Food
	json.NewDecoder(r.Body).Decode(&Foods)
	findFoods := db.DB.Find(&Foods)
	if findFoods.Error != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(findFoods.Error.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&Foods)

}

func GetFoodHandler(w http.ResponseWriter, r *http.Request) {
	var food models.Food
	params := mux.Vars(r)
	db.DB.First(&food, params["id"])

	if food.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Food not found"))
		return
	}

	json.NewEncoder(w).Encode(&food)
}

func PostFoodHandler(w http.ResponseWriter, r *http.Request) {
	var foods models.Food
	json.NewDecoder(r.Body).Decode(&foods)
	createdFoods := db.DB.Create(&foods)
	if createdFoods.Error != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(createdFoods.Error.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&foods)
}

func DeleteFoodHandler(w http.ResponseWriter, r *http.Request) {
	var food models.Food
	params := mux.Vars(r)
	db.DB.First(&food, params["id"])

	if food.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Food not found"))
		return
	}
	db.DB.Unscoped().Delete(&food)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Food Deleted"))
}
