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

func GetPersona(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	for _, item := range personas {
		if item.Usuario == parametros["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(&Persona{})
}

// Registra una nueva persona
func CrearPersona(w http.ResponseWriter, r *http.Request) {
	//params := mux.Vars(r)
	var persona Persona
	_ = json.NewDecoder(r.Body).Decode(&persona)
	personas = append(personas, persona)
	json.NewEncoder(w).Encode(personas)
}

func main() {
	router := mux.NewRouter().StrictSlash(false)
	personas = append(personas, Persona{Nombres: "Mateo Ya no", Correo: "mateo.llano@algo.com", Usuario: "mateito", Contrasena: "1234"})
	personas = append(personas, Persona{Nombres: "johndragon", Correo: "johnd.gonzalez@algo.com", Usuario: "jhndragon1", Contrasena: "7654"})
	personas = append(personas, Persona{Nombres: "Mateo Ya no", Correo: "mateo.llano@algo.com", Usuario: "mateito", Contrasena: "1234"})

	router.HandleFunc("/personas", GetPersonas).Methods("GET")
	router.HandleFunc("/personas/{id}", GetPersona).Methods("GET")
	router.HandleFunc("/personas", CrearPersona).Methods("POST")
	//router.HandleFunc("/people/{id}", DeletePerson).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}
