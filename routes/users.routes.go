package routes

import (
	"encoding/json"
	"net/http"

	"github.com/FernandoDanielZigarra/rest-api-postgres/db"
	"github.com/FernandoDanielZigarra/rest-api-postgres/models"
	"github.com/gorilla/mux"
)

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	db.DB.Preload("Tasks").Find(&users)
	json.NewEncoder(w).Encode(&users)
}
func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	params := mux.Vars(r)
	db.DB.Preload("Tasks").First(&user, "id = ?", params["id"])
	if user.ID == "" {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User not found"))
		return
	}
	/* db.DB.Model(&user).Association("Tasks").Find(&user.Tasks) */
	json.NewEncoder(w).Encode(&user)
}
func PostUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	db.DB.First(&user, "email = ?", user.Email)

	if user.ID != "" {
		w.WriteHeader(http.StatusConflict)
		w.Write([]byte("Email already exists"))
		return
	}

	createdUser := db.DB.Create(&user)
	err := createdUser.Error

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	json.NewEncoder(w).Encode(&user)

}

func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	params := mux.Vars(r)

	db.DB.Preload("Tasks").First(&user, "id = ?", params["id"])

	if user.ID == "" {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User not found"))
		return
	}

	// Decode the request body to update the user
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Failed to decode request body"))
		return
	}

	// Save the updated user to the database
	db.DB.Save(&user.FirstName)
	db.DB.Save(&user.LastName)
	db.DB.Save(&user.Email)

	// Respond with the updated user in the response body
	json.NewEncoder(w).Encode(&user)
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	params := mux.Vars(r)
	db.DB.First(&user, "id = ?", params["id"])

	if user.ID == "" {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User not found"))
		return
	}

	db.DB.Unscoped().Delete(&user)
	w.WriteHeader(http.StatusNoContent)

}
