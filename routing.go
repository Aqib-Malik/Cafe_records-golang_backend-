package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandlerRouting() {
	r := mux.NewRouter()
	r.HandleFunc("/cafe_add", CreateEmployee).Methods("POST")
	r.HandleFunc("/cafe_all", GetEmployees).Methods("GET", "OPTIONS")
	r.HandleFunc("/cafe_get/{eid}", GetEmployeesById).Methods("GET")
	r.HandleFunc("/cafe_update/{eid}", UpdateEmployeesById).Methods("PUT")
	r.HandleFunc("/cafe_delete/{eid}", DeleteEmployeesById).Methods("DELETE", "OPTIONS")

	// Authentication
	// r.HandleFunc("/signup", Signup).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", r))
}
