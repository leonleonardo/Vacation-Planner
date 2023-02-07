package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"vacation-planner/models"

	"github.com/gorilla/mux"
)

// Create user POST using mock data & parameter from the request
func (h DBRouter) CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	username := params["username"]
	var user models.User

	user.Firstname = "Kendy"
	user.Lastname = "Phillips"
	user.Username = username
	user.Password = "1234"

	// Creating new user and checking for error
	if result := h.DB.Create(&user); result.Error != nil {
		fmt.Println(result.Error)
	}

	json.NewEncoder(w).Encode(user)
}
