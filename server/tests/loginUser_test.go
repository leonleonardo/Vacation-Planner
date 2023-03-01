package tests

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"vacation-planner/database"
	"vacation-planner/routes"

	"github.com/gorilla/mux"
)

func TestUserLogin(t *testing.T) {
	//test each possible response from the server by inputting info and testing against the expected validation string

	//case 1: successful login

	//this email and pass combination is stored in the database
	email := "123@gmail.com"
	password := "123"

	payload := fmt.Sprintf(`{"Email": "%s", "Password": "%s"}`, email, password)

	//making a request to the database
	req, err := http.NewRequest("GET", "/loginUser", strings.NewReader(payload))
	if err != nil {
		t.Errorf("Error: login user request could not be completed")
	}

	//response recorder: takes in the returned bytestring
	rr := httptest.NewRecorder()

	db, err := database.Connect()
	h := routes.NewConnection(db)

	r := mux.NewRouter()
	r.HandleFunc("/loginUser", h.LoginUser).Methods("GET")
	r.ServeHTTP(rr, req)

	//error handling for wrong http status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	//with this email and password, we expect the user to be logged in
	expectedOutput := []byte("User successfully logged in.")

	//testing if the response matches output
	if !bytes.Equal(rr.Body.Bytes(), expectedOutput) {
		t.Errorf("handler returned unexpected response: got %v want %v", rr.Body.String(), expectedOutput)
	} else {
		fmt.Printf("Case 1 passed!!\n")
	}

	//case 2: email not in use
	email2 := "a@gmail.com"
	password2 := "password"

	payload2 := fmt.Sprintf(`{"Email": "%s", "Password": "%s"}`, email2, password2)

	req2, err := http.NewRequest("GET", "/loginUser", strings.NewReader(payload2))
	if err != nil {
		t.Errorf("Error: login user request could not be completed")
	}

	rr2 := httptest.NewRecorder()
	r2 := mux.NewRouter()
	r2.HandleFunc("/loginUser", h.LoginUser).Methods("GET")
	r2.ServeHTTP(rr2, req2)

	if status := rr2.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	//this email is not in the database of users
	expectedOutput2 := []byte("Email not in use!")

	//testing if the response matches output
	if !bytes.Equal(rr2.Body.Bytes(), expectedOutput2) {
		t.Errorf("handler returned unexpected response: got %v want %v", rr2.Body.String(), expectedOutput2)
	} else {
		fmt.Printf("Case 2 passed!!\n")
	}

	//case 3: password does not match
	email3 := "123@gmail.com"
	password3 := "password"

	payload3 := fmt.Sprintf(`{"Email": "%s", "Password": "%s"}`, email3, password3)

	req3, err := http.NewRequest("GET", "/loginUser", strings.NewReader(payload3))
	if err != nil {
		t.Errorf("Error: login user request could not be completed")
	}

	rr3 := httptest.NewRecorder()
	r3 := mux.NewRouter()
	r3.HandleFunc("/loginUser", h.LoginUser).Methods("GET")
	r3.ServeHTTP(rr3, req3)

	if status := rr3.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	//the password does not match with the user input in the database
	expectedOutput3 := []byte("Password given does not match password associated with this email!")

	//testing to see if expected response matches output
	if !bytes.Equal(rr3.Body.Bytes(), expectedOutput3) {
		t.Errorf("handler returned unexpected response: got %v want %v", rr3.Body.String(), expectedOutput3)
	} else {
		fmt.Printf("Case 3 passed!!\n")
	}

}
