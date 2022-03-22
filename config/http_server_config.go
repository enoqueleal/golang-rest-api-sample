package config

import (
	"log"
	"my-books/controller"
	"net/http"

	"github.com/gorilla/mux"
)

func init() {

	router := mux.NewRouter()

	router.HandleFunc("/contato", controller.GetPeople).Methods("GET")
	router.HandleFunc("/contato", controller.CreatePerson).Methods("POST")
	router.HandleFunc("/contato", controller.UpdatePerson).Methods("PUT")
	router.HandleFunc("/contato/{id}", controller.GetPerson).Methods("GET")
	router.HandleFunc("/contato/{id}", controller.DeletePerson).Methods("DELETE")

	log.Println("Server listening on port: 8080")

	err := http.ListenAndServe(":8080", router)

	if err == nil {
		log.Fatal(err)
	}

}
