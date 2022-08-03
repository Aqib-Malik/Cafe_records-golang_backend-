package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func CreateEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var emp Cafe
	json.NewDecoder(r.Body).Decode(&emp)
	Database.Create(&emp)
	json.NewEncoder(w).Encode(emp)

}
func GetEmployees(w http.ResponseWriter, r *http.Request) {
	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var employees []Cafe
	Database.Find(&employees)
	json.NewEncoder(w).Encode(employees)

}

func GetEmployeesById(w http.ResponseWriter, r *http.Request) {
	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	var employee Cafe
	Database.Find(&employee, mux.Vars(r)["eid"])
	json.NewEncoder(w).Encode(employee)

}

func UpdateEmployeesById(w http.ResponseWriter, r *http.Request) {
	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	var employee Cafe
	Database.Find(&employee, mux.Vars(r)["eid"])
	json.NewDecoder(r.Body).Decode(&employee)
	Database.Save(&employee)
	json.NewEncoder(w).Encode(employee)

}

func DeleteEmployeesById(w http.ResponseWriter, r *http.Request) {
	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE")
	var employee Cafe
	Database.Delete(&employee, mux.Vars(r)["eid"])

	json.NewEncoder(w).Encode("Record deleted successfully....")

}

// Authentication

// func Signup(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Access-Control-Allow-Origin", "*")

// 	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

// 	var user User
// 	json.NewDecoder(r.Body).Decode(&user)
// 	Database.Create(&user)
// 	json.NewEncoder(w).Encode(user)

// }
