package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

var (
	count    = 0
	contacts []Contact
)

type Contact struct {
	ID_Contacto    int            `json:"id"`
	NombreCompleto string         `json:"name"`
	Email          string         `json:"email"`
	TelMovil       string         `json:"phone"`
	FechaCaptura   string         `json:"date"`
	TemaDeInteres  *TemaDeInteres `json:"theme"`
	Mensaje        string         `json:"message"`
	Estado         *Estado        `json:"state"`
}

type TemaDeInteres struct {
	ID_TemaDeInteres int    `json:"id_theme"`
	Nombre           string `json:"name_theme"`
}

type Estado struct {
	ID_Estado int    `json:"id_state"`
	Nombre    string `json:"name_state"`
}

func main() {
	fmt.Println("Welcome")
	contacts = append(contacts, Contact{ID_Contacto: 1, NombreCompleto: "Jorge", Email: "", TelMovil: "", FechaCaptura: "", TemaDeInteres: &TemaDeInteres{ID_TemaDeInteres: 1, Nombre: "VideoJuegos"}, Mensaje: "", Estado: &Estado{ID_Estado: 1, Nombre: "Mexico"}})
	contacts = append(contacts, Contact{ID_Contacto: 2, NombreCompleto: "Angelica", Email: "", TelMovil: "", FechaCaptura: "", TemaDeInteres: &TemaDeInteres{ID_TemaDeInteres: 2, Nombre: "Cursos"}, Mensaje: "", Estado: &Estado{ID_Estado: 2, Nombre: "USA"}})
	contacts = append(contacts, Contact{ID_Contacto: 3, NombreCompleto: "Cigales", Email: "", TelMovil: "", FechaCaptura: "", TemaDeInteres: &TemaDeInteres{ID_TemaDeInteres: 3, Nombre: "Novelas"}, Mensaje: "", Estado: &Estado{ID_Estado: 3, Nombre: "Canada"}})

	router := mux.NewRouter()
	router.HandleFunc("/", LandingPage)
	router.HandleFunc("/contactos", GetContactosEndPoint).Methods("GET")
	router.HandleFunc("/contacto/{id}", GetContactoEndPoint).Methods("GET")
	router.HandleFunc("/contactos", PostCountContactos).Methods("POST")

	handleCORS := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	}).Handler(router)

	http.ListenAndServe("localhost:3030", handleCORS)
}

func LandingPage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Landing Page")
	count = count + 1
	visitsCount := strconv.Itoa(count)
	fmt.Println("Visitas: " + visitsCount)
}

func GetContactosEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Println("- Reporte de Contactos")
	json.NewEncoder(w).Encode(contacts)
}

func GetContactoEndPoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	indice, _ := strconv.Atoi(params["id"])
	json.NewEncoder(w).Encode(contacts[indice])
}

func PostCountContactos(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Nuevo Contacto")
	var newContact Contact
	json.NewDecoder(r.Body).Decode(&newContact)
	contacts = append(contacts, newContact)

	cont := strconv.Itoa(len(contacts))
	fmt.Println("Numero de Contactos: " + cont)
	json.NewEncoder(w).Encode(newContact)
}
