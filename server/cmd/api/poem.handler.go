package main

import "net/http"

func (app *application) GetPoems(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Poems"))
}

func (app *application) GetPoem(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Poem"))
}

func (app *application) CreatePoem(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create Poem"))
}

func (app *application) UpdatePoem(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Update Poem"))
}

func (app *application) DeletePoem(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete Poem"))
}
