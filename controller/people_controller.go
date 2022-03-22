package controller

import (
	"encoding/json"
	"net/http"
	"os/exec"

	"github.com/gorilla/mux"
)

type Person struct {
	ID        string   `json:"id,omitempty"`
	Firstname string   `json:"firstname,omitempty"`
	Lastname  string   `json:"lastname,omitempty"`
	Address   *Address `json:"address,omitempty"`
}

type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

var people []Person

func init() {
	people = append(people, Person{ID: "87965fb2-a9fb-11ec-b909-0242ac120002", Firstname: "John", Lastname: "Doe", Address: &Address{City: "City X", State: "State X"}})
	people = append(people, Person{ID: "955c4b48-a9fb-11ec-b909-0242ac120003", Firstname: "Koko", Lastname: "Doe", Address: &Address{City: "City Z", State: "State Y"}})

}

func GetPeople(w http.ResponseWriter, request *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(people)

}

func GetPerson(w http.ResponseWriter, request *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(request)

	for _, item := range people {

		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}

	}

	w.WriteHeader(http.StatusNotFound)

}

func CreatePerson(w http.ResponseWriter, request *http.Request) {

	var person Person

	json.NewDecoder(request.Body).Decode(&person)

	uuid, _ := exec.Command("uuidgen").Output()

	person.ID = string(uuid)

	people = append(people, person)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(people)

}

func UpdatePerson(w http.ResponseWriter, r *http.Request) {

	var person Person

	json.NewDecoder(r.Body).Decode(&person)

	for index, item := range people {

		if item.ID == person.ID {
			people = append(people[:index], people[index+1:]...)
			people = append(people, person)

			w.WriteHeader(http.StatusAccepted)
			return

		}

	}

	w.WriteHeader(http.StatusNotFound)

}

func DeletePerson(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index, item := range people {

		if item.ID == params["id"] {
			people = append(people[:index], people[index+1:]...)
			w.WriteHeader(http.StatusAccepted)
			return

		}

	}

	w.WriteHeader(http.StatusNotFound)

}
