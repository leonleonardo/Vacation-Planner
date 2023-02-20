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

	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	  }

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	// Changed this to not create a mock user until we have checked
	// the DB for any users with the given username
	var user models.User

	// Added a line to check the database for any users with the same username
	// as the new account
	result := h.DB.First(&models.User{}, "Username = ?", params["username"])

	// Checking if the rows that have the username is 0 therefore nobody has the username
	// still using mock information until figuring out how to work with http request BODY(?)
	if result.RowsAffected == 0 {
		user.Firstname = "Top"
		user.Lastname = "G"
		user.Username = params["username"]
		user.Password = "blackPowerRanger"
	
		// Creating new user and checking for error
		if newUser := h.DB.Create(&user); newUser.Error != nil {
			fmt.Println(newUser.Error)
		}

		json.NewEncoder(w).Encode(user)

	} else {
		w.Write([]byte("Username taken!"))
	}
}
