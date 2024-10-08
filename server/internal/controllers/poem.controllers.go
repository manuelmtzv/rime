package controllers

import "net/http"

type PoemController struct {
}

func (p PoemController) GetPoems(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Poems"))
}

func (p PoemController) GetPoem(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Poem"))
}

func (p PoemController) CreatePoem(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create Poem"))
}

func (p PoemController) UpdatePoem(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Update Poem"))
}

func (p PoemController) DeletePoem(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete Poem"))
}
