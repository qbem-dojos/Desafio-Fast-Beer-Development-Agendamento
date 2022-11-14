package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"math/rand"
	"fmt"
	
	"github.com/gorilla/mux"
)

type Contato struct {
	tipo 	string 		`json:"tipo"`
	value 	string 		`json:"value"`	
}

type Patient struct {
	Cpf     	string  	`json:"cpf"`
	Nome   		string  	`json:"nome"`
	Contato 	[]*Contato	`json:"contato"`
}

type Practitioner struct {
	Nome   	string  	`json:"nome"`
	Crm		string  	`json:"crm"`
}

type Agendamento struct {
	Data	 	string  	`json:"data"`
	Sala		string  	`json:"sala"`
}

type Appointment struct {
	ID string `json:"id"`
	Profisional *Practitioner `json:"practitioner"`
	Patiente *Patient `json:"patient"`
	Agendament *Agendamento `json:"agendamento"`
}

var appointments []Appointment
var lista_get []Appointment

func main() {
	r := mux.NewRouter()
	
	appointments = append(appointments, Appointment{ID: "1", Profisional: &Practitioner{Nome: "João", Crm: "654321"}, Patiente: &Patient{Cpf: "123456789", Nome: "Francisca", Contato: []*Contato{&Contato{tipo: "email", value: "879897"}, &Contato{tipo: "email", value: "879897"}}}, Agendament: &Agendamento{Data: "2020-10-10", Sala: "1"}})
	appointments = append(appointments, Appointment{ID: "2", Profisional: &Practitioner{Nome: "João", Crm: "123456"}, Patiente: &Patient{Cpf: "987654321", Nome: "Maria", Contato: []*Contato{&Contato{tipo: "email", value: "879819"}}}, Agendament: &Agendamento{Data: "2020-10-10", Sala: "1"}})
	appointments = append(appointments, Appointment{ID: "3", Profisional: &Practitioner{Nome: "João", Crm: "123456"}, Patiente: &Patient{Cpf: "656008979", Nome: "Joaquina", Contato: []*Contato{&Contato{tipo: "email", value: "879819"}}}, Agendament: &Agendamento{Data: "2020-10-10", Sala: "1"}})

	r.HandleFunc("/agendamentos", getAgendamentos).Methods("GET")
	r.HandleFunc("/agendamentos/{id}", getAgendamento).Methods("GET")
	r.HandleFunc("/agendamentos", createAgendamento).Methods("POST")
	r.HandleFunc("/agendamentos/{id}", updateAgendamento).Methods("PUT")
	r.HandleFunc("/agendamentos/{id}", deleteAgendamento).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", r))
}


func getAgendamentos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	crm, _ := strconv.Atoi(r.URL.Query().Get("crm"))
	
	lista_get = nil
	
	for _, item := range appointments {
		crmPro, _ := strconv.Atoi(item.Profisional.Crm)

		if (crmPro == crm) && (item.Agendament.Data == r.URL.Query().Get("data")){			
			lista_get = append(lista_get, item)
			
		}

	}
	json.NewEncoder(w).Encode(lista_get)
}
	
	

func getAgendamento(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)	// Get params
	// Loop through Agendamentos and find one with the id from the params
	for _, item := range appointments {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(appointments)
}

func createAgendamento(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var appointment Appointment
	_ = json.NewDecoder(r.Body).Decode(&appointments)
	appointment.ID = strconv.Itoa(rand.Intn(1000000))	// Mock ID - not safe
	appointments = append(appointments, appointment)
	json.NewEncoder(w).Encode(appointment)
}

func updateAgendamento(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)	// Get params
	// Loop through Agendamentos and find one with the id from the params
	for index, item := range appointments {
		if item.ID == params["id"] {
			appointments = append(appointments[:index], appointments[index+1:]...)
			var appointment Appointment
			_ = json.NewDecoder(r.Body).Decode(&appointments)
			appointment.ID = strconv.Itoa(rand.Intn(1000000))	// Mock ID - not safe
			appointments = append(appointments, appointment)
			json.NewEncoder(w).Encode(appointments)
			return
		}
	}
	json.NewEncoder(w).Encode(appointments)

}

func deleteAgendamento(w http.ResponseWriter, r *http.Request){
	
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)	// Get params
	// Loop through Agendamentos and find one with the id from the params
	for index, item := range appointments {
		if item.ID == params["id"] {
			appointments = append(appointments[:index], appointments[index+1:]...)			
			break
		}
	}
	json.NewEncoder(w).Encode(appointments)
}
