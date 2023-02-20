package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"vacation-planner/models"
)

// Create user POST, using HTTP request body information for email and password
func (h DBRouter) CreateUser(w http.ResponseWriter, r *http.Request) {

	// Only POST is allowed for this route
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	  }

	// Creating two new variables to use as reference
	var requestBody map[string]interface{}
	var user models.User

	// Decoding body of the http request for the information for the user account
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Added a line to check the database for any users with the same email as the new account
	
	result := h.DB.First(&models.User{}, "Email = ?", requestBody["Email"].(string))

	// Checking if the rows that have the email is 0 therefore nobody has the email
	if result.RowsAffected == 0 {
		// Assigning Email and Password to new User
		user.Email = requestBody["Email"].(string)
		user.Password = requestBody["Password"].(string)
	
		// Creating new user in the DB and checking for error
		if newUser := h.DB.Create(&user); newUser.Error != nil {
			fmt.Println(newUser.Error)
		}

		// Setting headers
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		
		// Not generating accurate response data(?), looking to meet to come to concensus on how exactly
		// we plan on passing information back and forth
		//
		// Essentially I would be encoding some sort of response with 
		// "json.NewEncoder(w).Encode(response)"
		//
		// For now, I'm just printing validation strings.

		w.Write([]byte("New user successfully created."))

	} else {
		// If Rows Affected (rows with email given) is greater than 0, therefore someone has an account with
		// the email given, we don't create a new user and tell them their email is taken.
		w.Write([]byte("Email taken!"))
	}
}
