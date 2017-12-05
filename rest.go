package main

import (
	"log"
	"fmt"
	"net/http"
	"encoding/json"

	"github.com/gorilla/mux"
)

type Person struct {
	ID		string `json:"id,omitempty"`
	FirstName	string `json:"firstname,omitempty"`
	LastName	string `json:"lastname,omitempty"`
	Address		*Address `json:"address,omitempty"`
}

type Address struct {
	City		string `json:"city,omitempty"`
	State		string `json:"state,omitempty"`
}

var people = make([]Person, 0, 10)

func main() {
	people = append(people, Person{ID: "1", FirstName: "John", LastName: "Doe", Address: &Address{City: "City X", State: "State X"}})
	people = append(people, Person{ID: "2", FirstName: "Koko", LastName: "Doe", Address: &Address{City: "City Z", State: "State Y"}})

	router := mux.NewRouter().StrictSlash(true)
	//router.HandleFunc("/", Index)
	router.HandleFunc("/people", GetPeople).Methods("GET")
	router.HandleFunc("/people/{id}", GetPerson).Methods("GET")
	router.HandleFunc("/people/{id}", CreatePerson).Methods("POST")
	router.HandleFunc("/people/{id}", DeletePerson).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}

func GetPeople(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Go REST Web Application V1 !\n")

	json.NewEncoder(w).Encode(people)
}

func GetPerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range people {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Person{})
}

func CreatePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var person Person
	json.NewDecoder(r.Body).Decode(&person)
	person.ID = params["id"]
	people = append(people, person)
	json.NewEncoder(w).Encode(people)
}


func DeletePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range people {
		if item.ID == params["id"] {
			fmt.Printf("find the id: %s\n", params["id"])
			people = append(people[:index], people[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(people)
}
