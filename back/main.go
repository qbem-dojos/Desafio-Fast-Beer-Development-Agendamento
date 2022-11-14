package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"math/rand"

	"github.com/gorilla/mux"
)


type Practitioner struct {
	ID     	string  	`json:"id"`
	Isbn   	string  	`json:"isbn"`
	Job		string  	`json:"job"`
	Author 	*Author 	`json:"author"`
}

// Author struct
type Author struct {
	Firstname 	string 	`json:"firstname"`
	Lastname  	string 	`json:"lastname"`
}

var practitioners []Practitioner

func main() {
	r := mux.NewRouter()

	practitioners = append(practitioners, Practitioner{ID: "1", Isbn: "448743", Job: "Doctor", Author: &Author{Firstname: "John", Lastname: "Doe"}})
	practitioners = append(practitioners, Practitioner{ID: "2", Isbn: "448743", Job: "Nurse", Author: &Author{Firstname: "Mary", Lastname: "Doe"}})

	r.HandleFunc("/practitioners", getPractitioners).Methods("GET")
	r.HandleFunc("/practitioners/{id}", getPractitioner).Methods("GET")
	r.HandleFunc("/practitioners", createPractitioner).Methods("POST")
	r.HandleFunc("/practitioners/{id}", updatePractitioner).Methods("PUT")
	r.HandleFunc("/practitioners/{id}", deletePractitioner).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", r))
}


func getPractitioners(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(practitioners)
}

func getPractitioner(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)	// Get params
	// Loop through practitioners and find one with the id from the params
	for _, item := range practitioners {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(practitioners)
}

func createPractitioner(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var practitioner Practitioner
	_ = json.NewDecoder(r.Body).Decode(&practitioner)
	practitioner.ID = strconv.Itoa(rand.Intn(1000000))	// Mock ID - not safe
	practitioners = append(practitioners, practitioner)
	json.NewEncoder(w).Encode(practitioner)
}

func updatePractitioner(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)	// Get params
	// Loop through practitioners and find one with the id from the params
	for index, item := range practitioners {
		if item.ID == params["id"] {
			practitioners = append(practitioners[:index], practitioners[index+1:]...)
			var practitioner Practitioner
			_ = json.NewDecoder(r.Body).Decode(&practitioner)
			practitioner.ID = strconv.Itoa(rand.Intn(1000000))	// Mock ID - not safe
			practitioners = append(practitioners, practitioner)
			json.NewEncoder(w).Encode(practitioner)
			return
		}
	}
	json.NewEncoder(w).Encode(practitioners)

}

func deletePractitioner(w http.ResponseWriter, r *http.Request){
	
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)	// Get params
	// Loop through practitioners and find one with the id from the params
	for index, item := range practitioners {
		if item.ID == params["id"] {
			practitioners = append(practitioners[:index], practitioners[index+1:]...)			
			break
		}
	}
	json.NewEncoder(w).Encode(practitioners)
}
