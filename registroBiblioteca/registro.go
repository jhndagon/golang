package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Persona struct {
	Nombres    string `json:"nombres,omitempty"`
	Correo     string `json:"correo,omitempty"`
	Usuario    string `json:"usuario,omitempty"`
	Contrasena string `json:"contrasena,omitempty"`
}

var personas []Persona

//se obtiene la informacion de todas las personas
func GetPersonas(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(personas)
}

//se obtiene información de un usuario
func GetPersona(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	for _, item := range personas {
		if item.Usuario == parametros["usuario"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Persona{})
}

// Registra una nueva persona
func CrearPersona(w http.ResponseWriter, r *http.Request) {
	var persona Persona
	_ = json.NewDecoder(r.Body).Decode(&persona)
	personas = append(personas, persona)
	json.NewEncoder(w).Encode(personas)
}

func main() {
	router := mux.NewRouter().StrictSlash(false)
	personas = append(personas, Persona{Nombres: "Andrés Goméz Zapata", Correo: "andres.gomez@correo.com",
		Usuario: "andresG", Contrasena: "1234"})
	personas = append(personas, Persona{Nombres: "Carlos Perez Villa", Correo: "carlos.Perez@correo.com",
		Usuario: "carlosP", Contrasena: "7654"})
	router.HandleFunc("/personas", GetPersonas).Methods("GET")
	router.HandleFunc("/personas/{usuario}", GetPersona).Methods("GET")
	router.HandleFunc("/personas", CrearPersona).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", router))
}
