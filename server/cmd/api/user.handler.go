package main

import "net/http"

func (app *application) GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Users"))
}

func (app *application) GetUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("User"))
}

func (app *application) CreateUser(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name     string `json:"name"`
		LastName string `json:"last_name"`
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.errorResponse(w, r, http.StatusBadRequest, err)
		return
	}

}

func (app *application) UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Update User"))
}

func (app *application) DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete User"))
}
