package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type Note struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreateAt    time.Time `json:"create_at"`
}

var noteStore = make(map[string]Note)

var id int

func GetNotesHandler(w http.ResponseWriter, r *http.Request) {
	var not []Note
	for _, v := range noteStore {
		not = append(not, v)
	}
	w.Header().Set("Content-Type", "application/json")
	j, err := json.Marshal(not)
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

func PostNotesHandler(w http.ResponseWriter, r *http.Request) {
	var not Note
	err := json.NewDecoder(r.Body).Decode(&not)
	if err != nil {
		panic(err)
	}
	not.CreateAt = time.Now()
	id++
	k := strconv.Itoa(id)
	noteStore[k] = not
	w.Header().Set("Content-Type", "application/json")
	j, err := json.Marshal(not)
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusCreated)
	w.Write(j)
}

func PutNotesHandler(w http.ResponseWriter, r *http.Request) {

}

func DeleteNotesHandler(w http.ResponseWriter, r *http.Request) {

}

func main() {
	r := mux.NewRouter().StrictSlash(false)

	r.HandleFunc("/api/notes", GetNotesHandler).Methods("GET")
	r.HandleFunc("/api/notes", PostNotesHandler).Methods("POST")
	r.HandleFunc("/api/notes/{id]", PutNotesHandler).Methods("PUT")
	r.HandleFunc("/api/notes/{id}", DeleteNotesHandler).Methods("DELETE")

	http.ListenAndServe(":8000", r)
}
