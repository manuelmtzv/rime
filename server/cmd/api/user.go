package main

import "net/http"

func (app *application) findUsers(w http.ResponseWriter, r *http.Request) {
	users, err := app.store.Users.FindAll(r.Context())
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	err = app.jsonResponse(w, http.StatusOK, users)

	if err != nil {
		app.internalServerError(w, r, err)
	}
}
