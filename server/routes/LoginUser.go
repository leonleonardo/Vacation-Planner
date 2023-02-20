package routes

import (
	"encoding/json"
	"net/http"
	"vacation-planner/models"
)

// Login user GET, using HTTP request body information for email and password
func (h DBRouter) LoginUser(w http.ResponseWriter, r *http.Request) {

	// Only GET is allowed for this route
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	  }

	// Creating new variable for storing request body
	var requestBody map[string]interface{}

	// Decoding body of the http request for the information for the user account
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Added a line to check the database if a user exist with same email given
	existingUser := &models.User{}
	result := h.DB.First(existingUser, "Email = ?", requestBody["Email"].(string))

	// Checking if the rows that have the email isn't 0 therefore somebody has the email
	if result.RowsAffected != 0 {

		// Checking if the password given in the request matches the password stored for the user in the DB
		if existingUser.Password != requestBody["Password"].(string) {
			// Will eventually encode a response for password not matching
			w.Write([]byte("Password given does not match password associated with this email!"))
		} else {

			// Setting headers
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)

			// Essentially I would be encoding some sort of response with 
			// "json.NewEncoder(w).Encode(response)"
			//
			// For now, I'm just printing validation strings.

			w.Write([]byte("User successfully logged in."))
		}
	} else {
		// If Rows Affected (rows with email given) is  0, therefore nobody has an account with
		// the email given, we can't login the user and tell them their email is not in use by our website.
		w.Write([]byte("Email not in use!"))
	}
}
